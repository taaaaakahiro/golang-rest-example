package template

import (
	errs "github.com/pkg/errors"
	"path/filepath"
)

const (
	htmlDir = "template/html"
)

type Template struct {
	Path string
}

func NewTemplate() *Template {
	path, err := filepath.Abs(htmlDir)
	if err != nil {
		errs.WithStack(err)
	}
	return &Template{
		Path: path,
	}
}
