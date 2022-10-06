package main

type BlogPost struct {
	Location string `json:"location"`
	Title    string `json:"title"`
	Slug     string `json:"slug"`
	Body     string `json:"body"`
}

func GetBlogPostBySlug(slug string) (BlogPost, *ApplicationError) {
	if slug == "bad" {
		err := ValidationError("bad slug")
		return BlogPost{}, err
	}
	return BlogPost{
		Location: "dummy",
		Title:    slug,
		Slug:     slug,
		Body:     "lorem ipsum dolor baby",
	}, nil
}
