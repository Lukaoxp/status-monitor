# Copilot instructions for status-monitor

## Visão geral do projeto
- API HTTP mínima em Go com entrada em [cmd/server/main.go](cmd/server/main.go) e domínio em [internal/health/health.go](internal/health/health.go).
- Endpoint único `/status` retorna JSON com `status`, `version` e `uptime`.
- `uptime` é calculado a partir de `StartTime` inicializado em `init()` no pacote `health`.

## Fluxo principal e padrões locais
- O handler HTTP está em `healthHandler` e usa `encoding/json` + `http.Error` para falhas de marshal.
- Configuração de porta via `getEnv("PORT", "8080")` (fallback explícito).
- O pacote `health` expõe `type Service` com `Version` e método `GetStatus()`.
  - A intenção atual é injetar `Version` no `main` via `health.NewService(v)` e chamar `service.GetStatus()`.
  - Observe que `main.go` ainda chama `health.GetStatus()` (não existe função livre). Ajustes devem seguir o padrão `Service`.

## Dependências e integrações
- Apenas biblioteca padrão; nenhum framework web ou lib externa.
- Comunicação é HTTP síncrona usando `net/http`.

## Fluxos de trabalho
- Executar localmente (padrão Go): `go run ./cmd/server`.
- Build: `go build ./cmd/server` (gera binário do servidor).

## Convenções específicas
- Manter separação de camadas: HTTP em `cmd/server`, lógica de domínio em `internal/`.
- Evitar DI de frameworks; usar structs simples e construtores (`NewService`).
- Manter erros explícitos (ex.: checar `json.Marshal` e responder `500`).

## Diretrizes de abordagem (obrigatório)
- Siga **exatamente** o conteúdo abaixo.
- **Quando o usuário mandar**, atualize o status da construção (mesmo que não haja mudanças no código).

### Status da Construção: Status Monitor API 🐹
- Arquitetura do Projeto: separação clara entre o ponto de entrada ([cmd/server/main.go](cmd/server/main.go)) e a lógica de domínio ([internal/health/health.go](internal/health/health.go)). ✅
- Cálculo de Uptime: rastreamento com `time.Since` e captura de `startTime` privado em `init()` no pacote `health`. ✅
- Infraestrutura Web: servidor HTTP nativo com `net/http` e `encoding/json`, erros tratados manualmente. ✅
- Configuração Dinâmica: leitura de `PORT` com fallback explícito (`getEnv`). ✅
- Injeção de Dependência: `Service` com `version` privado, injetado via `NewService(v string)` e acessível apenas através de `GetStatus()`. ✅
- Encapsulamento: campos privados (`version`, `startTime`), métodos públicos (`NewService`, `GetStatus`). ✅
- Testes Unitários: suite de testes em andamento para `health.Service`. 🔄

> Última atualização: 03/02/2026

### Diretrizes de Abordagem: Como Você Deve Ajudar
- Foco em Produção: evoluir com resiliência, performance e prontidão para deploy (Docker/Cloud).
- Ponte de Conhecimento: traçar paralelos com C# (ex.: DI e tratamento de exceções).
- Aprendizado Ativo: aplicar método socrático, incentivando exploração da biblioteca padrão do Go.
