package repository

import "github.com/Bottlist/bottlist/external/mysql"

type AuthRepository interface {
	A()
}

func NewAuthRepository(conn *mysql.Connector) AuthRepository {
	return &authRepository{conn: conn}
}

type authRepository struct {
	conn *mysql.Connector
}

func (a *authRepository) A() {

}
