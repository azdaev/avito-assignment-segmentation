package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateSegmentRequest struct {
	Name string `json:"name"`
}

func (h *Handler) CreateSegment(c *gin.Context) {
	req := CreateSegmentRequest{}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()

	if err := h.segmentationService.CreateSegment(ctx, req.Name); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
