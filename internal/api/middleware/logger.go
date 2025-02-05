package middleware

import (
	"fmt"
	"github.com/ivofreitas/device-api/internal/adapter/context"
	"github.com/ivofreitas/device-api/internal/adapter/log"
	"github.com/labstack/echo/v4"
	"time"
)

// Logger - Generates a JSON with information of request
func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		ctx := log.InitParams(c.Request().Context())
		c.SetRequest(c.Request().WithContext(ctx))

		httpLog := context.Get(ctx, log.HTTPKey).(*log.HTTP)
		req := c.Request()
		httpLog.Request.Host = req.Host
		httpLog.Request.Route = fmt.Sprintf("[%s] %s", req.Method, req.URL.Path)
		httpLog.Request.Header = req.Header

		defer func() {
			res := c.Response()

			httpLog.Latency = float64(time.Since(start)/time.Millisecond) / 1000

			httpLog.Response.Header = res.Header()
			httpLog.Response.Status = res.Status
			httpLog.Response.RemoteIP = c.RealIP()

			entry := log.NewEntry()
			entry = entry.WithField("http", httpLog)

			if httpLog.Error != "" {
				entry.Error()
				return
			}
			entry.Info()
		}()

		return next(c)
	}
}
