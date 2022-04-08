package blogrender_test

import (
	"blogrender"
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestBlogRender(t *testing.T) {

	approvals.UseFolder("testdata")

	buf := bytes.Buffer{}
	aPost := blogrender.Post{
		Title:       "hello world",
		Body:        "This is **a post**",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	rr, err := blogrender.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	if err := rr.Render(&buf, aPost); err != nil {
		t.Fatal(err)
	}

	approvals.VerifyString(t, buf.String())
}

func TestRenderViewModel(t *testing.T) {
	approvals.UseFolder("testdata")

	buf := bytes.Buffer{}
	aPost := blogrender.Post{
		Title:       "hello world",
		Body:        "This is **a post**",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	rr, err := blogrender.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	if err := rr.RenderViewModel(&buf, aPost); err != nil {
		t.Fatal(err)
	}

	approvals.VerifyString(t, buf.String())
}

func TestIndexRender(t *testing.T) {

	approvals.UseFolder("testdata")

	buf := bytes.Buffer{}
	posts := []blogrender.Post{
		{Title: "First post"},
		{Title: "Second post"},
	}

	rr, err := blogrender.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	if err := rr.RenderIndex(&buf, posts); err != nil {
		t.Error(err)
	}

	approvals.VerifyString(t, buf.String())
}

func BenchmarkBlogRender(b *testing.B) {
	aPost := blogrender.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	rr, err := blogrender.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rr.Render(io.Discard, aPost)
	}
}
