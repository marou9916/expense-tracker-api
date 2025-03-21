package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marou9916/expense-tracker-api.git/controllers"
	"github.com/marou9916/expense-tracker-api.git/middlewares"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	authRoutes := router.Group("/auth")
	authRoutes.Use(middlewares.JWTRequired())
	{
		authRoutes.POST("/register", controllers.RegisterHandler)
		authRoutes.POST("/login", controllers.LoginHandler)
		authRoutes.POST("/logout", controllers.LogoutHandler)
		authRoutes.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusNotFound, gin.H{"erreur": "Page not found"})
		})
	}

	expenseRoutes := router.Group("/:idUser/expenses")
	{
		expenseRoutes.POST("/", controllers.AddExpense)
		expenseRoutes.GET("/:idExpense", controllers.GetExpense)
		expenseRoutes.DELETE("/:idExpense", controllers.DeleteExpense)
	}

	return router
}
