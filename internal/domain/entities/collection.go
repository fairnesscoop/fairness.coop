package entities

import "fmt"

type Collection struct {
	Id    int
	Label string
	Name  string
	Path  string
	files []*File
}

func NewCollection(id int, label string, name string, path string) *Collection {
	return &Collection{
		Id:    id,
		Label: label,
		Name:  name,
		Path:  path,
		files: make([]*File, 0),
	}
}

func (c *Collection) Url() string {
	return fmt.Sprintf("/collections/%d", c.Id)
}

func (c *Collection) AddFile(f *File) {
	c.files = append(c.files, f)
}

func (c *Collection) FindFileByPath(path string) *File {
	for _, file := range c.files {
		if file.path == path {
			return file
		}
	}

	return nil
}

func (c *Collection) Files() []*File {
	return c.files
}
