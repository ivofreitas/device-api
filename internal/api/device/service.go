package device

import (
	context "context"
	"github.com/ivofreitas/device-api/internal/repository/device"
)

type Service struct {
	repository device.Repository
}

func NewService(repository device.Repository) *Service {
	return &Service{repository}
}

func (s *Service) Create(ctx context.Context, param interface{}) (interface{}, error) {
	return nil, nil
}

func (s *Service) Update(ctx context.Context, param interface{}) (interface{}, error) {
	return nil, nil
}

func (s *Service) GetAll(ctx context.Context, param interface{}) (interface{}, error) {
	return nil, nil
}

func (s *Service) GetById(ctx context.Context, param interface{}) (interface{}, error) {
	return nil, nil
}

func (s *Service) GetByBrand(ctx context.Context, param interface{}) (interface{}, error) {
	return nil, nil
}

func (s *Service) GetByState(ctx context.Context, param interface{}) (interface{}, error) {
	return nil, nil
}

func (s *Service) Delete(ctx context.Context, param interface{}) (interface{}, error) {
	return nil, nil
}
