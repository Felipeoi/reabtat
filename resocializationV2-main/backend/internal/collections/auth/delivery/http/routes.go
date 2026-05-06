package httpd

import (
	domain "github.com/jeje-gab/resocializationV2/backend/internal/collections/auth"
	"github.com/labstack/echo/v4"
)

func RegisterAuthRoutes(api *echo.Group, uc domain.Usecase, jwtMW echo.MiddlewareFunc) {
	h := NewHandler(uc)
	auth := api.Group("/auth")
	auth.POST("/signup", h.Signup)
	auth.POST("/login", h.Login)
	auth.GET("/me", h.Me, jwtMW) // Rota protegida para obter dados do usuário logado
}
