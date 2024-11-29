package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB // Instance de la base de données

// Initialisation de la base de données
func initDB() {
	db = InitializeDB()
}

// Gestionnaire pour l'inscription
func signupHandler(c *gin.Context) {
	type SignupForm struct {
		Email    string `form:"email" binding:"required,email"`
		Password string `form:"password" binding:"required,min=8"`
	}

	var form SignupForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Hash du mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	// Enregistrement dans la base
	user := User{Email: form.Email, Password: string(hashedPassword)}
	if err := db.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "User already exists"})
		return
	}

	c.JSON(200, gin.H{"message": "Account created successfully"})
}

// Gestionnaire pour la connexion
func loginHandler(c *gin.Context) {
	type LoginForm struct {
		Email    string `form:"email" binding:"required,email"`
		Password string `form:"password" binding:"required"`
	}

	var form LoginForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user User
	if err := db.Where("email = ?", form.Email).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "Invalid email or password"})
		return
	}

	// Vérification du mot de passe
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
		c.JSON(400, gin.H{"error": "Invalid email or password"})
		return
	}

	// Stocker l'état de connexion
	c.SetCookie("user_id", fmt.Sprint(user.ID), 3600, "/", "localhost", false, true)
	c.JSON(200, gin.H{"message": "Logged in successfully"})
}

// Middleware de protection
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := c.Cookie("user_id")
		if err != nil || userID == "" {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {
	// Initialisation de la base de données
	initDB()

	// Création d'un routeur Gin
	r := gin.Default()
	r.LoadHTMLGlob("../web/templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/graph", func(c *gin.Context) {
		c.HTML(http.StatusOK, "graph.html", nil)
	})		

	// Routes principales
	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", nil)
	})
	r.POST("/signup", signupHandler)

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", loginHandler)

	// Page protégée
	r.GET("/dashboard", authMiddleware(), func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", gin.H{"message": "Welcome to your dashboard!"})
	})

	// Démarrage du serveur
	r.Run(":8080")
}
