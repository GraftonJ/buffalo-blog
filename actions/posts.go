package actions

import (
				"github.com/gobuffalo/buffalo"
				"github.com/gobuffalo/pop"
				"github.com/pkg/errors"
				"github.com/GraftonJ/blog_app/models"
		)
// PostsIndex default implementation.
func PostsIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("posts/index.html"))
}

//Inserted
func PostsCreateGet(c buffalo.Context) error {
	c.Set("post", &models.Post{})
	return c.Render(200, r.HTML("posts/create"))
}

func PostsCreatePost(c buffalo.Context) error {
	// Allocate an empty Post
	post := &models.Post{}
	user := c.Value("current_user").(*models.User)
	// Bind post to the html form elements
	if err := c.Bind(post); err != nil {
		return errors.WithStack(err)
	}
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Validate the data from the html form
	post.AuthorID = user.ID
	verrs, err := tx.ValidateAndCreate(post)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		c.Set("post", post)
		c.Set("errors", verrs.Errors)
		return c.Render(422, r.HTML("posts/create"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "New post added successfully.")
	// and redirect to the index page
	return c.Redirect(302, "/")
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
