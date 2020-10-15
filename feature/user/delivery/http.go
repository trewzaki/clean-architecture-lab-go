package delivery

import (
	"clean-architecture-go/domain"

	"github.com/labstack/echo/v4"
)

// Handler ..
type Handler struct {
	usecase domain.UserUsecase
}

// NewHandler ..
func NewHandler(e *echo.Group, u domain.UserUsecase) *Handler {
	h := Handler{usecase: u}

	return &h
}
