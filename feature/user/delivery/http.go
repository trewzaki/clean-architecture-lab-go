package delivery

import (
	"clean-architecture-go/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler ..
type Handler struct {
	usecase domain.UserUsecase
}

// NewHandler ..
func NewHandler(e *echo.Group, u domain.UserUsecase) *Handler {
	h := Handler{usecase: u}

	e.GET("/users", h.List)

	return &h
}

// List ..
func (h *Handler) List(c echo.Context) error {
	result, err := h.usecase.ListUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": result})
}
