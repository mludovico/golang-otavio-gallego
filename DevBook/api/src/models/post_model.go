package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID          uint64    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	AuthorNick  string    `json:"authorNick,omitempty"`
	AuthorID    uint64    `json:"author_id,omitempty"`
	Likes       uint64    `json:"likes"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdateAt    time.Time `json:"updated_at,omitempty"`
	ILiked      bool      `json:"i_liked"`
}

func (p *Post) Prepare() error {
	if err := p.validate(); err != nil {
		return err
	}
	p.format()
	return nil
}

func (p *Post) validate() error {
	if p.Title == "" {
		return errors.New("title is required and cannot be empty")
	}

	if p.Description == "" {
		return errors.New("description is required and cannot be empty")
	}

	if p.AuthorID == 0 {
		return errors.New("author_id is required and cannot be empty")
	}

	return nil
}

func (p *Post) format() {
	p.Title = strings.TrimSpace(p.Title)
	p.Description = strings.TrimSpace(p.Description)
}
