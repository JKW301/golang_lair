# golang_lair
`go run main.go`
## main_docker.go
```
docker build -t go-docker-example .
docker run -p 8080:8080 go-docker-example
```
Tester avec `curl http://localhost:8080`

Pour réduire encore plus la taille de l'image, ajoutez scratch au `Dockerfile`, une image de base vide. Cela ne fonctionne que si l'application Go n'a pas besoin de bibliothèques système externes :

```
# Construire le binaire statique
FROM golang:1.21 AS builder
WORKDIR /app
COPY main_docker.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Créer une image minimale
FROM scratch
COPY --from=builder /app/app /
EXPOSE 8080
CMD ["/app"]
Avec cette méthode, l'image finale contiendra uniquement le binaire et pèsera seulement quelques mégaoctets.
```