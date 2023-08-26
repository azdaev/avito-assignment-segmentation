package api

import (
	"context"

	"github.com/gin-gonic/gin"
)

type SegmentationService interface {
	CreateSegment(ctx context.Context, segmentName string) error
	DeleteSegment(ctx context.Context, segmentName string) error
	SetUserSegments(ctx context.Context, addSegments, removeSegments []string, userID int) error
	GetUserSegments(ctx context.Context, userID int) ([]string, error)
}

type Logger interface {
	Info(string, ...any)
}

type Handler struct {
	segmentationService SegmentationService
	logger              Logger
}

func New(segmentationService SegmentationService, logger Logger) *gin.Engine {
	h := Handler{
		segmentationService: segmentationService,
		logger:              logger,
	}

	r := gin.New()

	api := r.Group("/api/")

	api.POST("/create", h.CreateSegment)
	api.DELETE("/delete/", h.DeleteSegment)
	api.POST("/set", h.SetUserSegments)
	api.GET("/user/:id", h.GetUserSegments)

	return r
}
