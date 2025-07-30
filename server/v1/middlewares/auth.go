package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// RequireSession ensures a session cookie is present before allowing access.
func RequireSession() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, err := c.Cookie("session")
			if err != nil || sess.Value == "" {
				return c.Redirect(http.StatusFound, "/acessar")
			}
			return next(c)
		}
	}
}
