package main

import (
	"log"
	"net/http"

	"github.com/Lukaoxp/status-monitor/internal/health"
)

const APIVersion = "1.0.0"

// Fluxo principal:
// - Inicializa o serviço de health com versionamento
// - Cria o handler HTTP injetando dependências
// - Lê a porta da variável de ambiente (com fallback)
// - Inicia o servidor HTTP
func main() {

	newService := health.NewService(APIVersion)
	server := &Server{healthService: newService}

	http.HandleFunc("/status", server.healthHandler)

	port := getEnv("PORT", "8080")

	log.Printf("Server starting on localhost:%s...", port)

	// srv := http.Server{
	// 	Addr:    ":" + port,
	// 	Handler: nil,
	// }

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}

}
