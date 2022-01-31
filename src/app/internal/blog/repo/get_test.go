package blog_repo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetFake(t *testing.T) {
	r := require.New(t)
	bs, err := ListFake()
	r.NoError(err)

	for _, b := range bs {
		fmt.Printf("\n\n\n")
		fmt.Println(b)
	}
}
