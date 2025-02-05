package api

import (
	gocontext "context"
	"fmt"
	"github.com/ivofreitas/device-api/config"
	"github.com/ivofreitas/device-api/internal/adapter/context"
	"github.com/ivofreitas/device-api/internal/adapter/log"
	"github.com/ivofreitas/device-api/internal/api/middleware"
	"github.com/ivofreitas/device-api/internal/domain"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	nethttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	echo   *echo.Echo
	logger *logrus.Entry
	signal chan struct{}
}

func NewServer() *Server {
	log.Init()

	return &Server{
		logger: log.NewEntry(),
		signal: make(chan struct{}),
	}
}

func (s *Server) Run() {
	s.start()
	s.logger.Println("Server started and waiting for the graceful signal...")
	<-s.signal
}

func (s *Server) start() {
	go s.watchStop()

	env := config.GetEnv()

	s.initHttp()

	s.logger.Infof("Server is starting in port %s.", env.Server.Port)

	register(s.echo)

	addr := fmt.Sprintf(":%s", env.Server.Port)
	go func() {
		if err := s.echo.Start(addr); err != nil {
			s.logger.WithError(err).Fatal("Shutting down the server now")
		}
	}()
}

func (s *Server) watchStop() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	ctx, cancel := gocontext.WithTimeout(gocontext.Background(), 5*time.Second)
	defer cancel()

	s.logger.Info("Server is stopping...")

	err := s.echo.Shutdown(ctx)
	if err != nil {
		s.logger.Errorln(err)
	}

	close(s.signal)
}

func (s *Server) initHttp() {
	s.echo = echo.New()
	s.echo.Use(middleware.Logger)
	s.echo.Use(echomiddleware.Recover())
	s.echo.Pre(echomiddleware.RemoveTrailingSlash())
	s.echo.HTTPErrorHandler = func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		responseErr := &domain.Error{
			Type:   "Path not found",
			Status: c.Response().Status,
			Detail: err.Error(),
		}
		httpLog := context.Get(c.Request().Context(), log.HTTPKey).(*log.HTTP)
		httpLog.Error = err.Error()

		if c.Request().Method == nethttp.MethodHead {
			err = c.NoContent(responseErr.Status)
		} else {
			err = c.JSON(responseErr.Status, responseErr)
		}
		if err != nil {
			s.echo.Logger.Error(err)
		}
	}
}
