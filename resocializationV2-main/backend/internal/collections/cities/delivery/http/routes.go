package http

import (
	"github.com/jeje-gab/resocializationV2/backend/internal/collections/cities"
	"github.com/labstack/echo/v4"
)

func RegisterCityRoutes(g *echo.Group, uc cities.Usecase) {
	h := NewHandler(uc)

	g.GET("/cities", h.List)
}
