package middlewares

import (
	"github.com/Bottlist/bottlist/pkg/domain/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SessionMiddleware struct {
	client repository.SessionRepository
}

func NewSessionMiddleware(client repository.SessionRepository) *SessionMiddleware {
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
				return echo.ErrUnauthorized
			} else {
				return next(c)
			}
		} else {
			userId, err := s.client.GetSession(ctx, cookie.Value)
			if err != nil {
				return echo.ErrUnauthorized
			}
			c.Set("user_id", userId)
		}
		return next(c)
	}
}
