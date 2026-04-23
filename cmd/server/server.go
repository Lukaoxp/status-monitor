package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Lukaoxp/status-monitor/internal/health"
)

// 1. Struct that holds our dependencies
type Server struct {
	healthService *health.Service
}

// 2. (s *Server) is the receiver. It binds this function to the Server struct.
// This lets us use 's.healthService' here without creating it from scratch!
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
	if _, err := w.Write(response); err != nil {
		log.Printf("error writing response: %v", err)
	}
}
