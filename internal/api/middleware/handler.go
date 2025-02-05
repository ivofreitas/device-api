package middleware

import (
	gocontext "context"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/ivofreitas/device-api/internal/adapter/context"
	"github.com/ivofreitas/device-api/internal/adapter/log"
	"github.com/ivofreitas/device-api/internal/domain"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
)

type ServiceFn func(ctx gocontext.Context, param interface{}) (interface{}, error)

type Handler struct {
	fn         ServiceFn
	param      interface{}
	httpStatus int
	echo.Binder
	*validator.Validate
}

func NewHandler(fn ServiceFn, httpStatus int, param interface{}) *Handler {
	return &Handler{fn, param, httpStatus, new(echo.DefaultBinder), validator.New()}
}

// Handle - Request's entry point - bind, validate and call internal business logic
func (ctrl *Handler) Handle(c echo.Context) error {

	ctx := c.Request().Context()
	httpLog := context.Get(ctx, log.HTTPKey).(*log.HTTP)

	if ctrl.param != nil {
		ctrl.param = reflect.New(reflect.TypeOf(ctrl.param).Elem()).Interface()
		if err := ctrl.bind(c); err != nil {
			var responseErr *domain.Error
			errors.As(err, &responseErr)
			httpLog.Error = responseErr.Error()
			return c.JSON(http.StatusBadRequest, responseErr)
		}

		if err := ctrl.validate(); err != nil {
			var responseErr *domain.Error
			errors.As(err, &responseErr)
			httpLog.Error = responseErr.Error()
			return c.JSON(http.StatusBadRequest, responseErr)
		}

		b, _ := json.Marshal(ctrl.param)
		httpLog.Request.Param = string(b)
	}

	result, err := ctrl.fn(ctx, ctrl.param)
	if err != nil {
		var responseErr *domain.Error
		if errors.As(err, &responseErr) {
			return c.JSON(responseErr.Status, responseErr)
		}

		httpLog.Error = err.Error()
		return c.JSON(http.StatusInternalServerError, err)
	}

	if result != nil {
		httpLog.Response.Body = result
		return c.JSON(ctrl.httpStatus, result)
	}

	return c.JSON(ctrl.httpStatus, nil)
}

func (ctrl *Handler) bind(c echo.Context) error {
	if err := ctrl.Bind(ctrl.param, c); err != nil {
		return &domain.Error{
			Type:   "bind_error",
			Status: http.StatusBadRequest,
			Detail: err.Error(),
		}
	}
	return nil
}

func (ctrl *Handler) validate() error {
	if err := ctrl.Struct(ctrl.param); err != nil {
		return &domain.Error{
			Type:   "validate_error",
			Status: http.StatusBadRequest,
			Detail: err.Error(),
		}
	}
	return nil
}
