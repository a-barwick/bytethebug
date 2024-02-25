package client

import (
	"bytethebug/internal/models"
	"bytethebug/internal/templates/components"
	"bytethebug/internal/templates/shared"

	"github.com/a-h/templ"
)

func BuildTodoPage(data []models.Todo) templ.Component {
	todoCmp := components.TodoApp(data)
	page := shared.BaseTemplate(todoCmp)
	return page
}

func BuildTodoList(data []models.Todo) templ.Component {
	return components.TodoList(data)
}
