package main

import (
	"log"
	"os"
	"path"
)

type BlogPost string

type BlogProvider struct {
	contentRoot string
}

func CreateBlogProvider(config ProjectConfig) *BlogProvider {
	b := BlogProvider{contentRoot: config.ContentRoot}
	return &b
}
func (b *BlogProvider) GetBlogPostBySlug(slug string) (BlogPost, *ApplicationError) {
	if slug == "bad" {
		err := ValidationError("bad slug")
		return "", err
	}
	postPath := path.Join(b.contentRoot, slug+".md")
	log.Println("Looking for post at: " + postPath)
	dat, err := os.ReadFile(postPath)
	if err != nil {
		return "", PostMissingError("Post not found")
	}
	return BlogPost(dat), nil
}
