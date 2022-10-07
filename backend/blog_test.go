package main

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"testing"
)

func populateTestDir(dir string, posts []BlogPost) {
	metadataFieldFormat := "[comment]: # (%s: %s)\n"
	for _, post := range posts {
		title := fmt.Sprintf(metadataFieldFormat, "TITLE", post.Title)
		slug := fmt.Sprintf(metadataFieldFormat, "SLUG", post.Slug)
		date := fmt.Sprintf(metadataFieldFormat, "DATE", post.Date)
		tags := fmt.Sprintf(metadataFieldFormat, "TAGS", post.Tags)
		data := []byte(title + slug + date + tags + post.Body)
		// Should this be 777? Probably not
		os.WriteFile(path.Join(dir, post.Slug+".md"), data, 0777)
	}
}
func TestGetBlogPostBySlug(t *testing.T) {
	tests := []struct {
		name     string
		posts    []BlogPost
		slug     string
		want     BlogPost
		want_err *ApplicationError
	}{
		{
			"no post",
			[]BlogPost{},
			"some_slug",
			BlogPost{},
			PostMissingError("some_slug"),
		},
		{
			"no post but data",
			[]BlogPost{{BlogPostMetadata{Title: "post", Tags: "post", Date: "date", Slug: "post"}, "body"}},
			"some_slug",
			BlogPost{},
			PostMissingError("some_slug"),
		},
		{
			"one post",
			[]BlogPost{{BlogPostMetadata{Title: "post", Tags: "post", Date: "date", Slug: "post"}, "body"}},
			"post",
			BlogPost{BlogPostMetadata{Title: "post", Tags: "post", Date: "date", Slug: "post"}, "body"},
			nil,
		},
		{
			"two post",
			[]BlogPost{{BlogPostMetadata{Title: "post", Tags: "post", Date: "date", Slug: "post"}, "body"}, {BlogPostMetadata{Title: "post2", Tags: "post2", Date: "date", Slug: "post2"}, "body"}},
			"post2",
			BlogPost{BlogPostMetadata{Title: "post2", Tags: "post2", Date: "date", Slug: "post2"}, "body"},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			contentRoot := t.TempDir()
			populateTestDir(contentRoot, tt.posts)
			b := CreateBlog(ProjectConfig{ContentRoot: contentRoot})
			got, err := b.GetBlogPostBySlug(tt.slug)

			if !reflect.DeepEqual(tt.want_err, err) {
				t.Errorf("want_err %v got_err %v", tt.want_err, err)
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want %v got %v", tt.want, got)
			}
		})
	}
}

func TestGetAllBlogMetadata(t *testing.T) {
	examplePost := BlogPost{BlogPostMetadata{Slug: "post", Title: "title", Tags: "hello", Date: "date"}, "# body\n\n\nhello"}
	var tests = []struct {
		name             string
		existingPosts    []BlogPost
		expectedMetadata []BlogPostMetadata
		expectedError    *ApplicationError
	}{
		{"success", []BlogPost{examplePost}, []BlogPostMetadata{examplePost.BlogPostMetadata}, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			contentRoot := t.TempDir()
			populateTestDir(contentRoot, test.existingPosts)
			blog := CreateBlog(ProjectConfig{ContentRoot: contentRoot})
			posts, err := blog.GetAllBlogPostMetadata()
			if !reflect.DeepEqual(test.expectedError, err) {
				t.Errorf("want %#v got %#v", test.expectedError, err)
			}
			if !reflect.DeepEqual(test.expectedMetadata, posts) {
				t.Errorf("want %#v got %#v", test.expectedMetadata, posts)
			}
		})
	}
}
