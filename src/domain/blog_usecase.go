package domain

type RepoCreate func(Blog) (*Blog, error)

type CreateBlogCfg struct {
	RepoCreate RepoCreate
}

func CreateBlog(b Blog, cfg CreateBlogCfg) (*Blog, error) {

	newB, err := cfg.RepoCreate(b)
	if err != nil {
		return nil, err
	}

	return newB, nil
}
