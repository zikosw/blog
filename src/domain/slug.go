// to keep tracking of every generated slug and make sure that every slug link is working
package domain

import (
	"time"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gosimple/slug"
)

type TargetType string

const (
	TargetBlog TargetType = "blog"
	TargetURL  TargetType = "url"
)

type Slug struct {
	slug      string
	createdAt time.Time
}

func (s Slug) Slug() string         { return s.slug }
func (s Slug) CreatedAt() time.Time { return s.createdAt }

func NewSlug(text string) (*Slug, error) {
	genSlug := slug.Make(text)
	s := Slug{
		slug:      genSlug,
		createdAt: time.Now(),
	}

	if err := s.validate(); err != nil {
		return nil, err
	}

	return &s, nil
}

func (s Slug) validate() error {
	return v.ValidateStruct(&s,
		v.Field(&s.slug, v.Required),
		v.Field(&s.createdAt, v.Required),
	)
}

func UnmarshalSlug(slug string, createdAt time.Time) (*Slug, error) {
	s := Slug{
		slug:      slug,
		createdAt: createdAt,
	}
	if err := s.validate(); err != nil {
		return nil, err
	}
	return &s, nil
}
