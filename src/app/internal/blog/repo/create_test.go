package blog_repo

import (
	"blog/domain"
	"blog/ent"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBlogCreateUsecase(t *testing.T) {
	t.Skip()
	r := require.New(t)

	r.Nil(nil)

	b := domain.Blog{}
	ctx := context.Background()
	client := &ent.Client{}
	domain.CreateBlog(b, domain.CreateBlogCfg{
		RepoCreate: CreateBlog(ctx, client),
	})

}
