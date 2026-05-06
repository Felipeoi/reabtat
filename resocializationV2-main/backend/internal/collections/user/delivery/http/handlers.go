package httpduser

import (
	"net/http"
	"strconv"

	domain "github.com/jeje-gab/resocializationV2/backend/internal/collections/user"
	"github.com/labstack/echo/v4"
)

type Handler struct{ uc domain.Usecase }

func NewHandler(uc domain.Usecase) *Handler { return &Handler{uc: uc} }

func (h *Handler) Create(c echo.Context) error {
	var in struct{ Name, Email, Password string }
	if err := c.Bind(&in); err != nil || in.Name == "" || in.Email == "" || in.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}
	id, err := h.uc.Create(in.Name, in.Email, in.Password)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, map[string]any{"id": id})
}

func (h *Handler) Get(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	u, err := h.uc.Get(id64)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	return c.JSON(http.StatusOK, u)
}

func (h *Handler) List(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	if limit <= 0 {
		limit = 20
	}
	us, err := h.uc.List(limit, offset)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, us)
}

func (h *Handler) Update(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	var in struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Status   string `json:"status"`
		Password string `json:"password"`
	}
	if err := c.Bind(&in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}

	// Se não forneceu status, mantém como ativo (padrão)
	if in.Status == "" {
		in.Status = "ativo"
	}

	// Valida status
	if in.Status != "ativo" && in.Status != "inativo" {
		return echo.NewHTTPError(http.StatusBadRequest, "status must be 'ativo' or 'inativo'")
	}

	if err := h.uc.UpdateByAdmin(id64, in.Name, in.Email, in.Status, in.Password); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) Delete(c echo.Context) error {
	id64, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	if err := h.uc.Delete(id64); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	return c.NoContent(http.StatusNoContent)
}
