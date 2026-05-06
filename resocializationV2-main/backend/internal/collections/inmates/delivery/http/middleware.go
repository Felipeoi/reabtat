package httpdinmates

import (
	"net/http"
	"strconv"

	"github.com/jeje-gab/resocializationV2/backend/internal/collections/inmates/usecase"
	"github.com/labstack/echo/v4"
)

// UserIDMiddleware extrai o user_id do echo.Context (setado pelo JWT middleware)
// e injeta no context.Context para que os usecases possam acessá-lo
func UserIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extrai o userID do echo.Context (setado pelo middleware JWT)
			userIDStr, ok := c.Get("userID").(string)
			if !ok || userIDStr == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "user not authenticated")
			}

			userID, err := strconv.ParseInt(userIDStr, 10, 64)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid user id")
			}

			// Injeta o userID no context.Context
			ctx := usecase.WithUserID(c.Request().Context(), userID)

			// Atualiza o request com o novo context
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
