package blogrender

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
	ParseMarkdown            func(string) template.HTML
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

var (
	//go:embed "templates/*"
	postTemplate embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {

	templ, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (rr *PostRenderer) Render(buf io.Writer, post Post) error {

	post.ParseMarkdown = func(s string) template.HTML {
		result := string(markdown.ToHTML([]byte(s), nil, nil))
		return template.HTML(result)
	}

	return rr.templ.ExecuteTemplate(buf, "blog.gohtml", post)
}

func (rr *PostRenderer) RenderViewModel(buf io.Writer, post Post) error {

	vm := &postViewModel{
		Post:     post,
		HTMLBody: template.HTML(string(markdown.ToHTML([]byte(post.Body), nil, nil))),
	}

	return rr.templ.ExecuteTemplate(buf, "blog.vm.gohtml", vm)
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

func (rr *PostRenderer) RenderIndex(buf io.Writer, posts []Post) error {
	return rr.templ.ExecuteTemplate(buf, "index.gohtml", posts)
}
