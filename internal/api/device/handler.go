package device

import "github.com/labstack/echo/v4"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Create(c echo.Context) error {
	return nil
}

func (h *Handler) Update(c echo.Context) error {
	return nil
}

func (h *Handler) GetAll(c echo.Context) error {
	return nil
}

func (h *Handler) GetById(c echo.Context) error {
	return nil
}

func (h *Handler) GetByBrand(c echo.Context) error {
	return nil
}

func (h *Handler) GetByState(c echo.Context) error {
	return nil
}

func (h *Handler) Delete(c echo.Context) error {
	return nil
}
