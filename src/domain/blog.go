package domain

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Blog struct {
	id        int
	title     string
	slug      Slug
	content   Content
	createdAt time.Time
}

func New(id int, title string, content Content, slug Slug) (*Blog, error) {

	b := Blog{
		id:        id,
		title:     title,
		slug:      slug,
		content:   content,
		createdAt: time.Now(),
	}
	if err := b.validate(); err != nil {
		return nil, err
	}
	return &b, nil
}

func (b Blog) validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.id, validation.Required),
		validation.Field(&b.title, validation.Required, validation.Length(4, 50)),
		validation.Field(&b.content, validation.Required),
		validation.Field(&b.createdAt, validation.Required),
	)
}
