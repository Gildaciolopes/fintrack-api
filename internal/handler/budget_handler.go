package handler

import (
	"net/http"
	"time"

	"github.com/Gildaciolopes/fintrack-api/internal/middleware"
	"github.com/Gildaciolopes/fintrack-api/internal/models"
	"github.com/Gildaciolopes/fintrack-api/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
 
type BudgetHandler struct {
	repo *repository.BudgetRepository
}
 
func NewBudgetHandler(repo *repository.BudgetRepository) *BudgetHandler {
	return &BudgetHandler{repo: repo}
}
 
func (h *BudgetHandler) Create(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	var req models.CreateBudgetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Error:   "Invalid request data",
			Message: err.Error(),
		})
		return
	}

	budget := &models.Budget{
		UserID:     userID,
		CategoryID: req.CategoryID,
		Amount:     req.Amount,
		Month:      req.Month,
	}

	if err := h.repo.Create(budget); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Failed to create budget",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "Budget created successfully",
		Data:    budget,
	})
}
 
func (h *BudgetHandler) GetAll(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	var month *time.Time
	if monthStr := c.Query("month"); monthStr != "" {
		parsedMonth, err := time.Parse("2006-01-02", monthStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{
				Success: false,
				Error:   "Invalid month format (use YYYY-MM-DD)",
			})
			return
		}
		month = &parsedMonth
	}

	budgets, err := h.repo.GetAll(userID, month)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Failed to retrieve budgets",
			Message: err.Error(),
		})
		return
	}

	if budgets == nil {
		budgets = []models.Budget{}
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data:    budgets,
	})
}
 
func (h *BudgetHandler) GetBudgetsWithSpent(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	monthStr := c.Query("month")
	if monthStr == "" {
		monthStr = time.Now().Format("2006-01-02")
	}

	month, err := time.Parse("2006-01-02", monthStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Error:   "Invalid month format (use YYYY-MM-DD)",
		})
		return
	}

	budgets, err := h.repo.GetBudgetsWithSpent(userID, month)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Failed to retrieve budgets",
			Message: err.Error(),
		})
		return
	}

	if budgets == nil {
		budgets = []models.BudgetWithSpent{}
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data:    budgets,
	})
}
 
func (h *BudgetHandler) GetByID(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Error:   "Invalid budget ID",
		})
		return
	}

	budget, err := h.repo.GetByID(id, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Success: false,
			Error:   "Budget not found",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data:    budget,
	})
}
 
func (h *BudgetHandler) Update(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Error:   "Invalid budget ID",
		})
		return
	}

	var req models.UpdateBudgetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Error:   "Invalid request data",
			Message: err.Error(),
		})
		return
	}

	if err := h.repo.Update(id, userID, req.Amount, req.Month); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Failed to update budget",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Budget updated successfully",
	})
}
 
func (h *BudgetHandler) Delete(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Error:   "Invalid budget ID",
		})
		return
	}

	if err := h.repo.Delete(id, userID); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Failed to delete budget",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Budget deleted successfully",
	})
}
