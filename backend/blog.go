package main

import (
	"log"
	"os"
	"path"
)

type BlogPost struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Blog struct {
	contentRoot string
}

func CreateBlog(config ProjectConfig) *Blog {
	b := Blog{contentRoot: config.ContentRoot}
	return &b
}
func (b *Blog) GetBlogPostBySlug(slug string) (BlogPost, *ApplicationError) {
	if slug == "bad" {
		err := ValidationError("bad slug")
		return BlogPost{}, err
	}
	postPath := path.Join(b.contentRoot, slug+".md")
	log.Println("Looking for post at: " + postPath)
	dat, err := os.ReadFile(postPath)
	if err != nil {
		return BlogPost{}, PostMissingError("Post not found")
	}
	return BlogPost{Title: slug, Body: string(dat)}, nil
}
