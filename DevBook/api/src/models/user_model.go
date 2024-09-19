package models

import (
	"devbook_api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

const (
	InsertPreparation = true
	UpdatePreparation = false
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (u *User) Prepare(preparationType bool) error {
	if u.sanitize(preparationType) != nil {
		return ErrFailedToHashPassword
	}
	if err := u.validate(preparationType); err != nil {
		return err
	}
	return nil
}

func (u User) validate(shouldValidatePassword bool) error {
	if u.Name == "" {
		return ErrNameIsRequired
	}
	if u.Nick == "" {
		return ErrNickIsRequired
	}
	if u.Email == "" {
		return ErrEmailIsRequired
	}
	if checkmail.ValidateFormat(u.Email) != nil {
		return ErrEmailInvalidFormat
	}
	if shouldValidatePassword && u.Password == "" {
		return ErrPasswordIsRequired
	}
	return nil
}

func (u *User) sanitize(shouldHashPassword bool) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
	if shouldHashPassword {
		hashedPassword, err := security.Hash(u.Password)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

var (
	ErrNameIsRequired       = errors.New("name is required and cannot be empty")
	ErrNickIsRequired       = errors.New("nick is required and cannot be empty")
	ErrEmailIsRequired      = errors.New("email is required and cannot be empty")
	ErrEmailInvalidFormat   = errors.New("email is invalid")
	ErrPasswordIsRequired   = errors.New("password is required and cannot be empty")
	ErrFailedToHashPassword = errors.New("failed to hash password")
)
