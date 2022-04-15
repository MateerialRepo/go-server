package main

import (
	"fmt"
	"github.com/MateerialRepo/go-webapp/pkg/config"
	"github.com/MateerialRepo/go-webapp/pkg/handlers"
	"github.com/MateerialRepo/go-webapp/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

//PORT is a constant because we do not expect it to change
const PORT = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	//change this to true in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	//tc is template cache....a cache of all the templates in the application
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache. Error: ", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println("This application is running on port ", PORT)

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
