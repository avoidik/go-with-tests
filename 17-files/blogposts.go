package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

const (
	titleSeparator       = "Title: "
	descriptionSeperator = "Description: "
	tagsSeparator        = "Tags: "
)

func readBody(s *bufio.Scanner) string {
	s.Scan()

	buf := bytes.Buffer{}
	for s.Scan() {
		fmt.Fprintln(&buf, s.Text())
	}
	body := strings.TrimSuffix(buf.String(), "\n")

	return body
}

func newPost(fsFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(fsFile)

	readLine := func(s *bufio.Scanner, tagName string) string {
		s.Scan()
		return strings.TrimPrefix(s.Text(), tagName)
	}

	return Post{
		Title:       readLine(scanner, titleSeparator),
		Description: readLine(scanner, descriptionSeperator),
		Tags:        strings.Split(readLine(scanner, tagsSeparator), ", "),
		Body:        readBody(scanner),
	}, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	fsFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, nil
	}
	defer fsFile.Close()

	return newPost(fsFile)
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	entries, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range entries {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
