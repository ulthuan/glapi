package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"github.com/ulthuan/glapi/models"
)

// SettingsShow default implementation.
func UserSettings(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	scmProviders := &models.ScmProviders{}
	// Retrieve all Projects from the DB
	if err := tx.All(scmProviders); err != nil {
		return errors.WithStack(err)
	}

	var providers = map[string]bool{
		"github":    false,
		"gitlab":    false,
		"bitbucket": false,
	}
	for _, prov := range *scmProviders {
		providers[prov.Name] = true
	}

	c.Set("providers", providers)
	return c.Render(200, r.HTML("settings.html"))
}
