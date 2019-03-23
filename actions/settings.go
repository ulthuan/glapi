package actions

import "github.com/gobuffalo/buffalo"

// SettingsShow default implementation.
func UserSettings(c buffalo.Context) error {
	return c.Render(200, r.HTML("settings.html"))
}
