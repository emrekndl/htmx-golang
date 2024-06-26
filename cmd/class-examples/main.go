package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"net/http"
	"strconv"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Block struct {
	Id int
}

type Blocks struct {
	Start  int
	Next   int
	More   bool
	Blocks []Block
}

func main() {
	// Echo instance
	e := echo.New()

	e.Renderer = NewTemplates()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/blocks", blocksUri)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func blocksUri(ctx echo.Context) error {
	startStr := ctx.QueryParam("start")
	start, err := strconv.Atoi(startStr)
	if err != nil {
		start = 0
	}

	blocks := []Block{}
	for i := start; i < start+10; i++ {
		blocks = append(blocks, Block{Id: i})
	}

	template := "blocks"
	if start == 0 {
		template = "blocks-index"
	}

	return ctx.Render(http.StatusOK, template, Blocks{
		Start:  start,
		Next:   start + 10,
		More:   start+10 < 100,
		Blocks: blocks,
	})

}
