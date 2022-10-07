package main

import (
	"log"
	"os"
	"path"
)

type BlogPostMetadata struct {
	Title string `json:"title"`
}

type BlogPost struct {
	BlogPostMetadata
	Body string `json:"body"`
}

type BlogManager interface {
	GetBlogPostBySlug(string) (BlogPost, *ApplicationError)
	GetAllBlogPostMetadata() ([]BlogPostMetadata, *ApplicationError)
}

type FileBasedBlogManager struct {
	contentRoot string
}

func CreateBlog(config ProjectConfig) BlogManager {
	b := FileBasedBlogManager{contentRoot: config.ContentRoot}
	return &b
}

func (b *FileBasedBlogManager) GetBlogPostBySlug(slug string) (BlogPost, *ApplicationError) {
	postPath := path.Join(b.contentRoot, slug+".md")
	log.Println("Looking for post at: " + postPath)
	dat, err := os.ReadFile(postPath)
	if err != nil {
		return BlogPost{}, PostMissingError(slug)
	}
	return BlogPost{BlogPostMetadata: BlogPostMetadata{Title: slug}, Body: string(dat)}, nil
}

func (b *FileBasedBlogManager) GetAllBlogPostMetadata() ([]BlogPostMetadata, *ApplicationError) {
	panic("not implemented")
}
