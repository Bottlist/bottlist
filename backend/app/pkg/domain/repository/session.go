package repository

import (
	"context"
	"time"
)

type SessionRepository interface {
	SetSession(ctx context.Context, sessionId string, userId string, ttl time.Duration) error
	GetSession(ctx context.Context, sessionId string) (string, error)
	DeleteSession(ctx context.Context, sessionId string) error
}
