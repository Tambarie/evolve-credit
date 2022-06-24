package server

import (
	"context"
	"fmt"
	"github.com/Tambarie/evolve-credit/application/handler"
	"github.com/Tambarie/evolve-credit/application/userRepo"
	"github.com/Tambarie/evolve-credit/domain/service"
	"github.com/Tambarie/evolve-credit/infrastructure/repository/postgresDB"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Start() {
	router := initializeRouter()

	db := postgresDB.Init()
	h := handler.Handler{
		UserService: service.NewUserService(userRepo.NewUserRepositoryDB(db)),
	}
	DefineRouter(router, &h)
	PORT := fmt.Sprintf(":%s", os.Getenv("service_port"))
	if PORT == ":" {
		PORT += "8080"
	}
	s := &http.Server{
		Handler: router,
		Addr:    PORT,
	}
	wait := make(chan os.Signal) // creates a channel that will be used to wait for a signal

	log.Printf("Server Started at Port%s", PORT)

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("An error occurred with the server: %s", err)
			return
		}
	}() // go routine to start the server
	// sends a signal to the wait channel if there is an interrupt signal
	signal.Notify(wait, os.Interrupt)

	<-wait // waits here until a signal is received
	log.Printf("Shutting down the server...")

	time.Sleep(time.Second * 2) // sleep for 1 second

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// shuts down the server gracefully
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("An error occurred: %s", err)
	}
	log.Printf("Server exits successfully")
}
