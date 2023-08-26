package segments

import (
	"context"
)

type Repository interface {
	CreateSegment(ctx context.Context, name string) error
	DeleteSegment(ctx context.Context, name string) error
	AddUserSegments(ctx context.Context, addSegments []string, userID int) error
	RemoveUserSegments(ctx context.Context, removeSegments []string, userID int) error
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

func (s *Service) CreateSegment(ctx context.Context, segmentName string) error {
	return s.repo.CreateSegment(ctx, segmentName)
}

func (s *Service) DeleteSegment(ctx context.Context, segmentName string) error {
	return s.repo.DeleteSegment(ctx, segmentName)
}

func (s *Service) SetUserSegments(ctx context.Context, addSegments, removeSegments []string, userID int) error {
	err := s.repo.AddUserSegments(ctx, addSegments, userID)
	if err != nil {
		return err
	}
	return s.repo.RemoveUserSegments(ctx, removeSegments, userID)
}

func (s *Service) GetUserSegments(ctx context.Context, userID int) ([]string, error) {
	return s.repo.GetUserSegments(ctx, userID)
}
