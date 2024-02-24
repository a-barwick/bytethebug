package main

import (
	"e-commerce/models"
	"e-commerce/templates"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func createTestData() []models.Todo {
	return []models.Todo{
		{
			Id:        "1",
			Body:      "Do something",
			IsChecked: false,
			CreatedAt: "2021-01-01T00:00:00Z",
			UpdatedAt: "2021-01-01T00:00:00Z",
		},
		{
			Id:        "2",
			Body:      "Do something else",
			IsChecked: true,
			CreatedAt: "2021-01-01T00:00:00Z",
			UpdatedAt: "2021-01-01T00:00:00Z",
		},
	}
}

func createItemHandler(data []models.Todo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		formValues := r.Form
		body := formValues.Get("todo")
		if body != "" {
			data = append(data, models.Todo{
				Id:        "3",
				Body:      body,
				IsChecked: false,
				CreatedAt: "2021-01-01T00:00:00Z",
				UpdatedAt: "2021-01-01T00:00:00Z",
			})
		}
		templates.ItemList(data).Render(r.Context(), w)
	}
}

func main() {
	data := createTestData()

	todoPage := templates.TodoComponent(data)
	base := templates.BasePage(todoPage)

	http.Handle("/", templ.Handler(base))
	http.HandleFunc("/create", createItemHandler(data))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
