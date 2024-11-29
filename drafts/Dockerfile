# Étape 1 : Construire l'application
FROM golang:1.21 AS builder

# Définir le répertoire de travail dans le conteneur
WORKDIR /app

# Copier le code source dans le conteneur
COPY main_docker.go .

# Compiler l'application
RUN go build -o app .

# Étape 2 : Préparer l'image minimale
FROM alpine:latest

# Copier le binaire Go depuis l'étape précédente
COPY --from=builder /app/app /app

# Exposer le port de l'application
EXPOSE 8080

# Commande d'exécution
CMD ["/app"]
