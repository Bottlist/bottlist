package repository

import (
	"github.com/Bottlist/bottlist/external/mysql"
	"github.com/Bottlist/bottlist/pkg/domain/model"
)

type AuthRepository interface {
	InsertUser(user *model.UserCreate) error
	GetUserByEmail(email string) (*model.User, error)
	InsertProvisionalUser(user *model.CreateProvisionalUser) error
	GetProvisionalUserByEmail(email string) (*model.ProvisionalUser, error)
	GetProvisionalUserByToken(token string) (*model.ProvisionalUser, error)
	DeleteProvisionalUser(id int) error
}

func NewAuthRepository(conn *mysql.Connector) AuthRepository {
	return &authRepository{conn: conn}
}

type authRepository struct {
	conn *mysql.Connector
}

func (a *authRepository) InsertUser(user *model.UserCreate) error {
	query := `INSERT INTO users (email, first_name, last_name, first_name_hira, last_name_hira, screen_name, birthday, password, image) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := a.conn.Conn.Exec(query, user.Email, user.FirstName, user.LastName, user.FirstNameHira, user.LastNameHira, user.ScreenName, user.Birthday.GetDate(), user.Password, user.Image)
	return err
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

func (a *authRepository) InsertProvisionalUser(user *model.CreateProvisionalUser) error {
	query := `INSERT INTO provisional_users (email, first_name, last_name, first_name_hira, last_name_hira, screen_name, birthday, password, token, expired_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := a.conn.Conn.Exec(query, user.Email, user.FirstName, user.LastName, user.FirstNameHira, user.LastNameHira, user.ScreenName, user.Birthday.GetDate(), user.Password, user.Token, user.ExpiredAt)
	return err
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

func (a *authRepository) GetProvisionalUserByToken(token string) (*model.ProvisionalUser, error) {
	var user model.ProvisionalUser
	query := `SELECT * FROM provisional_users WHERE token = ?`
	err := a.conn.Conn.Get(&user, query, token)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *authRepository) DeleteProvisionalUser(id int) error {
	query := `DELETE FROM provisional_users WHERE id = ?`
	_, err := a.conn.Conn.Exec(query, id)
	return err
}
