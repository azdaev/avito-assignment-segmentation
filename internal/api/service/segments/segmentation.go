package segments

import (
	"context"
	"errors"
	"segmentation-avito/internal/models"
)

type Repository interface {
	CreateSegment(ctx context.Context, name string) error
	SetPercentageSegments(ctx context.Context, name string, percentage int) ([]int, error)
	DeleteSegment(ctx context.Context, name string) error
	AddUserSegments(ctx context.Context, addSegments []string, userID int) ([]string, error)
	RemoveUserSegments(ctx context.Context, removeSegments []string, userID int) ([]string, error)
	GetUserSegments(ctx context.Context, userID int) ([]string, error)
}

type Service struct {
	repo Repository
}

func New(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateSegment(ctx context.Context, segmentName string, percentage int) ([]int, error) {
	err := s.repo.CreateSegment(ctx, segmentName)
	if err != nil {
		return nil, err
	}

	if percentage < 0 {
		return nil, nil
	}

	return s.repo.SetPercentageSegments(ctx, segmentName, percentage)
}

func (s *Service) DeleteSegment(ctx context.Context, segmentName string) error {
	return s.repo.DeleteSegment(ctx, segmentName)
}

func (s *Service) SetUserSegments(ctx context.Context, addSegments, removeSegments []string, userID int) (*models.SetUserSegmentsResponse, error) {
	var notAdded, notRemoved []string
	var notAddedErr, notRemovedErr error

	if len(addSegments) != 0 {
		notAdded, notAddedErr = s.repo.AddUserSegments(ctx, addSegments, userID)
	}

	if len(removeSegments) != 0 {
		notRemoved, notRemovedErr = s.repo.RemoveUserSegments(ctx, removeSegments, userID)
	}

	if notAddedErr == nil && notRemovedErr == nil {
		return nil, nil
	}

	return &models.SetUserSegmentsResponse{
		AddError:    notAdded,
		RemoveError: notRemoved,
	}, errors.New("SetUserSegments: invalid segments")
}

func (s *Service) GetUserSegments(ctx context.Context, userID int) ([]string, error) {
	return s.repo.GetUserSegments(ctx, userID)
}
