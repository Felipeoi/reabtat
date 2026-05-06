package httpdinmates

import (
	domain "github.com/jeje-gab/resocializationV2/backend/internal/collections/inmates"
	"github.com/labstack/echo/v4"
)

func RegisterInmatesRoutes(api *echo.Group, uc domain.UseCase, jwtMW echo.MiddlewareFunc) {
	h := NewHandler(uc)
	// Aplica JWT middleware primeiro, depois o UserID middleware
	inmates := api.Group("/inmates", jwtMW, UserIDMiddleware())

	inmates.POST("", h.Create)
	inmates.GET("", h.List)
	inmates.GET("/:id", h.Get)
	inmates.PUT("/:id", h.Update)
	inmates.DELETE("/:id", h.Delete)
}
