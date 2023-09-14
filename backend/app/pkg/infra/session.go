package infra

import (
	"context"
	redisClient "github.com/Bottlist/bottlist/external/redis"
	"github.com/Bottlist/bottlist/pkg/domain/repository"
	"time"
)

type SessionRepositoryImpl struct {
	client *redisClient.Connector
}

func NewSessionRepository(client *redisClient.Connector) repository.SessionRepository {
	return &SessionRepositoryImpl{
		client: client,
	}
}
func (r *SessionRepositoryImpl) SetSession(ctx context.Context, sessionId string, userId string, ttl time.Duration) error {
	err := r.client.Conn.Set(ctx, sessionId, userId, ttl).Err()
	return err
}
func (r *SessionRepositoryImpl) GetSession(ctx context.Context, sessionId string) (string, error) {
	val, err := r.client.Conn.Get(ctx, sessionId).Result()
	return val, err
}
func (r *SessionRepositoryImpl) DeleteSession(ctx context.Context, sessionId string) error {
	err := r.client.Conn.Del(ctx, sessionId).Err()
	return err
}
