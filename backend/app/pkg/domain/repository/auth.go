package repository

import (
	"github.com/Bottlist/bottlist/external/mysql"
	"github.com/Bottlist/bottlist/pkg/domain/model"
)

type AuthRepository interface {
	GetUserByEmail(email string) (*model.User, error)
	GetProvisionalUserByEmail(email string) (*model.ProvisionalUser, error)
	InsertProvisionalUser(user *model.CreateProvisionalUser) error
}

func NewAuthRepository(conn *mysql.Connector) AuthRepository {
	return &authRepository{conn: conn}
}

type authRepository struct {
	conn *mysql.Connector
}

func (a *authRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	query := `SELECT * FROM users WHERE email = ?`
	err := a.conn.Conn.Get(&user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *authRepository) GetProvisionalUserByEmail(email string) (*model.ProvisionalUser, error) {
	var user model.ProvisionalUser
	query := `SELECT * FROM provisional_users WHERE email = ?`
	err := a.conn.Conn.Get(&user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *authRepository) InsertProvisionalUser(user *model.CreateProvisionalUser) error {
	query := `INSERT INTO provisional_users (email, first_name, last_name, first_name_hira, last_name_hira, screen_name, birthday, password, token, expired_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := a.conn.Conn.Exec(query, user.Email, user.FirstName, user.LastName, user.FirstNameHira, user.LastNameHira, user.ScreenName, user.Birthday.GetDate(), user.Password, user.Token, user.ExpiredAt)
	return err
}
