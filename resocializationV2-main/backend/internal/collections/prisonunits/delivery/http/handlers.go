package http

import (
	"net/http"

	"github.com/jeje-gab/resocializationV2/backend/internal/collections/prisonunits"
	"github.com/jeje-gab/resocializationV2/backend/pkg/pgxerr"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	uc prisonunits.Usecase
}

func NewHandler(uc prisonunits.Usecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) List(c echo.Context) error {
	ctx := c.Request().Context()
	ufCode := c.QueryParam("uf_code")

	units, err := h.uc.List(ctx, ufCode)
	if err != nil {
		if pgxerr.IsUndefinedRelation(err) {
			return c.JSON(http.StatusServiceUnavailable, map[string]string{
				"error": "Tabelas ainda não criadas. Rode as migrations de unidades prisionais.",
			})
		}
		c.Logger().Error("prison units list: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao listar unidades prisionais",
		})
	}

	return c.JSON(http.StatusOK, units)
}
