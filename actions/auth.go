package actions

import (
	"fmt"
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/markbates/going/defaults"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/pkg/errors"
	"github.com/ulthuan/glapi/models"
	"github.com/ulthuan/goth/providers/glo"
)

func init() {
	gothic.Store = App().SessionStore

	goth.UseProviders(
		glo.New(
			os.Getenv("GLO_KEY"),
			os.Getenv("GLO_SECRET"),
			fmt.Sprintf("%s%s", App().Host, "/auth/glo/callback"),
			"board:write,user:read",
		),
		github.New(
			os.Getenv("GITHUB_KEY"),
			os.Getenv("GITHUB_SECRET"),
			fmt.Sprintf("%s%s", App().Host, "/auth/scm/github/callback"),
			"user:email",
		),
	)
}

func ScmAuthCallback(c buffalo.Context) error {
	currentUser := c.Value("current_user").(*models.User)
	gu, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(401, err)
	}
	tx := c.Value("tx").(*pop.Connection)
	if gu.Email != currentUser.Email.String {
		return c.Error(401, errors.New("Unexpected error"))
	}
	findProvider := false
	provider := &models.ScmProvider{}
	for _, *provider = range currentUser.ScmProviders {
		if provider.Name == gu.Provider {
			provider.ScmProviderToken = gu.AccessToken
			findProvider = true
		}
	}

	if findProvider == false {
		provider = &models.ScmProvider{}
		provider.Name = gu.Provider
		provider.ScmProviderToken = gu.AccessToken
		provider.UserID = currentUser.ID
		provider.ScmProviderID = gu.UserID
	}

	if err = tx.Save(provider); err != nil {
		return errors.WithStack(err)
	}

	c.Flash().Add("success", "You Connect your account")
	return c.Redirect(302, "/settings")
}

func UserAuthCallback(c buffalo.Context) error {
	gu, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return c.Error(401, err)
	}
	tx := c.Value("tx").(*pop.Connection)
	q := tx.Where("provider = ? and provider_id = ?", gu.Provider, gu.UserID)
	exists, err := q.Exists("users")
	if err != nil {
		return errors.WithStack(err)
	}
	u := &models.User{}
	if exists {
		if err = q.First(u); err != nil {
			return errors.WithStack(err)
		}
	}
	u.Name = defaults.String(gu.Name, gu.NickName)
	u.Provider = gu.Provider
	u.ProviderID = gu.UserID
	u.ProviderToken = gu.AccessToken
	u.Email = nulls.NewString(gu.Email)
	u.Picture = gu.AvatarURL
	if err = tx.Save(u); err != nil {
		return errors.WithStack(err)
	}

	c.Session().Set("current_user_id", u.ID)
	if err = c.Session().Save(); err != nil {
		return errors.WithStack(err)
	}

	c.Flash().Add("success", "You have been logged in")
	rTo := c.Session().GetOnce("login_redirect_to")
	rToStr, ok := rTo.(string)
	if !ok || rToStr == "" {
		rToStr = "/"
	}

	return c.Redirect(302, rToStr)
}

func AuthDestroy(c buffalo.Context) error {
	c.Session().Clear()
	c.Flash().Add("success", "You have been logged out")
	return c.Redirect(302, "/")
}

func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			if err := tx.Find(u, uid); err != nil {
				return errors.WithStack(err)
			}
			c.Set("current_user", u)
		}
		return next(c)
	}
}

func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid == nil {
			c.Flash().Add("danger", "You must be authorized to see that page")
			c.Session().Set("login_redirect_to", c.Request().URL.String())
			return c.Redirect(302, "/login")
		}
		return next(c)
	}
}

func Logout(c buffalo.Context) error {
	session := c.Session()
	session.Delete("current_user")
	session.Delete("current_user_id")
	session.Save()
	return c.Redirect(301, "/login")
}
