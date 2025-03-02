package blog_test

import (
	"errors"
	"io/fs"
	"master/blog"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct{}

func (f StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, I always fail")
}

func TestNewPostsFromFS(t *testing.T) {

	t.Run("Load the file with specific text", func(t *testing.T) {
		files := fstest.MapFS{
			"hello world.md":  {Data: []byte("Title: Post 1")},
			"hello-world2.md": {Data: []byte("Title: Post 2")},
		}

		posts, err := blog.NewPostsFromFS(files)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(files) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(files))
		}
		got := posts[0]
		want := blog.Post{Title: "Post 1"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})
}
