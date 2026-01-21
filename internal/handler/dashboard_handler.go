package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Gildaciolopes/fintrack-api/internal/middleware"
	"github.com/Gildaciolopes/fintrack-api/internal/models"
	"github.com/Gildaciolopes/fintrack-api/internal/repository"
	"github.com/gin-gonic/gin"
)
 
type DashboardHandler struct {
	dashboardRepo   *repository.DashboardRepository
	transactionRepo *repository.TransactionRepository
}
 
func NewDashboardHandler(
	dashboardRepo *repository.DashboardRepository,
	transactionRepo *repository.TransactionRepository,
) *DashboardHandler {
	return &DashboardHandler{
		dashboardRepo:   dashboardRepo,
		transactionRepo: transactionRepo,
	}
}
 
func (h *DashboardHandler) GetStats(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
 
	startDate := time.Now().AddDate(0, 0, -30)
	endDate := time.Now()

	if start := c.Query("start_date"); start != "" {
		if parsed, err := time.Parse("2006-01-02", start); err == nil {
			startDate = parsed
		}
	}

	if end := c.Query("end_date"); end != "" {
		if parsed, err := time.Parse("2006-01-02", end); err == nil {
			endDate = parsed
		}
	}

	stats, err := h.dashboardRepo.GetStats(userID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Failed to retrieve stats",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data:    stats,
	})
}
 
func (h *DashboardHandler) GetExpensesByCategory(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	startDate := time.Now().AddDate(0, 0, -30)
	endDate := time.Now()

	if start := c.Query("start_date"); start != "" {
		if parsed, err := time.Parse("2006-01-02", start); err == nil {
			startDate = parsed
		}
	}

	if end := c.Query("end_date"); end != "" {
		if parsed, err := time.Parse("2006-01-02", end); err == nil {
			endDate = parsed
		}
	}

	expenses, err := h.dashboardRepo.GetExpensesByCategory(userID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Failed to retrieve expenses",
			Message: err.Error(),
		})
		return
	}

	if expenses == nil {
		expenses = []models.CategoryExpense{}
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data:    expenses,
	})
}
 
func (h *DashboardHandler) GetMonthlyData(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	months := 6
	if m := c.Query("months"); m != "" {
		if parsed, err := strconv.Atoi(m); err == nil && parsed > 0 && parsed <= 12 {
			months = parsed
		}
	}

	monthlyData, err := h.dashboardRepo.GetMonthlyData(userID, months)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Failed to retrieve monthly data",
			Message: err.Error(),
		})
		return
	}

	if monthlyData == nil {
		monthlyData = []models.MonthlyData{}
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data:    monthlyData,
	})
}
 
func (h *DashboardHandler) GetDailyData(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	startDate := time.Now().AddDate(0, 0, -30)
	endDate := time.Now()

	if start := c.Query("start_date"); start != "" {
		if parsed, err := time.Parse("2006-01-02", start); err == nil {
			startDate = parsed
		}
	}

	if end := c.Query("end_date"); end != "" {
		if parsed, err := time.Parse("2006-01-02", end); err == nil {
			endDate = parsed
		}
	}

	dailyData, err := h.dashboardRepo.GetDailyData(userID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Failed to retrieve daily data",
			Message: err.Error(),
		})
		return
	}

	if dailyData == nil {
		dailyData = []models.DailyData{}
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data:    dailyData,
	})
}
 
func (h *DashboardHandler) GetRecentTransactions(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	limit := 10
	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 && parsed <= 50 {
			limit = parsed
		}
	}

	transactions, err := h.transactionRepo.GetRecentTransactions(userID, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Failed to retrieve transactions",
			Message: err.Error(),
		})
		return
	}

	if transactions == nil {
		transactions = []models.Transaction{}
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data:    transactions,
	})
}
