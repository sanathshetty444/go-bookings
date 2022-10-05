package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/sanathshetty444/go-bookings/pkg/config"
	"github.com/sanathshetty444/go-bookings/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplateCache(a *config.AppConfig) {
	app = a
}

//Capital means public
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsed_template, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsed_template.Execute(w, nil)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Error: %+v", err))
		return
	}
}

func RenderTemplateBase(w http.ResponseWriter, tmpl string, td models.TemplateData) {

	var ts map[string]*template.Template

	if app.UseCache {
		ts = app.TemplateCache
	} else {
		ts, _ = CreateTemplateCache()
	}

	t, ok := ts[tmpl]

	if !ok {
		log.Fatal()
	}

	buf := new(bytes.Buffer)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Errorf("Error in RenderTemplateBase", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			fmt.Println("inside")
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
		fmt.Println(myCache)
	}
	return myCache, nil
}
