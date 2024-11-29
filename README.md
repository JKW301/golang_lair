# golang_lair
`go run main.go`

# Dépendance
`go mod init web_console` **initialiser un module avant d'installer les dépendances**
`go install github.com/gin-gonic/gin@latest`
`go get github.com/gorilla/websocket`
`go get github.com/pquerna/otp{}`
`go get -u gorm.io/gorm`
`go get -u gorm.io/driver/sqlite`


##
**Servir la page html avec GIN**
```
r.LoadHTMLGlob("web/templates/*")
r.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", nil)
})

```

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