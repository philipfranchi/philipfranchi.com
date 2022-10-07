package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

type MockBlogManager struct {
	mockGetBlogPostBySlug      func(string) (BlogPost, *ApplicationError)
	mockGetAllBlogPostMetadata func() ([]BlogPostMetadata, *ApplicationError)
}

func (m *MockBlogManager) GetBlogPostBySlug(slug string) (BlogPost, *ApplicationError) {
	return m.mockGetBlogPostBySlug(slug)
}

func (m *MockBlogManager) GetAllBlogPostMetadata() ([]BlogPostMetadata, *ApplicationError) {
	return m.mockGetAllBlogPostMetadata()
}

func createMockGetBlogPostBySlug(
	post BlogPost,
	err *ApplicationError,
) func(s string) (BlogPost, *ApplicationError) {
	return func(s string) (BlogPost, *ApplicationError) {
		return post, err
	}
}

func createMockGetAllBlogPostMetadata(
	data []BlogPostMetadata,
	err *ApplicationError,
) func() ([]BlogPostMetadata, *ApplicationError) {
	return func() ([]BlogPostMetadata, *ApplicationError) {
		return data, err
	}
}

func createFailMockGetBlogPostBySlug(t *testing.T) func(s string) (BlogPost, *ApplicationError) {
	return func(s string) (BlogPost, *ApplicationError) {
		panic("should not call")
	}
}

func createFailMockGetAllBlogPostMetadata(
	t *testing.T,
) func() ([]BlogPostMetadata, *ApplicationError) {
	return func() ([]BlogPostMetadata, *ApplicationError) {
		panic("should not call")
	}
}

func TestSinglePostHandler(t *testing.T) {
	dummyPost := BlogPost{Body: "body", BlogPostMetadata: BlogPostMetadata{Title: "slug"}}
	createBody := func(post BlogPost) []byte {
		b, _ := json.Marshal(post)
		return b
	}

	var tests = []struct {
		name   string
		slug   string
		body   []byte
		status int
		mock   func(s string) (BlogPost, *ApplicationError)
	}{
		{
			"respects error",
			"doesnt_matter",
			[]byte("bad slug"),
			http.StatusBadRequest,
			createMockGetBlogPostBySlug(BlogPost{}, ValidationError()),
		},
		{
			"validates slug",
			"",
			[]byte("bad slug"),
			http.StatusBadRequest,
			createFailMockGetBlogPostBySlug(t),
		},
		{
			"success",
			"slug",
			createBody(dummyPost),
			http.StatusOK,
			createMockGetBlogPostBySlug(dummyPost, nil),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := CreateAPIHandler(
				&MockBlogManager{
					mockGetBlogPostBySlug:      test.mock,
					mockGetAllBlogPostMetadata: createFailMockGetAllBlogPostMetadata(t),
				},
			)
			request := httptest.NewRequest(
				http.MethodGet,
				"/api/blog/"+test.slug,
				strings.NewReader(""),
			)
			request = mux.SetURLVars(request, map[string]string{
				"slug": test.slug,
			})

			response := httptest.NewRecorder()
			handler.HandleGetSingleBlogPost(response, request)
			if test.status != response.Result().StatusCode {
				t.Errorf("want %d got %d", test.status, response.Code)
			}
			if string(test.body) != response.Body.String() {
				t.Errorf("want %#v got %#v", string(test.body), response.Body.String())
			}
		})
	}
}

func TestGetBlogPostTitles(t *testing.T) {
	type singleTest struct {
		name          string
		status        int
		expectedBody  []byte
		expectedError *ApplicationError
		mock          func() ([]BlogPostMetadata, *ApplicationError)
	}

	var tests = []singleTest{
		func() singleTest {
			expectedBody, _ := json.Marshal([]BlogPostMetadata{})
			return singleTest{"success", 200, expectedBody, nil, createMockGetAllBlogPostMetadata(
				[]BlogPostMetadata{},
				nil,
			)}
		}(),
		func() singleTest {
			return singleTest{"respects_error", 500, []byte("some error"), nil, createMockGetAllBlogPostMetadata(
				[]BlogPostMetadata{},
				MarshallingError("some error"),
			)}
		}(),
	}
	for _, test := range tests {
		handler := CreateAPIHandler(
			&MockBlogManager{
				mockGetBlogPostBySlug:      createFailMockGetBlogPostBySlug(t),
				mockGetAllBlogPostMetadata: test.mock,
			},
		)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/api/blog/", strings.NewReader(""))
		handler.HandleGetAllBlogPostMetadata(response, request)

		if test.status != response.Result().StatusCode {
			t.Errorf("want %d got %d", test.status, response.Code)
		}

		if string(test.expectedBody) != response.Body.String() {
			t.Errorf("want %#v got %#v", string(test.expectedBody), response.Body.String())
		}
	}
}
