package main

import (
	"html/template"
	"io"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

//	type Count struct {
//		Count int
//	}
var id = 0

type Contact struct {
	Name  string
	Email string
	Id    int
}

func NewContacts(name, email string) Contact {
	id++
	return Contact{
		Name:  name,
		Email: email,
		Id:    id,
	}
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

func NewData() Data {
	return Data{
		Contacts: Contacts{
			NewContacts("John", "jn@foo.com"),
			NewContacts("Jane", "je@foo.com"),
			NewContacts("James", "js@bar.com"),
		},
	}
}

func (d *Data) hasEmail(email string) bool {
	for _, contact := range d.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}

func (d *Data) indexOf(id int) int {
	for i, contact := range d.Contacts {
		if contact.Id == id {
			return i
		}
	}
	return -1
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func newFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type Page struct {
	Data     Data
	FormData FormData
}

func NewPage() Page {
	return Page{
		Data:     NewData(),
		FormData: newFormData(),
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	// count := Count{Count: 0}
	// data := NewData()
	page := NewPage()
	e.Renderer = newTemplate()

	e.Static("/images", "images")
	e.Static("/css", "css")

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", page)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		if page.Data.hasEmail(email) {
			formData := newFormData()
			formData.Values["name"] = name
			formData.Values["email"] = email
			formData.Errors["email"] = "Email already exists!"
			return c.Render(422, "form", formData)
		}

		contact := NewContacts(name, email)
		page.Data.Contacts = append(page.Data.Contacts, contact)

		err := c.Render(200, "form", newFormData())
		if err != nil {
			return err
		}
		return c.Render(200, "oob-contact", contact)
	})

	e.DELETE("/contacts/:id", func(c echo.Context) error {
		time.Sleep(time.Second * 2)
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.String(400, "Invalid ID")
		}

		index := page.Data.indexOf(id)
		if index == -1 {
			return c.String(404, "Contact not found")
		}
		page.Data.Contacts = append(page.Data.Contacts[:index], page.Data.Contacts[index+1:]...)

		return c.NoContent(200)

	})

	// e.GET("/", func(c echo.Context) error {
	// 	return c.Render(200, "index", count)
	// })
	// e.POST("/count", func(c echo.Context) error {
	// 	count.Count++
	// 	return c.Render(200, "count", count)
	// })

	e.Logger.Fatal(e.Start(":1323"))
}
