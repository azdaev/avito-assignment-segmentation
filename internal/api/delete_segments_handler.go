package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteSegmentRequest struct {
	Name string `json:"name"`
}

func (h *Handler) DeleteSegment(c *gin.Context) {
	req := DeleteSegmentRequest{}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()

	if err := h.segmentationService.DeleteSegment(ctx, req.Name); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
