package handler

import (
	"net/http"

	"github.com/Gildaciolopes/fintrack-api/internal/middleware"
	"github.com/Gildaciolopes/fintrack-api/internal/models"
	"github.com/Gildaciolopes/fintrack-api/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
 
type GoalHandler struct {
	repo *repository.GoalRepository
}
 
func NewGoalHandler(repo *repository.GoalRepository) *GoalHandler {
	return &GoalHandler{repo: repo}
}
 
func (h *GoalHandler) Create(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	var req models.CreateGoalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Error:   "Invalid request data",
			Message: err.Error(),
		})
		return
	}

	goal := &models.FinancialGoal{
		UserID:        userID,
		Title:         req.Title,
		TargetAmount:  req.TargetAmount,
		CurrentAmount: req.CurrentAmount,
		Deadline:      req.Deadline,
	}

	if err := h.repo.Create(goal); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Failed to create goal",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "Goal created successfully",
		Data:    goal,
	})
}
 
func (h *GoalHandler) GetAll(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	status := c.Query("status")

	goals, err := h.repo.GetAll(userID, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Failed to retrieve goals",
			Message: err.Error(),
		})
		return
	}

	if goals == nil {
		goals = []models.FinancialGoal{}
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data:    goals,
	})
}
 
func (h *GoalHandler) GetByID(c *gin.Context) {
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
			Error:   "Invalid goal ID",
		})
		return
	}

	goal, err := h.repo.GetByID(id, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Success: false,
			Error:   "Goal not found",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data:    goal,
	})
}
 
func (h *GoalHandler) Update(c *gin.Context) {
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
			Error:   "Invalid goal ID",
		})
		return
	}

	var req models.UpdateGoalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Error:   "Invalid request data",
			Message: err.Error(),
		})
		return
	}

	updates := make(map[string]interface{})
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.TargetAmount > 0 {
		updates["target_amount"] = req.TargetAmount
	}
	if req.CurrentAmount >= 0 {
		updates["current_amount"] = req.CurrentAmount
	}
	if req.Deadline != nil {
		updates["deadline"] = req.Deadline
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}

	if err := h.repo.Update(id, userID, updates); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Failed to update goal",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Goal updated successfully",
	})
}
 
func (h *GoalHandler) Delete(c *gin.Context) {
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
			Error:   "Invalid goal ID",
		})
		return
	}

	if err := h.repo.Delete(id, userID); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Failed to delete goal",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Goal deleted successfully",
	})
}
 
func (h *GoalHandler) Contribute(c *gin.Context) {
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
			Error:   "Invalid goal ID",
		})
		return
	}

	var req models.ContributeGoalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Error:   "Invalid request data",
			Message: err.Error(),
		})
		return
	}

	if err := h.repo.Contribute(id, userID, req.Amount); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Failed to contribute to goal",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Contribution added successfully",
	})
}
