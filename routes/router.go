package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marou9916/expense-tracker-api.git/controllers"
	"github.com/marou9916/expense-tracker-api.git/middlewares"
)

// SetupRoutes initializes all API routes and returns the Gin router instance.
func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Authentication routes
	authRoutes := router.Group("/auth")
	{
		// User registration endpoint
		authRoutes.POST("/register", controllers.RegisterHandler)
		// User login endpoint (returns JWT token)
		authRoutes.POST("/login", controllers.LoginHandler)
		// User logout endpoint (requires valid JWT)
		authRoutes.POST("/logout", middlewares.JWTRequired(), controllers.LogoutHandler)
		// Fallback route for undefined authentication endpoints
		authRoutes.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		})
	}

	// Expense routes (require user authorization)
	expenseRoutes := router.Group("/expenses")

	expenseRoutes.Use(middlewares.JWTRequired())

	{
		// Add a new expense (requires authorization)
		expenseRoutes.POST("/", middlewares.CheckUserAuthorization(), controllers.AddExpense)
		// Get a specific expense by ID (requires authorization)
		expenseRoutes.GET("/:idExpense", middlewares.CheckUserAuthorization(), controllers.GetExpense)
		// Delete an expense by ID (requires authorization)
		expenseRoutes.DELETE("/:idExpense", middlewares.CheckUserAuthorization(), controllers.DeleteExpense)
	}

	return router
}
