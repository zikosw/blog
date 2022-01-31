package domain

import (
	"time"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

type Blog struct {
	id        int
	title     string
	slug      Slug
	content   Content
	createdAt time.Time
}

func NewBlog(title string, content Content, slug Slug) (*Blog, error) {

	b := Blog{
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
	return v.ValidateStruct(&b,
		v.Field(&b.id, v.Min(0)), // validation.Required),
		v.Field(&b.title, v.Required, v.Length(4, 50)),
		v.Field(&b.content, v.Required),
		v.Field(&b.createdAt, v.Required),
	)
}

func (b Blog) ID() int              { return b.id }
func (b Blog) Title() string        { return b.title }
func (b Blog) Slug() Slug           { return b.slug }
func (b Blog) Content() Content     { return b.content }
func (b Blog) CreatedAt() time.Time { return b.createdAt }

func UnmarshalBlog(id int, title string, content Content, slug Slug, createdAt time.Time) (*Blog, error) {
	b := Blog{
		id:        id,
		title:     title,
		slug:      slug,
		content:   content,
		createdAt: createdAt,
	}

	if err := b.validate(); err != nil {
		return nil, err
	}

	return &b, nil
}
