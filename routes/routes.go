package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harsh-solanki21/golang-gin-crud-api/controllers"
	"github.com/harsh-solanki21/golang-gin-crud-api/middlewares"
)

func SetupRoutes(router *gin.Engine, authController *controllers.AuthController, userController *controllers.UserController, productController *controllers.ProductController) {
	// Public routes
	public := router.Group("/api/v1")
	{
		public.POST("/register", userController.CreateUser)
		public.POST("/login", authController.Login)
		public.POST("/logout", authController.Logout)
	}

	// Protected routes
	protected := router.Group("/api/v1")
	protected.Use(middlewares.AuthMiddleware())
	{
		// User routes group
		users := protected.Group("/users")
		{
			users.GET("/", middlewares.AuthorizeMiddleware("admin"), userController.ListUsers)
			users.GET("/:id", userController.GetUser)
			users.PUT("/:id", userController.UpdateUser)
			users.DELETE("/:id", userController.DeleteUser)
		}

		// Product routes group
		products := protected.Group("/products")
		{
			products.POST("/", productController.CreateProduct)
			products.GET("/", productController.ListProducts)
			products.GET("/:id", productController.GetProduct)
			products.PUT("/:id", productController.UpdateProduct)
			products.DELETE("/:id", productController.DeleteProduct)
		}
	}
}
