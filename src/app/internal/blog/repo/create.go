package blog_repo

import (
	"context"
	"fmt"
	"log"

	"blog/domain"
	"blog/ent"

	"entgo.io/ent/dialect/sql/sqlgraph"
)

func CreateBlog(ctx context.Context, client *ent.Client) func(domain.Blog) (*domain.Blog, error) {
	return func(b domain.Blog) (*domain.Blog, error) {

		eb, err := client.Blog.
			Create().SetTitle(b.Title()).
			SetContent(string(b.Content())).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed creating blog: %w", err)
		}
		log.Println("blog was created: ", eb)

		s := b.Slug()

		es, err := client.Slug.
			Create().
			SetSlug(s.Slug()).
			SetLink(eb).
			Save(ctx)
		if err != nil {
			if is := sqlgraph.IsUniqueConstraintError(err); is {
				// TODO: feature: handle when slug duplicated
				fmt.Println("is-uniq", is)
			}
			return nil, fmt.Errorf("failed creating slug: %w", err)
		}
		log.Println("sl was created: ", es)

		dbS, err := domain.UnmarshalSlug(es.Slug, es.CreatedAt)
		if err != nil {
			return nil, err
		}
		dbB, err := domain.UnmarshalBlog(eb.ID, eb.Title, domain.Content(eb.Content), *dbS, eb.CreatedAt)

		if err != nil {
			return nil, err
		}
		return dbB, nil
	}
}
