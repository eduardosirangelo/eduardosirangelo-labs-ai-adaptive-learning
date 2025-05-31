# Stage 1: Build binary
FROM golang:1.23-alpine AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia apenas o go.mod e go.sum para aproveitar cache de dependências
COPY go.mod go.sum ./

# Baixa módulos do Go (cache)
RUN go mod download

# Copia o restante do código fonte do serviço
COPY . .

# Compila o binário estático e otimizado
# - CGO\_ENABLED=0: sem dependências C, gera executável estático
# - GOOS=linux: alvo Linux
# - -trimpath: remove caminhos do sistema do binário
RUN CGO\_ENABLED=0 GOOS=linux go build -trimpath -o assessment-api ./cmd/assessment-api

# Stage 2: Runtime minimalista
FROM scratch

# Copia certificados CA para TLS (caso o serviço faça chamadas HTTPS)
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Copia o executável compilado
COPY --from=builder /app/assessment-api /usr/local/bin/assessment-api

# Porta exposta pelo serviço (configurável via env var)
EXPOSE 8070

# Define o usuário não-root (opcional, aumente segurança)
# USER 65532:65532

# Inicia o serviço
ENTRYPOINT ["/usr/local/bin/assessment-api"]
