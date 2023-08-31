package api

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SetUserSegmentsRequest struct {
	AddSegments    []string `json:"add"`
	RemoveSegments []string `json:"remove"`
	UserID         int      `json:"user_id"` // TODO: add TTL
}

func (h *Handler) SetUserSegments(c *gin.Context) {
	req := SetUserSegmentsRequest{}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()

	resp, err := h.segmentationService.SetUserSegments(ctx, req.AddSegments, req.RemoveSegments, req.UserID)
	log.Println("resp:", resp)
	if err != nil {
		c.JSON(200, resp)
		return
	}

	c.Status(http.StatusOK)
}
