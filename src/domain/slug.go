// to keep tracking of every generated slug and make sure that every slug link is working
package domain

import (
	"fmt"
	"time"

	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gosimple/slug"
)

type SlugTarget string

const (
	TargetBlog SlugTarget = "blog"
)

type Slug struct {
	slug      string
	target    SlugTarget
	targetID  string
	createdAt time.Time
}

func NewSlug(text string, target SlugTarget, targetID string) (*Slug, error) {
	genSlug := slug.Make(fmt.Sprintf("%s-%s", text, targetID))
	s := Slug{
		slug:      genSlug,
		target:    target,
		targetID:  targetID,
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
		v.Field(&s.target, v.Required, v.In(TargetBlog)),
		v.Field(&s.targetID, v.Required),
		v.Field(&s.createdAt, v.Required),
	)
}
