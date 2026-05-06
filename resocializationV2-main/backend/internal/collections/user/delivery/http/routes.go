package httpduser

import (
	domain "github.com/jeje-gab/resocializationV2/backend/internal/collections/user"
	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(api *echo.Group, uc domain.Usecase, adminMW echo.MiddlewareFunc) {
	h := NewHandler(uc)
	// Rotas de users exigem role "admin"
	users := api.Group("/users", adminMW)

	users.POST("", h.Create)
	users.GET("", h.List)
	users.GET("/:id", h.Get)
	users.PUT("/:id", h.Update)
	users.DELETE("/:id", h.Delete)
}
