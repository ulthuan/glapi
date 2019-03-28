package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func LoginHandler(c buffalo.Context) error {
	c.Set("provider", "glo")
	return c.Render(200, r.HTML("login.html", "page_layout.html"))
}
