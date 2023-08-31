package api

import (
	"context"

	"github.com/gin-gonic/gin"
)

type CreateSegmentRequest struct {
	Name       string `json:"name"`
	Percentage int    `json:"percentage"`
}

func (h *Handler) CreateSegment(c *gin.Context) {
	req := CreateSegmentRequest{}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()

	affectedUsers, err := h.segmentationService.CreateSegment(ctx, req.Name, req.Percentage)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if len(affectedUsers) > 0 {
		c.JSON(200, gin.H{"affected_users": affectedUsers})
		return
	}

	c.Status(200)
}
