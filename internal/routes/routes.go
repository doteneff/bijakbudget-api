package routes

import (
	"github.com/doteneff/bijakbudget-api/internal/config"
	"github.com/doteneff/bijakbudget-api/internal/handlers"
	"github.com/doteneff/bijakbudget-api/internal/repositories"
	"github.com/doteneff/bijakbudget-api/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	db := config.DB

	// Categories Layer injection Setup
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Transactions Layer injection Setup
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// Users Layer injection Setup
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Group routes under /api/v1
	api := r.Group("/api/v1")
	{
		// Categories routes
		categories := api.Group("/categories")
		{
			categories.POST("/", categoryHandler.Create)
			categories.GET("/", categoryHandler.GetAll)
			categories.GET("/:id", categoryHandler.GetByID)
			categories.PUT("/:id", categoryHandler.Update)
			categories.DELETE("/:id", categoryHandler.Delete)
		}

		// Transactions routes
		transactions := api.Group("/transactions")
		{
			transactions.POST("/", transactionHandler.Create)
			transactions.GET("/", transactionHandler.GetAll)
			transactions.GET("/:id", transactionHandler.GetByID)
			transactions.PUT("/:id", transactionHandler.Update)
			transactions.DELETE("/:id", transactionHandler.Delete)
		}

		// Users routes
		users := api.Group("/users")
		{
			users.POST("/register", userHandler.Register)
			users.POST("/login", userHandler.Login)
			users.POST("/sso/google", userHandler.SSOGoogle)
			users.GET("/:id", userHandler.GetByID)
		}
	}
}
