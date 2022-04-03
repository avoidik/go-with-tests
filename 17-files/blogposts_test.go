package blogposts_test

import (
	"blogposts"
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

const (
	firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, golang
---
Hello
World
`
	secondBody = `Title: Post 2
Description: Description 2
Tags: tdd
---
Hello
Cats
`
)

type StubFailingFs struct{}

func (s StubFailingFs) Open(name string) (fs.File, error) {
	return nil, errors.New("dirfs failure")
}

func assertPost(t *testing.T, postGot blogposts.Post, postWant blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(postGot, postWant) {
		t.Errorf("got %+v, want %+v", postGot, postWant)
	}
}

func TestNewBlogposts(t *testing.T) {

	fs := fstest.MapFS{
		"hello.md":  {Data: []byte(firstBody)},
		"hello2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "golang"},
		Body: `Hello
World`},
	)

	assertPost(t, posts[1], blogposts.Post{
		Title:       "Post 2",
		Description: "Description 2",
		Tags:        []string{"tdd"},
		Body: `Hello
Cats`},
	)

	_, err = blogposts.NewPostsFromFS(StubFailingFs{})
	if err == nil {
		t.Error("expected error")
	}
}
