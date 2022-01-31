package blog_repo

import (
	"blog/domain"
	"blog/ent"
	"context"

	gfi "github.com/brianvoe/gofakeit/v6"
)

func Get(ctx context.Context, client *ent.Client) func(id int) (*domain.Blog, error) {
	return func(id int) (*domain.Blog, error) {

		eb, err := client.Blog.Get(ctx, id)
		if err != nil {
			return nil, err
		}

		es, err := eb.QuerySlugs().First(ctx)
		if err != nil {
			return nil, err
		}

		s, err := domain.UnmarshalSlug(es.Slug, es.CreatedAt)
		if err != nil {
			return nil, err
		}

		b, err := domain.UnmarshalBlog(eb.ID, eb.Title, domain.Content(eb.Content), *s, eb.CreatedAt)
		if err != nil {
			return nil, err
		}

		return b, nil
	}
}

func GetFake() domain.Blog {
	id := gfi.IntRange(1, 10000)
	title := gfi.LoremIpsumSentence(3)
	content := gfi.LoremIpsumParagraph(5, 10, 10, "\n")
	at := gfi.Date()

	s, err := domain.NewSlug(title)
	if err != nil {
		panic(err)
	}
	b, err := domain.UnmarshalBlog(
		id,
		title,
		domain.Content(content),
		*s,
		at,
	)
	if err != nil {
		panic(err)
	}
	return *b
}

func ListFake() ([]domain.Blog, error) {
	bs := []domain.Blog{}
	for i := 0; i < 3; i++ {
		b := GetFake()
		bs = append(bs, b)
	}
	return bs, nil
}
