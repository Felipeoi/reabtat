package http

import (
	"net/http"

	"github.com/jeje-gab/resocializationV2/backend/internal/collections/ufs"
	"github.com/jeje-gab/resocializationV2/backend/pkg/pgxerr"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	uc ufs.Usecase
}

func NewHandler(uc ufs.Usecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) List(c echo.Context) error {
	ctx := c.Request().Context()

	ufs, err := h.uc.List(ctx)
	if err != nil {
		if pgxerr.IsUndefinedRelation(err) {
			return c.JSON(http.StatusServiceUnavailable, map[string]string{
				"error": "Tabelas ainda não criadas. Na raiz do projeto: sh scripts/apply-migrations-docker.sh",
			})
		}
		c.Logger().Error("ufs list: ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao listar UFs",
		})
	}

	return c.JSON(http.StatusOK, ufs)
}
