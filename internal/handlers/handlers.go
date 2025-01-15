package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos []Todo

func HandleHome(c echo.Context) error {
	return c.Render(http.StatusOK, "base.html", nil)
}

func HandleGetTodos(c echo.Context) error {
	return c.Render(http.StatusOK, "todo_list.html", map[string]interface{}{
		"Todos": todos,
	})
}

func HandleCreateTodo(c echo.Context) error {
	title := c.FormValue("title")
	todo := Todo{
		ID:        len(todos) + 1,
		Title:     title,
		Completed: false,
	}
	todos = append(todos, todo)
	return c.Render(http.StatusOK, "todo_list.html", map[string]interface{}{
		"Todos": todos,
	})
}

func HandleToggleTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = !todos[i].Completed
			break
		}
	}
	return c.Render(http.StatusOK, "todo_list.html", map[string]interface{}{
		"Todos": todos,
	})

}

func HandleDeleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i := range todos {
		if todos[i].ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	return c.Render(http.StatusOK, "todo_list.html", map[string]interface{}{
		"Todos": todos,
	})
}
