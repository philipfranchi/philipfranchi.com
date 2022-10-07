package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

type BlogPostMetadata struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
	Date  string `json:"date"`
	Tags  string `json:"tags"`
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
	postFile, err := os.Open(postPath)
	if err != nil {
		log.Println("Unable to find post at: " + postPath)
		return BlogPost{}, PostMissingError(slug)
	}
	return ParseBlog(postFile)
}

func (b *FileBasedBlogManager) GetAllBlogPostMetadata() ([]BlogPostMetadata, *ApplicationError) {
	posts, _ := os.ReadDir(b.contentRoot)
	var data []BlogPostMetadata = []BlogPostMetadata{}
	for _, postPath := range posts {
		postFile, err := os.Open(path.Join(b.contentRoot, postPath.Name()))
		if err != nil {
			return []BlogPostMetadata{}, InternalServerError(err.Error())
		}
		defer postFile.Close()
		blog, err := ParseBlog(postFile)
		data = append(data, blog.BlogPostMetadata)
	}
	return data, nil
}

func ParseMetadataField(row string) (string, string, *ApplicationError) {
	// Metadata fields take the form [comment]: # (Key: Value)
	fieldRegex := regexp.MustCompile(`\[comment\]: # \((\w*): (.*)\)`)
	results := fieldRegex.FindAllStringSubmatch(row, -1)
	if len(results) != 1 {
		return "", "", BlogParseError("Cannot read field: " + row)
	}
	if len(results[0]) != 3 {
		return "", "", BlogParseError("Cannot read field: " + row)
	}
	return results[0][1], results[0][2], nil
}

func ParseBlog(rs io.ReadSeeker) (BlogPost, *ApplicationError) {
	var lines []string
	scanner := bufio.NewScanner(rs)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) < 4 {
		return BlogPost{}, BlogParseError("Cannot read metadata from blog")
	}

	_, title, err := ParseMetadataField(lines[0])
	if err != nil {
		return BlogPost{}, err
	}
	_, slug, err := ParseMetadataField(lines[1])
	if err != nil {
		return BlogPost{}, err
	}
	_, date, err := ParseMetadataField(lines[2])
	if err != nil {
		return BlogPost{}, err
	}
	_, tags, err := ParseMetadataField(lines[3])
	if err != nil {
		return BlogPost{}, err
	}

	metadata := BlogPostMetadata{Title: title, Slug: slug, Date: date, Tags: tags}
	body := strings.Join(lines[4:], "\n")
	return BlogPost{BlogPostMetadata: metadata, Body: body}, nil
}
