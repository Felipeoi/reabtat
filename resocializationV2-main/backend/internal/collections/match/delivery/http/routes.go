package http

import (
	"github.com/jeje-gab/resocializationV2/backend/internal/collections/match"
	"github.com/labstack/echo/v4"
)

func RegisterMatchRoutes(api *echo.Group, uc match.Usecase, jwtMW echo.MiddlewareFunc) {
	h := NewHandler(uc)

	// Aplica JWT middleware primeiro, depois o UserID middleware
	matches := api.Group("/matches", jwtMW, UserIDMiddleware())

	matches.GET("", h.FindMatches)
	matches.GET("/:myInmateId/:matchedInmateId", h.FindMatchByID)
}
