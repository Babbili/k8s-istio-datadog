package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	book struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Author string `json:"author"`
	}
)

var (
	books = map[int]*book{}
	seq   = 1
	lock  = sync.Mutex{}
)

// Handlers

func createBook(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := &book{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	books[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getBook(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, books[id])
}

func deleteBook(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	delete(books, id)
	return c.NoContent(http.StatusNoContent)
}

func getAllBooks(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	return c.JSON(http.StatusOK, books)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/books", getAllBooks)
	e.POST("/books", createBook)
	e.GET("/books/:id", getBook)
	e.DELETE("/books/:id", deleteBook)

	// Start server
	e.Logger.Fatal(e.Start(":50051"))
}
