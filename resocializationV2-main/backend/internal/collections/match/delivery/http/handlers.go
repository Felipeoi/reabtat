package http

import (
	"net/http"
	"strconv"

	"github.com/jeje-gab/resocializationV2/backend/internal/collections/match"
	"github.com/jeje-gab/resocializationV2/backend/pkg/pgxerr"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	uc match.Usecase
}

func NewHandler(uc match.Usecase) *Handler {
	return &Handler{uc: uc}
}

// FindMatches busca todos os matches possíveis para os inmates do usuário logado
func (h *Handler) FindMatches(c echo.Context) error {
	matches, err := h.uc.FindMatches(c.Request().Context())
	if err != nil {
		if pgxerr.IsUndefinedRelation(err) {
			return c.JSON(http.StatusServiceUnavailable, map[string]string{
				"error": "Tabelas ainda não criadas. Na raiz do projeto: sh scripts/apply-migrations-docker.sh",
			})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Erro ao buscar matches")
	}

	return c.JSON(http.StatusOK, matches)
}

// FindMatchByID busca um match específico pelos IDs dos inmates
func (h *Handler) FindMatchByID(c echo.Context) error {
	myInmateIDStr := c.Param("myInmateId")
	matchedInmateIDStr := c.Param("matchedInmateId")

	myInmateID, err := strconv.Atoi(myInmateIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID do seu inmate inválido")
	}

	matchedInmateID, err := strconv.Atoi(matchedInmateIDStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "ID do inmate matched inválido")
	}

	match, err := h.uc.FindMatchByID(c.Request().Context(), myInmateID, matchedInmateID)
	if err != nil {
		if pgxerr.IsUndefinedRelation(err) {
			return c.JSON(http.StatusServiceUnavailable, map[string]string{
				"error": "Tabelas ainda não criadas. Na raiz do projeto: sh scripts/apply-migrations-docker.sh",
			})
		}
		return echo.NewHTTPError(http.StatusNotFound, "Match não encontrado")
	}

	return c.JSON(http.StatusOK, match)
}
