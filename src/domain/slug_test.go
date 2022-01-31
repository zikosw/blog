package domain_test

import (
	d "blog/domain"

	"github.com/pkg/errors"
)

func newSlug(title string) d.Slug {
	s, err := d.NewSlug(title, d.TargetBlog, "999")
	if err != nil {
		panic(errors.Wrap(err, "new slug"))
	}
	return *s
}
