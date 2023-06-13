package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/beejan2986/Bookings/pkg/config"
	"github.com/beejan2986/Bookings/pkg/handlers"
	"github.com/beejan2986/Bookings/pkg/render"
	"github.com/joho/godotenv"
)

var portNumber string
var app config.AppConfig
var session *scs.SessionManager

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	portNumber = os.Getenv("PORT")
	if portNumber == "" {
		portNumber = ":8080"
	} else {
		portNumber = ":" + portNumber
	}

	app.InProduction = false
	app.UseCache = false

	session = scs.New()
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	http.HandleFunc("/", repo.Home)
	http.HandleFunc("/about", repo.About)

	fmt.Println("start listen to port ", portNumber)

	mux := routes(app)
	_ = http.ListenAndServe(portNumber, mux)
}
