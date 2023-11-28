package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type IFSessionRepository interface {
	SetSession(ctx context.Context, sessionId string, userId string, ttl time.Duration) error
	GetSession(ctx context.Context, sessionId string) (string, error)
	DeleteSession(ctx context.Context, sessionId string) error
}

type SessionRepository struct {
	client *redis.Client
}

func NewSessionRepository(client *redis.Client) *SessionRepository {
	return &SessionRepository{
		client: client,
	}
}
func (r *SessionRepository) SetSession(ctx context.Context, sessionId string, userId string, ttl time.Duration) error {
	err := r.client.Set(ctx, sessionId, userId, ttl).Err()
	return err
}
func (r *SessionRepository) GetSession(ctx context.Context, sessionId string) (string, error) {
	val, err := r.client.Get(ctx, sessionId).Result()
	return val, err
}
func (r *SessionRepository) DeleteSession(ctx context.Context, sessionId string) error {
	err := r.client.Del(ctx, sessionId).Err()
	return err
}
