package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Lukaoxp/status-monitor/internal/health"
)

const APIVersion = "1.0.0"

func main() {
	newService := health.NewService(APIVersion)
	server := &Server{healthService: newService}

	http.HandleFunc("/status", server.healthHandler)

	port := getEnv("PORT", "8080")

	log.Printf("Server starting on localhost:%s...", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

// 1. Definimos a struct que guardará nossas dependências
type Server struct {
	healthService *health.Service
}

// 2. O (s *Server) é o receiver. Ele diz que esta função pertence ao Server.
// Agora podemos usar 's.healthService' aqui dentro sem criá-lo do zero!
func (s *Server) healthHandler(w http.ResponseWriter, _ *http.Request) {

	// service := health.NewService(APIVersion)
	data := s.healthService.GetStatus()
	response, err := json.Marshal(data)
	if err != nil {
		log.Printf("error marshalling health status: %v", err)
		fmt.Println(err.Error())

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
