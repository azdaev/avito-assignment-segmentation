package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserSegments(c *gin.Context) {
	userIDstr := c.Param("id")[1:]
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user_id",
		})
		return
	}

	ctx := context.Background()
	segments, err := h.segmentationService.GetUserSegments(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"segments": segments,
	})
}
