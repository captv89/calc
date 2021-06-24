package actions

import (
	"net/http"
	"strconv"

	"github.com/gobuffalo/buffalo"
	"github.com/gorilla/mux"
)

type Cat struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var cats = []Cat{
	{ID: 1, Name: "Elsa", Age: 16},
	{ID: 2, Name: "Tijger", Age: 12},
	{ID: 3, Name: "Pussy", Age: 19},
}

// CatsList default implementation.
func CatsList(c buffalo.Context) error {
	return c.Render(200, r.JSON(cats))
}

// CatsGetByID default implementation.
func CatsGetByID(c buffalo.Context) error {
	params := mux.Vars(c.Request())
	id, _ := strconv.Atoi(params["id"])
	for _, cat := range cats {
		if cat.ID == id {
			return c.Render(http.StatusOK, r.JSON(cat))
		}
	}
	return c.Render(http.StatusNotFound, nil)
}
