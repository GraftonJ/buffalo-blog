package grifts

import (
	"github.com/GraftonJ/blog_app/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
