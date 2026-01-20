package healtcheck

import healthcheck "fiber-file-converter-api/internal/domain/health"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Check() healthcheck.Status {
	return healthcheck.Up
}
