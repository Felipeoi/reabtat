package http

import (
	"github.com/jeje-gab/resocializationV2/backend/internal/collections/prisonunits"
	"github.com/labstack/echo/v4"
)

func RegisterPrisonUnitRoutes(g *echo.Group, uc prisonunits.Usecase) {
	h := NewHandler(uc)
	g.GET("/prison-units", h.List)
}
