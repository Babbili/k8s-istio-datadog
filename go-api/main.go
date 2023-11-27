package main

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	echotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo.v4"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

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

// Kubernetes healthProbes
func healthProbes(c echo.Context) error {
	started := time.Now()
	duration := time.Now().Sub(started)
	if duration.Seconds() > 10 {
		return echo.NewHTTPError(http.StatusInternalServerError)
	} else {
		return c.String(http.StatusOK, "OK")
	}
}

func main() {
	// start dd-trace
	tracer.Start(
		tracer.WithEnv("prod"),
		tracer.WithService("goapi"),
		tracer.WithServiceVersion("v1"),
	)
	// When the tracer is stopped, it will flush everything it has to the Datadog Agent before quitting.
	defer tracer.Stop()

	e := echo.New()

	// Middleware
	e.Use(echotrace.Middleware(echotrace.WithServiceName("goapi"), echotrace.WithCustomTag("env", "prod")))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Go books API!\n")
	})
	e.GET("/books", getAllBooks)
	e.POST("/books", createBook)
	e.GET("/books/:id", getBook)
	e.DELETE("/books/:id", deleteBook)

	e.GET("/healthz", healthProbes)

	// Start server
	e.Logger.Fatal(e.Start(":50051"))
}
