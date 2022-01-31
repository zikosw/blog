package main

import (
	blog_repo "blog/app/internal/repo/blog"
	"blog/domain"
	"blog/ent"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/labstack/echo/v4"
)

func initDB() error {
	log.Default().Println("initDB")
	// client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	fmt.Println("connect db")
	client, err := ent.Open("sqlite3", "file:blog.sqlite?&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	fmt.Println("migrate schema")
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	ctx := context.Background()

	rpCreate := blog_repo.CreateBlog(ctx, client)

	title := "title"
	content := domain.Content("content: " + time.Now().String())

	s, err := domain.NewSlug(title)
	if err != nil {
		return err
	}

	b, err := domain.NewBlog(title, content, *s)
	if err != nil {
		return err
	}

	// newB, err := rpCreate(*b)
	// if err != nil {
	// 	return err
	// }

	newB, err := domain.CreateBlog(*b, domain.CreateBlogCfg{
		RepoCreate: rpCreate,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n\n\n", newB)

	getB, err := blog_repo.Get(ctx, client)(newB.ID())
	if err != nil {
		return err
	}

	fmt.Printf("get_b:: %+v\n\n", getB)

	// b, err := repo.CreateBlog(ctx, client)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n\n", b)

	// qb, err := repo.QueryBlog(ctx, client)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n\n", qb)

	return nil
}

func main() {
	fmt.Println("blog starting...")
	if err := initDB(); err != nil {
		panic(err)
	}

	// http server
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8000"))
	fmt.Println("blog stopped")
}
