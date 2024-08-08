package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/fabiomattes2016/go-crud/interfaces"
	"github.com/fabiomattes2016/go-crud/internal/application"
	"github.com/fabiomattes2016/go-crud/internal/infrastructure/db"
	"github.com/fabiomattes2016/go-crud/internal/infrastructure/persistence"
	"github.com/fabiomattes2016/go-crud/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	enviroment := os.Getenv("ENVIRONMENT")

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	db.Init(db_host, db_user, db_pass, db_name, db_port)

	secret := os.Getenv("JWT_SECRET")

	userRepository := persistence.NewUserRepository(db.DB)
	userService := application.NewUserService(userRepository)
	authService := application.NewAuthService(userRepository, secret)
	userHandler := interfaces.NewUserHandler(userService)
	authHandler := interfaces.NewAuthHandler(authService)

	if enviroment == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	authorized := r.Group("/")
	authorized.Use(middleware.JWTAuthMiddleware(secret))
	{
		authorized.GET("/users", userHandler.GetUsers)
		authorized.GET("/users/:id", userHandler.GetUser)
		authorized.POST("/users", userHandler.CreateUser)
		authorized.PUT("/users/:id", userHandler.UpdateUser)
		authorized.DELETE("/users/:id", userHandler.DeleteUser)

		// Admin-only route
		admin := authorized.Group("/admin")
		admin.Use(middleware.RBACMiddleware("admin"))
		{
			admin.GET("/admin-route", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "Welcome, admin!"})
			})
		}
	}

	r.Run()
}
