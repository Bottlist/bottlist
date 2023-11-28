package repository

import "github.com/jmoiron/sqlx"

type IFUserRepository interface {
}

func NewUserRepository(conn *sqlx.DB) *UserRepository {
	return &UserRepository{conn: conn}
}

type UserRepository struct {
	conn *sqlx.DB
}
