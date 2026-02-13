# Estágio 1: Builder (Onde a mágica da compilação acontece)
FROM golang:1.26-alpine AS builder

# Definimos onde vamos trabalhar dentro do container
WORKDIR /app

# Copiamos o arquivo de definição de módulos
COPY go.mod ./

# Baixamos o que for necessário (no momento, nada externo, mas prepara o ambiente)
RUN go mod download

# Agora precisamos trazer os outros arquivos 
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/main.go


# Inicializa uma imagem vazia de SO
FROM scratch
WORKDIR /root/
# Puxa o arquivo server do estagio BUILDER anterior
COPY --from=builder /app/server .
# Liberamos a porta 8080
EXPOSE 8080
# Inicializamos o executavel
CMD [ "./server" ] 

# Pra buildar: docker build -t simple-api:v1 .
# Pra rodar: docker run -p 8080:8080 simple-api:v1