package main

import (
	"os"
	"path"
	"testing"
)

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
			[]BlogPost{{BlogPostMetadata{"post"}, "body"}},
			"some_slug",
			BlogPost{},
			PostMissingError("some_slug"),
		},
		{
			"one post",
			[]BlogPost{{BlogPostMetadata{"post"}, "body"}},
			"post",
			BlogPost{BlogPostMetadata{"post"}, "body"},
			nil,
		},
		{
			"two post",
			[]BlogPost{{BlogPostMetadata{"post"}, "body"}, {BlogPostMetadata{"post2"}, "body"}},
			"post2",
			BlogPost{BlogPostMetadata{"post2"}, "body"},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			contentRoot := t.TempDir()
			for _, post := range tt.posts {
				// Should this be 777? Probably not
				os.WriteFile(path.Join(contentRoot, post.Title+".md"), []byte(post.Body), 0777)
			}

			b := CreateBlog(ProjectConfig{ContentRoot: contentRoot})
			got, err := b.GetBlogPostBySlug(tt.slug)

			if err != nil && tt.want_err != nil && *err != *tt.want_err {
				t.Errorf("want_err %#v got_err %#v", tt.want_err, err)
			} else if err == nil && tt.want_err != nil {
				t.Errorf("want_err %#v got_err %#v", tt.want_err, err)
			} else if err != nil && tt.want_err == nil {
				t.Errorf("want_err %#v got_err %#v", tt.want_err, err)
			} else if err == nil && tt.want_err == nil && got != tt.want {
				t.Errorf("want %#v got %#v", tt.want, got)
			}
		})
	}
}
