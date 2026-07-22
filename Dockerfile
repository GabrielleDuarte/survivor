# Etapa de build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Dependências
COPY go.mod go.sum ./
RUN go mod download

# Código
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Imagem final
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]