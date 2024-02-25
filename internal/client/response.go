package client

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

func WriteHtmlResponse(component templ.Component, context context.Context, writer http.ResponseWriter) {
	component.Render(context, writer)
}
