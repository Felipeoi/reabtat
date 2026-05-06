package http

import (
	"net/http"

	"github.com/jeje-gab/resocializationV2/backend/internal/collections/cities"
	"github.com/jeje-gab/resocializationV2/backend/pkg/pgxerr"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	uc cities.Usecase
}

func NewHandler(uc cities.Usecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) List(c echo.Context) error {
	ctx := c.Request().Context()
	ufCode := c.QueryParam("uf_code")

	cities, err := h.uc.List(ctx, ufCode)
	if err != nil {
		if pgxerr.IsUndefinedRelation(err) {
			return c.JSON(http.StatusServiceUnavailable, map[string]string{
				"error": "Tabelas ainda não criadas. Na raiz do projeto: sh scripts/apply-migrations-docker.sh",
			})
		}
		c.Logger().Error("cities list: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao listar cidades",
		})
	}

	return c.JSON(http.StatusOK, cities)
}
