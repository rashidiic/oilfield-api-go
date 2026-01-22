package mock

import (
	"net/http"
	"strconv"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Handler struct {
	db        *gorm.DB
	validator *validator.Validate
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		db:        db,
		validator: validator.New(),
	}
}

func toResponse(item MockItem) MockItemResponse {
	return MockItemResponse{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		CreatedAt:   item.CreatedAt.UTC().Format(time.RFC3339),
	}
}

// POST /mock-items

// Create godoc
// @Summary      Create mock item
// @Description  Creates a new MockItem
// @Tags         mock-items
// @Accept       json
// @Produce      json
// @Param        body  body      CreateMockItemRequest  true  "Mock item payload"
// @Success      201   {object}  MockItemResponse
// @Failure      400   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /api/mock-items [post]
func (h *Handler) Create(c *gin.Context) {
	var req CreateMockItemRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid request payload",
			Details: err.Error(),
		})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Validation failed",
			Details: err.Error(),
		})
		return
	}

	item := MockItem{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := h.db.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "Failed to create mock item",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, toResponse(item))
}

// GET /mock-items/:id

// List godoc
// @Summary      List mock items
// @Description  Returns all MockItems
// @Tags         mock-items
// @Produce      json
// @Success      200  {array}   MockItemResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/mock-items [get]
func (h *Handler) List(c *gin.Context) {
	var items []MockItem

	if err := h.db.Order("id desc").Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "database error",
			Details: err.Error(),
		})
		return
	}

	resp := make([]MockItemResponse, 0, len(items))
	for _, item := range items {
		resp = append(resp, toResponse(item))
	}
	c.JSON(http.StatusOK, resp)
}

// PUT /mock-items/:id

// Update godoc
// @Summary      Update mock item
// @Description  Updates MockItem by id
// @Tags         mock-items
// @Accept       json
// @Produce      json
// @Param        id    path      int                   true  "Mock item ID"
// @Param        body  body      UpdateMockItemRequest  true  "Mock item payload"
// @Success      200   {object}  MockItemResponse
// @Failure      400   {object}  ErrorResponse
// @Failure      404   {object}  ErrorResponse
// @Failure      500   {object}  ErrorResponse
// @Router       /api/mock-items/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID parameter"})
		return
	}

	var req UpdateMockItemRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid json",
			Details: err.Error(),
		})
		return
	}

	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Validation failed",
			Details: err.Error(),
		})
		return
	}

	var item MockItem

	if err := h.db.First(&item, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "database error",
			Details: err.Error(),
		})
		return
	}

	item.Name = req.Name
	item.Description = req.Description

	if err := h.db.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "database error",
			Details: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, toResponse(item))
}

// DELETE /mock-items/:id

// Delete godoc
// @Summary      Delete mock item
// @Description  Deletes MockItem by id
// @Tags         mock-items
// @Produce      json
// @Param        id   path      int  true  "Mock item ID"
// @Success      204  "No Content"
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/mock-items/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid id"})
		return
	}

	res := h.db.Delete(&MockItem{}, id)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "database error",
			Details: res.Error.Error(),
		})
		return
	}

	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
