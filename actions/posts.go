package actions

import "github.com/gobuffalo/buffalo"

// PostsIndex default implementation.
func PostsIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("posts/index.html"))
}

// PostsCreate default implementation.
func PostsCreate(c buffalo.Context) error {
	return c.Render(200, r.HTML("posts/create.html"))
}

// PostsEdit default implementation.
func PostsEdit(c buffalo.Context) error {
	return c.Render(200, r.HTML("posts/edit.html"))
}

// PostsDelete default implementation.
func PostsDelete(c buffalo.Context) error {
	return c.Render(200, r.HTML("posts/delete.html"))
}

// PostsDetail default implementation.
func PostsDetail(c buffalo.Context) error {
	return c.Render(200, r.HTML("posts/detail.html"))
}
