# ==========
# Build stage
# ==========
FROM golang:1.24 AS builder

# On travaille dans /app
WORKDIR /app

# 1) On copie le module du serveur (go.mod/go.sum du dossier server)
COPY server/go.mod server/go.sum ./server/

# On va dans server et on télécharge les dépendances
WORKDIR /app/server
RUN go mod download

# 2) On copie le reste du code (api, etc.)
WORKDIR /app
COPY . .

# On reviens dans server et on build le binaire
WORKDIR /app/server
RUN CGO_ENABLED=0 GOOS=linux go build -o library-server .

# ==========
# Runtime stage
# ==========
FROM alpine:latest

WORKDIR /app

# On copie juste le binaire compilé
COPY --from=builder /app/server/library-server /app/library-server

# Ports :
# - 50051 = gRPC
# - 2112 = Prometheus metrics
EXPOSE 50051 2112

# On lance le serveur
CMD ["/app/library-server"]
