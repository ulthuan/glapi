package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/ulthuan/glapi/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
