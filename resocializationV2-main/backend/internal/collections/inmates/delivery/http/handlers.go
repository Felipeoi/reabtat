package httpdinmates

import (
	"net/http"
	"strconv"

	domain "github.com/jeje-gab/resocializationV2/backend/internal/collections/inmates"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
	"github.com/jeje-gab/resocializationV2/backend/pkg/pgxerr"
	"github.com/labstack/echo/v4"
)

type Handler struct{ uc domain.UseCase }

func NewHandler(uc domain.UseCase) *Handler { return &Handler{uc: uc} }

func (h *Handler) Create(c echo.Context) error {
	var in entity.InmatesFlat
	if err := c.Bind(&in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}

	// Validações básicas
	if in.OriginID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "origin_id is required")
	}
	if in.Custody == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "custody is required")
	}
	if in.Custody != "CLOSED" && in.Custody != "SEMI_OPEN" && in.Custody != "OPEN" {
		return echo.NewHTTPError(http.StatusBadRequest, "custody must be CLOSED, SEMI_OPEN or OPEN")
	}
	if in.DestinationID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "destination_id is required")
	}

	// O userID já foi injetado no context pelo UserIDMiddleware
	err, _, inmate := h.uc.Create(c.Request().Context(), in)
	if err != nil {
		if pgxerr.IsUndefinedRelation(err) {
			return c.JSON(http.StatusServiceUnavailable, map[string]string{
				"error": "Tabelas ainda não criadas. Na raiz do projeto: sh scripts/apply-migrations-docker.sh",
			})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, inmate)
}

func (h *Handler) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	err, _, inmate := h.uc.ById(c.Request().Context(), id)
	if err != nil {
		if pgxerr.IsUndefinedRelation(err) {
			return c.JSON(http.StatusServiceUnavailable, map[string]string{
				"error": "Tabelas ainda não criadas. Na raiz do projeto: sh scripts/apply-migrations-docker.sh",
			})
		}
		return echo.NewHTTPError(http.StatusNotFound, "inmate not found")
	}

	return c.JSON(http.StatusOK, inmate)
}

func (h *Handler) List(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	if limit <= 0 {
		limit = 20
	}

	err, _, inmates := h.uc.List(c.Request().Context(), limit, offset)
	if err != nil {
		if pgxerr.IsUndefinedRelation(err) {
			return c.JSON(http.StatusServiceUnavailable, map[string]string{
				"error": "Tabelas ainda não criadas. Na raiz do projeto: sh scripts/apply-migrations-docker.sh",
			})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, inmates)
}

func (h *Handler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	var in entity.InmatesFlat
	if err := c.Bind(&in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}

	// Adiciona o ID ao objeto
	in.Id = &id

	// Validações básicas
	if in.OriginID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "origin_id is required")
	}
	if in.Custody == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "custody is required")
	}
	if in.Custody != "CLOSED" && in.Custody != "SEMI_OPEN" && in.Custody != "OPEN" {
		return echo.NewHTTPError(http.StatusBadRequest, "custody must be CLOSED, SEMI_OPEN or OPEN")
	}
	if in.DestinationID <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "destination_id is required")
	}

	err, _, inmate := h.uc.Update(c.Request().Context(), in)
	if err != nil {
		if pgxerr.IsUndefinedRelation(err) {
			return c.JSON(http.StatusServiceUnavailable, map[string]string{
				"error": "Tabelas ainda não criadas. Na raiz do projeto: sh scripts/apply-migrations-docker.sh",
			})
		}
		return echo.NewHTTPError(http.StatusNotFound, "inmate not found")
	}

	return c.JSON(http.StatusOK, inmate)
}

func (h *Handler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	err, _ = h.uc.Delete(c.Request().Context(), id)
	if err != nil {
		if pgxerr.IsUndefinedRelation(err) {
			return c.JSON(http.StatusServiceUnavailable, map[string]string{
				"error": "Tabelas ainda não criadas. Na raiz do projeto: sh scripts/apply-migrations-docker.sh",
			})
		}
		return echo.NewHTTPError(http.StatusNotFound, "inmate not found")
	}

	return c.NoContent(http.StatusNoContent)
}
