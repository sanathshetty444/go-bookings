package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/sanathshetty444/go-bookings/pkg/config"
	"github.com/sanathshetty444/go-bookings/pkg/models"
	"github.com/sanathshetty444/go-bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	app *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{app: a}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	repo.app.Session.Put(r.Context(), "remoteIP", r.RemoteAddr)

	n, _ := fmt.Fprintf(w, fmt.Sprintf("Hello world %d", 5))
	fmt.Println("No of bytes:", n)
}
func (repo *Repository) Divide(w http.ResponseWriter, r *http.Request) {
	n, err := divideValues(100, 0)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Error: %+v", err))
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("Value is: %f", n))
}
func divideValues(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return x / y, nil
}

func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	// render.RenderTemplate(w, "about.page.tmpl")
	StringMap := make(map[string]string)
	StringMap["test"] = "Hello world"
	StringMap["remoteIP"] = repo.app.Session.GetString(r.Context(), "remoteIP")

	fmt.Printf("%+v", repo.app.Session.Cookie)
	td := models.TemplateData{StringMap: StringMap}
	render.RenderTemplateBase(w, "about.page.tmpl", td)
}
