package httpd

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	domain "github.com/jeje-gab/resocializationV2/backend/internal/collections/auth"
	"github.com/labstack/echo/v4"
)

type Handler struct{ uc domain.Usecase }

func NewHandler(uc domain.Usecase) *Handler { return &Handler{uc: uc} }

func (h *Handler) Signup(c echo.Context) error {
	var in struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Pass     string `json:"password"`
		Telefone string `json:"telefone"`
	}
	if err := c.Bind(&in); err != nil {
		c.Logger().Error("signup: bind error:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Corpo da requisição inválido."})
	}
	in.Name = strings.TrimSpace(in.Name)
	in.Email = strings.ToLower(strings.TrimSpace(in.Email))
	in.Telefone = strings.TrimSpace(in.Telefone)
	if in.Name == "" || in.Email == "" || in.Pass == "" {
		c.Logger().Warn("signup: missing required fields")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nome, e-mail e senha são obrigatórios."})
	}
	c.Logger().Info("signup: creating user", "email", in.Email)
	id, err := h.uc.Signup(in.Name, in.Email, in.Pass, in.Telefone)
	if err != nil {
		c.Logger().Error("signup: usecase error:", err)
		msg := err.Error()
		if strings.Contains(msg, "duplicate key") || strings.Contains(msg, "23505") {
			msg = "Este e-mail já está cadastrado."
		}
		return c.JSON(http.StatusConflict, map[string]string{"error": msg})
	}
	c.Logger().Info("signup: user created successfully", "id", id)
	return c.JSON(http.StatusCreated, map[string]any{"user_id": id})
}

func (h *Handler) Login(c echo.Context) error {
	var in struct {
		Email string `json:"email"`
		Pass  string `json:"password"`
	}
	if err := c.Bind(&in); err != nil || strings.TrimSpace(in.Email) == "" || in.Pass == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "E-mail e senha são obrigatórios."})
	}
	in.Email = strings.ToLower(strings.TrimSpace(in.Email))
	tok, err := h.uc.Login(in.Email, in.Pass)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrInactiveAccount):
			return c.JSON(http.StatusForbidden, map[string]string{"error": "Conta inativa. Procure o administrador."})
		case errors.Is(err, domain.ErrInvalidLogin):
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "E-mail ou senha incorretos."})
		default:
			c.Logger().Error("login: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Falha ao autenticar. Tente novamente."})
		}
	}
	return c.JSON(http.StatusOK, map[string]any{"token": tok})
}

func (h *Handler) Me(c echo.Context) error {
	// Pega o userID do contexto (inserido pelo middleware JWT)
	userIDStr, ok := c.Get("userID").(string)
	if !ok || userIDStr == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Não autenticado."})
	}

	// Converte string para int64
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Identificador de usuário inválido."})
	}

	user, err := h.uc.Me(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Usuário não encontrado."})
	}

	return c.JSON(http.StatusOK, user)
}
