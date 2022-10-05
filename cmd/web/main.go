package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sanathshetty444/go-bookings/pkg/config"
	"github.com/sanathshetty444/go-bookings/pkg/handlers"
	"github.com/sanathshetty444/go-bookings/pkg/render"
)

var port string = ":3333"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	ts, e := render.CreateTemplateCache()
	if e != nil {
		log.Fatal(e)
	}

	app.TemplateCache = ts
	app.UseCache = true
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = app.InProduction
	session.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = session

	render.NewTemplateCache(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Println("Starting application on port", port)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
