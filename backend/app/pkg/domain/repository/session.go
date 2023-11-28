package repository

import (
	"context"
	"encoding/json"
	"github.com/Bottlist/bottlist/pkg/types"
	"github.com/go-redis/redis/v8"
	"time"
)

type IFSessionRepository interface {
	SetSession(ctx context.Context, sessionId string, user types.User, ttl time.Duration) error
	GetSession(ctx context.Context, sessionId string) (*types.User, error)
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
func (r *SessionRepository) SetSession(ctx context.Context, sessionId string, user types.User, ttl time.Duration) error {
	userJson, _ := json.Marshal(user)
	err := r.client.Set(ctx, sessionId, userJson, ttl).Err()
	return err
}
func (r *SessionRepository) GetSession(ctx context.Context, sessionId string) (*types.User, error) {
	val, err := r.client.Get(ctx, sessionId).Result()
	if err != nil {
		return nil, err
	}
	var user = new(types.User)
	err = json.Unmarshal([]byte(val), user)
	if err != nil {
		return nil, err
	}
	return user, err
}
func (r *SessionRepository) DeleteSession(ctx context.Context, sessionId string) error {
	err := r.client.Del(ctx, sessionId).Err()
	return err
}
