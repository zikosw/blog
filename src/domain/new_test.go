package domain_test

import (
	d "blog/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewInvalid(t *testing.T) {

	r := require.New(t)

	title := "title"
	content := d.Content("content")
	sl := newSlug(title)

	b, err := d.NewBlog(title, content, sl)
	r.NoError(err)
	r.NotNil(b)

}
