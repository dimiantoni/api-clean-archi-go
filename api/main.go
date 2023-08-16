package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dimiantoni/api-clean-archi-go/infra/database"
	"github.com/dimiantoni/api-clean-archi-go/infra/repository"
	"github.com/dimiantoni/api-clean-archi-go/usecase/user"

	"github.com/codegangsta/negroni"
	"github.com/dimiantoni/api-clean-archi-go/api/handler"
	"github.com/dimiantoni/api-clean-archi-go/api/middleware"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	userRepo := repository.NewUserRepository(database.NewMongodb())
	userService := user.NewService(userRepo)

	r := mux.NewRouter()

	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Auth),
		negroni.NewLogger(),
	)

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// User
	handler.MakeUserHandlers(r, *n, userService)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Panic(err.Error())
	}

}
