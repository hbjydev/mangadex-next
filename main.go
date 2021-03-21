package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gorilla/mux"
	"github.com/hbjydev/mangadex-next/database"
	"github.com/hbjydev/mangadex-next/middlewares"
	"github.com/hbjydev/mangadex-next/routers"
)

func main() {
	if os.Getenv("SENTRY_DSN") != "" {
		err := sentry.Init(sentry.ClientOptions{
			Dsn:              os.Getenv("SENTRY_DSN"),
			Environment:      os.Getenv("SENTRY_ENV"),
			Release:          os.Getenv("SENTRY_RELEASE"),
			AttachStacktrace: true,
			Debug:            true,
		})

		if err != nil {
			log.Fatalf("Error initializing Sentry integration: %v", err)
		}

		defer sentry.Flush(2 * time.Second)
	}
	log.Println("Starting Mangadex API service...")
	log.Println("Connecting to database...")
	database.Connect()
	log.Println("Connected!")

	r := mux.NewRouter()

	r.Use(middlewares.LogMiddleware)
	r.Use(middlewares.SentryMiddleware)
	r.Use(middlewares.JSONMiddleware)

	healthRouter := routers.HealthRouter{}
	healthRouter.RegisterRoutes(r)

	userRouter := routers.UserRouter{}
	userRouter.RegisterRoutes(r)

	log.Println("Starting server on :3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
