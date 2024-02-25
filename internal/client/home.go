package client

import (
	"bytethebug/internal/templates/components"
	"bytethebug/internal/templates/shared"

	"github.com/a-h/templ"
)

func BuildHomePage() templ.Component {
	homeCmp := components.Home()
	page := shared.BaseTemplate(homeCmp)
	return page
}
