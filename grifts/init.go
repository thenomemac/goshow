package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/thenomemac/goshow/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
