package http

import (
	"github.com/jeje-gab/resocializationV2/backend/internal/collections/ufs"
	"github.com/labstack/echo/v4"
)

func RegisterUFRoutes(g *echo.Group, uc ufs.Usecase) {
	h := NewHandler(uc)

	g.GET("/ufs", h.List)
}
