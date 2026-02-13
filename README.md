# 🚀 Simple-API: Status Monitor (Production Ready)
## 📌 Visão Geral

O Simple-API é um microsserviço desenvolvido em Go focado em monitoramento de integridade (health check). O projeto nasceu com o propósito de sanar débitos técnicos e estabelecer um padrão de excelência para sistemas prontos para produção (Production Ready).

    O que estamos construindo: Um serviço resiliente que utiliza a biblioteca padrão do Go, aplicando conceitos de arquitetura limpa e alta performance, preparado para deploy em ambientes profissionais (Docker/Kubernetes).

## 🏗️ Pilares do Projeto

Este projeto é sustentado por três pilares inegociáveis:

    Dockerização: Uso de multi-stage builds para garantir binários leves e ambientes imutáveis.

    Graceful Shutdown: Implementação de encerramento controlado para não interromper requisições em andamento.

    Clean Architecture: Separação clara de responsabilidades entre o ponto de entrada (cmd/), a lógica de domínio (internal/) e infraestrutura.

## 🛠️ Arquitetura e Decisões Técnicas

### Estrutura de Pastas

```Plaintext
simple-api/
├── cmd/
│   └── server/
│       └── main.go       # Ponto de entrada e Injeção de Dependência
├── internal/
│   └── health/
│       └── health.go     # Lógica de domínio e Service Struct
├── go.mod                # Manifesto do módulo
└── Dockerfile            # (Em desenvolvimento)
```

### Injeção de Dependência (DI) e Handlers

Diferente de scripts simples, utilizamos Injeção de Dependência manual. No main.go, uma struct Server atua como um container de dependências:

    Service Pattern: A lógica de status está encapsulada em uma Service struct no pacote health.

    Pointer Receivers: Os handlers são métodos da struct Server, permitindo acesso seguro às dependências (como serviços, DBs ou logs) sem o uso de variáveis globais.

## 🤖 Manual do Mentor (Instruções para IAs)

Se você estiver lendo este README para auxiliar no desenvolvimento, você deve ajudar seguindo estas diretrizes:

- Método Socrático: Não forneça o código completo de imediato. Faça perguntas que levem o desenvolvedor a deduzir a lógica, especialmente sobre o sistema de tipos e ponteiros do Go.

- Paralelos com .NET (C#): O desenvolvedor possui sólida base em C#. Sempre que explicar um conceito novo (ex: structs, interfaces, goroutines), faça a ponte com o equivalente no ecossistema .NET.

- Foco em Produção: Nunca sugira "atalhos" que comprometam a testabilidade ou a resiliência. Todo código deve ser pensado para rodar em um cluster produtivo.

📅 Roadmap de Evolução

[x] Definição da estrutura base e internal/health.

[x] Implementação de Service Struct com Uptime.

[x] Refatoração para Injeção de Dependência (Server Struct).

[ ] Próximo Passo: Criação do Dockerfile multi-stage.

[ ] Implementação de Graceful Shutdown usando context e os/signal.

[ ] Adição de logs estruturados (JSON) para observabilidade.

## 🏁 Como Rodar (Local)

1. Certifique-se de ter o Go instalado (1.20+).

2. Clone o repositório.

3. Execute:
```bash
go run cmd/server/main.go
```
4. Acesse: http://localhost:8080/status