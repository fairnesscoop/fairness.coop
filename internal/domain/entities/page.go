package entities

import (
	"fmt"
	"os"
	"path/filepath"
)

type File struct {
	collection *Collection
	path       string
}

func NewFile(collection *Collection, path string) *File {
	return &File{
		collection: collection,
		path:       path,
	}
}

func (f *File) Path() string {
	return f.path
}

func (f *File) Collection() *Collection {
	return f.collection
}

func (f *File) Url() string {
	return fmt.Sprintf("/collections/%d/posts/%s", f.collection.Id, f.path)
}

func (f *File) Read() (string, error) {
	path := filepath.Join(f.collection.Path, f.path)

	content, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(content), nil
}

func (f *File) Write(content string) error {
	path := filepath.Join(f.collection.Path, f.path)
	return os.WriteFile(path, []byte(content), 0666)
}
