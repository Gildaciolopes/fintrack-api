package handler

import (
	"net/http"

	"github.com/Gildaciolopes/fintrack-backend/internal/middleware"
	"github.com/Gildaciolopes/fintrack-backend/internal/models"
	"github.com/Gildaciolopes/fintrack-backend/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
 
type CategoryHandler struct {
	repo *repository.CategoryRepository
}
 
func NewCategoryHandler(repo *repository.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{repo: repo}
} 

func (h *CategoryHandler) Create(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	var req models.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Error:   "Invalid request data",
			Message: err.Error(),
		})
		return
	}

	category := &models.Category{
		UserID: userID,
		Name:   req.Name,
		Type:   req.Type,
		Color:  req.Color,
		Icon:   req.Icon,
	}

	if err := h.repo.Create(category); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Failed to create category",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "Category created successfully",
		Data:    category,
	})
}
 
func (h *CategoryHandler) GetAll(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	categoryType := c.Query("type")

	categories, err := h.repo.GetAll(userID, categoryType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Failed to retrieve categories",
			Message: err.Error(),
		})
		return
	}

	if categories == nil {
		categories = []models.Category{}
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data:    categories,
	})
}
 
func (h *CategoryHandler) GetByID(c *gin.Context) {
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
			Error:   "Invalid category ID",
		})
		return
	}

	category, err := h.repo.GetByID(id, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Success: false,
			Error:   "Category not found",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Data:    category,
	})
}
 
func (h *CategoryHandler) Update(c *gin.Context) {
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
			Error:   "Invalid category ID",
		})
		return
	}

	var req models.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Error:   "Invalid request data",
			Message: err.Error(),
		})
		return
	}

	updates := make(map[string]interface{})
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Type != "" {
		updates["type"] = req.Type
	}
	if req.Color != "" {
		updates["color"] = req.Color
	}
	if req.Icon != "" {
		updates["icon"] = req.Icon
	}

	if err := h.repo.Update(id, userID, updates); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Failed to update category",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Category updated successfully",
	})
}
 
func (h *CategoryHandler) Delete(c *gin.Context) {
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
			Error:   "Invalid category ID",
		})
		return
	}

	if err := h.repo.Delete(id, userID); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Failed to delete category",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Category deleted successfully",
	})
}
