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

// Fluxo principal:
// - Inicializa o serviço de health com versionamento
// - Cria o handler HTTP injetando dependências
// - Lê a porta da variável de ambiente (com fallback)
// - Inicia o servidor HTTP
func main() {

	// 1. Setup de Dependências (Clean Architecture)
	newService := health.NewService(APIVersion)
	server := &Server{healthService: newService}

	// 2. Configuração do Servidor
	// 2.1. Registra as rotas ANTES de ligar o servidor
	http.HandleFunc("/status", server.healthHandler)

	// 2.2. Cria a configuração do servidor
	srv := returnHttpServer()

	// 3. Preparação dos Sinais (O "Ouvido" do sistema)
	sigChan := make(chan os.Signal, 1)
	// ESSA LINHA É A CHAVE: Liga o canal aos sinais do SO
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 4. Execução Assíncrona (A Goroutine)
	go func() {
		// http.ErrServerClosed é um erro "bom", significa que o shutdown funcionou
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error: %v\n", err)
		}
	}()

	// 5. O Bloqueio (Esperando a chamada do SO)
	<-sigChan
	log.Println("Stop signal received...")

	// 6. O Graceful Shutdown (O Prazo de validade)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Limpa o timer da memória ao terminar

	log.Println("Shutting down server...")

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server shutdown successfull")
}

// Retorna um ponteiro de um novo httserver
func returnHttpServer() (srv *http.Server) {
	port := getEnv("PORT", "8080")

	srv = &http.Server{
		Addr:    ":" + port,
		Handler: nil,
	}

	log.Printf("Server starting on localhost:%s...", port)
	return
}
