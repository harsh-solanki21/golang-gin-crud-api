package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/harsh-solanki21/golang-gin-crud-api/configs"
	"github.com/harsh-solanki21/golang-gin-crud-api/controllers"
	"github.com/harsh-solanki21/golang-gin-crud-api/middlewares"
	"github.com/harsh-solanki21/golang-gin-crud-api/repositories"
	"github.com/harsh-solanki21/golang-gin-crud-api/routes"
	"github.com/harsh-solanki21/golang-gin-crud-api/services"
)

func main() {
	// Load environment variables
	if err := configs.LoadEnv(); err != nil {
		log.Fatal("Error loading env file:", err)
	}

	// Set Gin mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create a new router
	router := gin.Default()

	// Add necessary middleware
	router.Use(middlewares.ErrorMiddleware())

	// Set trusted proxies
	trustedProxies := os.Getenv("TRUSTED_PROXIES")
	proxyList := strings.Split(trustedProxies, ",")

	if err := router.SetTrustedProxies(proxyList); err != nil {
		log.Fatalf("Error setting trusted proxies: %v", err)
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := configs.ConnectDB(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal("Error disconnecting from MongoDB:", err)
		}
	}()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(client)
	productRepo := repositories.NewProductRepository(client)

	// Initialize services
	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)
	productService := services.NewProductService(productRepo)

	// Initialize controllers
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)
	productController := controllers.NewProductController(productService)

	// Set up routes
	routes.SetupRoutes(router, authController, userController, productController)

	// Get port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Add log before starting the server
	log.Printf("Server is running on http://localhost:%s\n", port)

	// Run the server
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
