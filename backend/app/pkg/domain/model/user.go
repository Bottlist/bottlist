package model

import (
	"time"

	"github.com/Bottlist/bottlist/pkg/types"
)

type UserBase struct {
	ID int64 `json:"id"`
}

type User struct {
	ID            int        `db:"id" json:"id"`
	Email         string     `db:"email" json:"email"`
	FirstName     string     `db:"first_name" json:"first_name"`
	LastName      string     `db:"last_name" json:"last_name"`
	FirstNameHira string     `db:"first_name_hira" json:"first_name_hira"`
	LastNameHira  string     `db:"last_name_hira" json:"last_name_hira"`
	ScreenName    string     `db:"screen_name" json:"screen_name"`
	Birthday      types.Date `db:"birthday" json:"birthday"`
	Password      string     `db:"password" json:"-"`
	Image         string     `db:"image" json:"image"`
}

type UserCreate struct {
	Email         string     `db:"email" json:"email"`
	FirstName     string     `db:"first_name" json:"first_name"`
	LastName      string     `db:"last_name" json:"last_name"`
	FirstNameHira string     `db:"first_name_hira" json:"first_name_hira"`
	LastNameHira  string     `db:"last_name_hira" json:"last_name_hira"`
	ScreenName    string     `db:"screen_name" json:"screen_name"`
	Birthday      types.Date `db:"birthday" json:"birthday"`
	Password      string     `db:"password" json:"-"`
	Image         string     `db:"image" json:"image"`
}

type ProvisionalUser struct {
	ID            int        `db:"id" json:"id"`
	Email         string     `db:"email" json:"email"`
	FirstName     string     `db:"first_name" json:"first_name"`
	LastName      string     `db:"last_name" json:"last_name"`
	FirstNameHira string     `db:"first_name_hira" json:"first_name_hira"`
	LastNameHira  string     `db:"last_name_hira" json:"last_name_hira"`
	ScreenName    string     `db:"screen_name" json:"screen_name"`
	Birthday      types.Date `db:"birthday" json:"birthday"`
	Password      string     `db:"password" json:"-"`
	Token         string     `db:"token" json:"token"`
	CreatedAt     time.Time  `db:"created_at" json:"created_at"`
	ExpiredAt     time.Time  `db:"expired_at" json:"expired_at"`
}

type CreateProvisionalUser struct {
	Email         string     `db:"email" json:"email"`
	FirstName     string     `db:"first_name" json:"first_name"`
	LastName      string     `db:"last_name" json:"last_name"`
	FirstNameHira string     `db:"first_name_hira" json:"first_name_hira"`
	LastNameHira  string     `db:"last_name_hira" json:"last_name_hira"`
	ScreenName    string     `db:"screen_name" json:"screen_name"`
	Birthday      types.Date `db:"birthday" json:"birthday"`
	Password      string     `db:"password" json:"-"`
	Token         string     `db:"token" json:"token"`
	ExpiredAt     time.Time  `db:"expired_at" json:"expired_at"`
}
