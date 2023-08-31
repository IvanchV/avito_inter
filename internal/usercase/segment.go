package usercase

import "avito/internal/models"

type SegmentRepo interface {
	CreateSegment(*models.Segment) error
	DeleteSegment(string) error
}

type SegmentService struct {
	repo SegmentRepo
}

func NewSegmentService(repo SegmentRepo) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) CreateSegment(seg *models.Segment) error {
	return s.repo.CreateSegment(seg)
}

func (s *SegmentService) DeleteSegment(name string) error {
	return s.repo.DeleteSegment(name)
}
