package jwt

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Service struct {
	secret  []byte
	expires time.Duration
}

func NewJWT(secret string, expires time.Duration) *Service {
	return &Service{secret: []byte(secret), expires: expires}
}

// Generate cria um token JWT com userID e role
func (s *Service) Generate(userID, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"role": role,
		"exp":  time.Now().Add(s.expires).Unix(),
		"iat":  time.Now().Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(s.secret)
}

// parseToken valida e extrai claims do token
func (s *Service) parseToken(tokenStr string) (jwt.MapClaims, error) {
	tok, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, "Algoritmo de assinatura do token inválido.")
		}
		return s.secret, nil
	})
	if err != nil || !tok.Valid {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Token inválido ou expirado.")
	}
	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Token com formato inválido.")
	}
	return claims, nil
}

// Middleware de autenticação - verifica apenas se o token é válido
func (s *Service) Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")
			if len(auth) < 8 || auth[:7] != "Bearer " {
				return echo.NewHTTPError(http.StatusUnauthorized, "Envie o header Authorization: Bearer <token>.")
			}
			tokenStr := auth[7:]
			claims, err := s.parseToken(tokenStr)
			if err != nil {
				return err
			}
			// Armazena userID e role no contexto
			if sub, _ := claims["sub"].(string); sub != "" {
				c.Set("userID", sub)
			}
			if role, _ := claims["role"].(string); role != "" {
				c.Set("userRole", role)
			}
			return next(c)
		}
	}
}

// RequireRole retorna middleware que verifica se o usuário tem um dos roles permitidos
func (s *Service) RequireRole(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Primeiro executa o middleware de autenticação
			auth := c.Request().Header.Get("Authorization")
			if len(auth) < 8 || auth[:7] != "Bearer " {
				return echo.NewHTTPError(http.StatusUnauthorized, "Envie o header Authorization: Bearer <token>.")
			}
			tokenStr := auth[7:]
			claims, err := s.parseToken(tokenStr)
			if err != nil {
				return err
			}

			// Armazena userID e role no contexto
			if sub, _ := claims["sub"].(string); sub != "" {
				c.Set("userID", sub)
			}

			userRole, _ := claims["role"].(string)
			c.Set("userRole", userRole)

			// Verifica se o role do usuário está na lista de permitidos
			for _, allowedRole := range allowedRoles {
				if userRole == allowedRole {
					return next(c)
				}
			}

			return echo.NewHTTPError(http.StatusForbidden, "Sem permissão para esta operação.")
		}
	}
}
