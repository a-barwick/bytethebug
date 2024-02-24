package main

import (
	"bytethebug/internal/models"
	"bytethebug/internal/provider"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
)

const (
	testdata = "testdata"
)

func seedTestData() []models.Todo {
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
		provider.WriteHtmlResponse(provider.BuildTodoList(data), r.Context(), w)
	}
}

func parseArgs() string {
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "--testdata" || arg == "-t" {
			return testdata
		}
	}
	return ""
}

func main() {
	homePage := provider.BuildHomePage()

	http.Handle("/", templ.Handler(homePage))

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
