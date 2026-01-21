package main

import (
	"log"

	"github.com/Gildaciolopes/fintrack-api/internal/config"
	"github.com/Gildaciolopes/fintrack-api/internal/handler"
	"github.com/Gildaciolopes/fintrack-api/internal/middleware"
	"github.com/Gildaciolopes/fintrack-api/internal/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const VERSION = "1.0.0"
 func main() { 
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
 
	db, err := cfg.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
 
	categoryRepo := repository.NewCategoryRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	goalRepo := repository.NewGoalRepository(db)
	budgetRepo := repository.NewBudgetRepository(db)
	dashboardRepo := repository.NewDashboardRepository(db)
 
	healthHandler := handler.NewHealthHandler(VERSION)
	categoryHandler := handler.NewCategoryHandler(categoryRepo)
	transactionHandler := handler.NewTransactionHandler(transactionRepo)
	goalHandler := handler.NewGoalHandler(goalRepo)
	budgetHandler := handler.NewBudgetHandler(budgetRepo)
	dashboardHandler := handler.NewDashboardHandler(dashboardRepo, transactionRepo)
 
	authMiddleware := middleware.NewAuthMiddleware(cfg.Supabase.JWTSecret)
 
	if cfg.Server.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.ErrorHandler())

	// CORS configuration
	corsConfig := cors.Config{
		AllowOrigins:     cfg.CORS.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}
	r.Use(cors.New(corsConfig))
 
	r.GET("/health", healthHandler.Health)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "FinTrack API",
			"version": VERSION,
			"status":  "running",
		})
	})

	// API routes
	api := r.Group("/api/" + cfg.Server.APIVersion)
	{   
		protected := api.Group("")
		protected.Use(authMiddleware.Authenticate())
		{ 
			dashboard := protected.Group("/dashboard")
			{
				dashboard.GET("/stats", dashboardHandler.GetStats)
				dashboard.GET("/expenses-by-category", dashboardHandler.GetExpensesByCategory)
				dashboard.GET("/monthly-data", dashboardHandler.GetMonthlyData)
				dashboard.GET("/daily-data", dashboardHandler.GetDailyData)
				dashboard.GET("/recent-transactions", dashboardHandler.GetRecentTransactions)
			}
 
			categories := protected.Group("/categories")
			{
				categories.POST("", categoryHandler.Create)
				categories.GET("", categoryHandler.GetAll)
				categories.GET("/:id", categoryHandler.GetByID)
				categories.PUT("/:id", categoryHandler.Update)
				categories.DELETE("/:id", categoryHandler.Delete)
			}
 
			transactions := protected.Group("/transactions")
			{
				transactions.POST("", transactionHandler.Create)
				transactions.GET("", transactionHandler.GetAll)
				transactions.GET("/:id", transactionHandler.GetByID)
				transactions.PUT("/:id", transactionHandler.Update)
				transactions.DELETE("/:id", transactionHandler.Delete)
			}
 
			goals := protected.Group("/goals")
			{
				goals.POST("", goalHandler.Create)
				goals.GET("", goalHandler.GetAll)
				goals.GET("/:id", goalHandler.GetByID)
				goals.PUT("/:id", goalHandler.Update)
				goals.DELETE("/:id", goalHandler.Delete)
				goals.POST("/:id/contribute", goalHandler.Contribute)
			}
 
			budgets := protected.Group("/budgets")
			{
				budgets.POST("", budgetHandler.Create)
				budgets.GET("", budgetHandler.GetAll)
				budgets.GET("/with-spent", budgetHandler.GetBudgetsWithSpent)
				budgets.GET("/:id", budgetHandler.GetByID)
				budgets.PUT("/:id", budgetHandler.Update)
				budgets.DELETE("/:id", budgetHandler.Delete)
			}
		}
	}
 
	addr := ":" + cfg.Server.Port
	log.Printf("🚀 Server starting on port %s (env: %s)", cfg.Server.Port, cfg.Server.Env)
	log.Printf("📚 API documentation available at http://localhost:%s/api/%s", cfg.Server.Port, cfg.Server.APIVersion)

	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
