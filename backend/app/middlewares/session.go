package middlewares

import (
	"github.com/Bottlist/bottlist/pkg/domain/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SessionMiddleware struct {
	client repository.IFSessionRepository
}

func NewSessionMiddleware(client *repository.SessionRepository) *SessionMiddleware {
	return &SessionMiddleware{
		client: client,
	}
}

func (s *SessionMiddleware) Session(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		cookie, err := c.Cookie("session_id")
		if err != nil {
			if err != http.ErrNoCookie {
				return echo.NewHTTPError(http.StatusUnauthorized, err)
			} else {
				return next(c)
			}
		} else {
			user, err := s.client.GetSession(ctx, cookie.Value)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err)
			}
			c.Set("user_id", user.ID)
			c.Set("user_type", user.UserType)
		}
		return next(c)
	}
}
