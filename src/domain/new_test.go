package domain_test

import (
	d "blog/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewInvalid(t *testing.T) {

	r := require.New(t)

	b, err := d.New(1, "title", "content")
	r.NoError(err)
	r.NotNil(b)

}
