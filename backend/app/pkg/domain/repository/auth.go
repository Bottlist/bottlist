package repository

import (
	"github.com/Bottlist/bottlist/pkg/domain/model"
	"github.com/jmoiron/sqlx"
)

type IFAuthRepository interface {
	InsertUser(user *model.UserCreate) error
	GetUserByEmail(email string) (*model.User, error)
	InsertProvisionalUser(user *model.CreateProvisionalUser) error
	GetProvisionalUserByEmail(email string) (*model.ProvisionalUser, error)
	GetProvisionalUserByToken(token string) (*model.ProvisionalUser, error)
	DeleteProvisionalUser(id int) error
}

func NewAuthRepository(conn *sqlx.DB) *AuthRepository {
	return &AuthRepository{conn: conn}
}

type AuthRepository struct {
	conn *sqlx.DB
}

func (a *AuthRepository) InsertUser(user *model.UserCreate) error {
	query := `INSERT INTO users (email, first_name, last_name, first_name_hira, last_name_hira, screen_name, birthday, password, image) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := a.conn.Exec(query, user.Email, user.FirstName, user.LastName, user.FirstNameHira, user.LastNameHira, user.ScreenName, user.Birthday.GetDate(), user.Password, user.Image)
	return err
}

func (a *AuthRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	query := `SELECT * FROM users WHERE email = ?`
	err := a.conn.Get(&user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *AuthRepository) InsertProvisionalUser(user *model.CreateProvisionalUser) error {
	query := `INSERT INTO provisional_users (email, first_name, last_name, first_name_hira, last_name_hira, screen_name, birthday, password, token, expired_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := a.conn.Exec(query, user.Email, user.FirstName, user.LastName, user.FirstNameHira, user.LastNameHira, user.ScreenName, user.Birthday.GetDate(), user.Password, user.Token, user.ExpiredAt)
	return err
}

func (a *AuthRepository) GetProvisionalUserByEmail(email string) (*model.ProvisionalUser, error) {
	var user model.ProvisionalUser
	query := `SELECT * FROM provisional_users WHERE email = ?`
	err := a.conn.Get(&user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *AuthRepository) GetProvisionalUserByToken(token string) (*model.ProvisionalUser, error) {
	var user model.ProvisionalUser
	query := `SELECT * FROM provisional_users WHERE token = ?`
	err := a.conn.Get(&user, query, token)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *AuthRepository) DeleteProvisionalUser(id int) error {
	query := `DELETE FROM provisional_users WHERE id = ?`
	_, err := a.conn.Exec(query, id)
	return err
}
