package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	product_repository "github.com/williamkoller/golang-domain-driven-design/internal/domain/product/repository"
	product_handlers "github.com/williamkoller/golang-domain-driven-design/internal/infra/http/handlers"
	product_router "github.com/williamkoller/golang-domain-driven-design/internal/infra/http/router"
	shared_events "github.com/williamkoller/golang-domain-driven-design/internal/shared/domain/events"
)

func main() {
	dispatcher := shared_events.NewEventDispatcher()
	repo := product_repository.NewRepository()

	productHandler := product_handlers.NewProductHandler(repo, dispatcher)

	r := product_router.SetupProductRouter(productHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		fmt.Println("Server running on http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Failed to start server: %v\n", err)
		}
	}()

	GracefulShutdown(server, 5*time.Second)
}

func GracefulShutdown(server *http.Server, timeout time.Duration) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	fmt.Println("\nReceived termination signal. Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error during server shutdown: %v\n", err)
	} else {
		fmt.Println("Server shut down gracefully.")
	}
}
