package mock

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(api *gin.RouterGroup, db *gorm.DB) {
	h := NewHandler(db)

	group := api.Group("mock-items")
	{
		group.POST("", h.Create)
		group.GET("", h.List)
		group.PUT("/:id", h.Update)
		group.DELETE("/:id", h.Delete)
	}
}
