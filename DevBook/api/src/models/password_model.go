package models

import "errors"

type Password struct {
	New          string `json:"new,omitempty"`
	Current      string `json:"current,omitempty"`
	Confirmation string `json:"confirmation,omitempty"`
}

func (p *Password) Prepare() error {
	if err := p.validate(); err != nil {
		return err
	}
	return nil
}

func (p Password) validate() error {
	if p.New == "" {
		return ErrNewPasswordIsRequired
	}
	if p.Current == "" {
		return ErrCurrentPasswordIsRequired
	}
	if p.Confirmation == "" {
		return ErrPasswordConfirmationIsRequired
	}
	if p.New != p.Confirmation {
		return ErrPasswordDoesNotMatchConfirmation
	}
	return nil
}

var (
	ErrNewPasswordIsRequired            = errors.New("new password is required and cannot be empty")
	ErrCurrentPasswordIsRequired        = errors.New("current password is required and cannot be empty")
	ErrPasswordConfirmationIsRequired   = errors.New("password confirmation is required and cannot be empty")
	ErrPasswordDoesNotMatchConfirmation = errors.New("new password and confirmation do not match")
)
