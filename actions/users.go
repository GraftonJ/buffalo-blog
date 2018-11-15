package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/GraftonJ/blog_app/models"
)

// UserRegisterGet displays a register form
func UsersRegisterGet(c buffalo.Context) error {
	// Make user available inside the html template
	c.Set("user", &models.User{})
	return c.Render(200, r.HTML("users/register.html"))
}
