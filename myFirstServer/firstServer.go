package myFirstServer

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var tempStorage = make(map[string]string)

func MyFirstServer() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", sayHello)

	r.Put("/greeting/{text}", updateGreeting)

	tempStorage["greeting"] = "hello students!"

	http.ListenAndServe(":3000", r)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(tempStorage["greeting"]))
}

func updateGreeting(w http.ResponseWriter, r *http.Request) {
	greeting := chi.URLParam(r, "text")

	tempStorage["greeting"] = greeting

	w.Write([]byte("успех!"))
}
