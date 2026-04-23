package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Lukaoxp/status-monitor/internal/health"
)

const APIVersion = "1.0.0"

// Main flow:
// - Initializes the health service with versioning
// - Creates the HTTP handler injecting dependencies
// - Reads the port from environment variable (with fallback)
// - Starts the HTTP server
func main() {

	// 1. Dependency Setup (Clean Architecture)
	newService := health.NewService(APIVersion)
	server := &Server{healthService: newService}

	// 2. Server Configuration
	// 2.1. Register routes BEFORE starting the server
	http.HandleFunc("/status", server.healthHandler)

	// 2.2. Create the server configuration
	srv := returnHttpServer()

	// 3. Signal Setup (The system's "ear")
	sigChan := make(chan os.Signal, 1)
	// THIS LINE IS KEY: Connects the channel to OS signals
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 4. Async Execution (The Goroutine)
	go func() {
		// http.ErrServerClosed is a "good" error — it means shutdown worked
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error: %v\n", err)
		}
	}()

	// 5. The Block (Waiting for the OS signal)
	<-sigChan
	log.Println("Stop signal received...")

	// 6. Graceful Shutdown (The deadline)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Frees the timer from memory when done

	log.Println("Shutting down server...")

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server shutdown successful")
}

// Returns a pointer to a new http server
func returnHttpServer() (srv *http.Server) {
	port := getEnv("PORT", "8080")

	srv = &http.Server{
		Addr:              ":" + port,
		Handler:           nil,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Server starting on localhost:%s...", port)
	return
}
