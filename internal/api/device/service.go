package device

import (
	"context"
	"github.com/ivofreitas/device-api/internal/domain"
	"net/http"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository}
}

func (s *Service) Create(ctx context.Context, param interface{}) (interface{}, error) {
	device := param.(domain.Device)
	id, err := s.repository.Create(ctx, device)
	if err != nil {
		return nil, domain.Error{Type: "create_error", Status: http.StatusInternalServerError, Detail: err.Error()}
	}
	return id, nil
}

func (s *Service) Update(ctx context.Context, param interface{}) (interface{}, error) {
	update := param.(domain.Update)
	existingDevice, err := s.repository.GetById(ctx, update.Id)
	if err != nil || existingDevice == nil {
		return nil, domain.Error{Type: "not_found", Status: http.StatusNotFound, Detail: "device not found"}
	}

	if existingDevice.State == domain.InUseState &&
		(update.Name != existingDevice.Name || update.Brand != existingDevice.Brand) {
		return nil, domain.Error{
			Type:   "update_error",
			Status: http.StatusForbidden,
			Detail: "cannot update name or brand of a device in use"}
	}

	if err := s.repository.Update(ctx, update.Device); err != nil {
		return nil, domain.Error{Type: "update_error", Status: http.StatusInternalServerError, Detail: err.Error()}
	}
	return nil, nil
}

func (s *Service) Patch(ctx context.Context, param interface{}) (interface{}, error) {
	patch := param.(domain.Patch)
	existingDevice, err := s.repository.GetById(ctx, patch.Id)
	if err != nil || existingDevice == nil {
		return nil, domain.Error{Type: "not_found", Status: http.StatusNotFound, Detail: "device not found"}
	}

	if existingDevice.State == domain.InUseState &&
		(patch.Name != "" && patch.Name != existingDevice.Name ||
			patch.Brand != "" && patch.Brand != existingDevice.Brand) {
		return nil, domain.Error{
			Type:   "update_error",
			Status: http.StatusForbidden,
			Detail: "cannot update name or brand of a device in use"}
	}

	updatedDevice := *existingDevice
	if patch.Name != "" {
		updatedDevice.Name = patch.Name
	}
	if patch.Brand != "" {
		updatedDevice.Brand = patch.Brand
	}
	if patch.State != 0 {
		updatedDevice.State = patch.State
	}

	if err := s.repository.Update(ctx, updatedDevice); err != nil {
		return nil, domain.Error{Type: "update_error", Status: http.StatusInternalServerError, Detail: err.Error()}
	}
	return nil, nil
}

func (s *Service) GetAll(ctx context.Context, param interface{}) (interface{}, error) {
	devices, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, domain.Error{Type: "fetch_error", Status: http.StatusInternalServerError, Detail: err.Error()}
	}
	return devices, nil
}

func (s *Service) GetById(ctx context.Context, param interface{}) (interface{}, error) {
	idParam := param.(domain.GetById)
	device, err := s.repository.GetById(ctx, idParam.Id)
	if err != nil || device == nil {
		return nil, domain.Error{Type: "not_found", Status: http.StatusNotFound, Detail: "device not found"}
	}
	return device, nil
}

func (s *Service) GetByBrand(ctx context.Context, param interface{}) (interface{}, error) {
	brandParam := param.(domain.GetByBrand)
	devices, err := s.repository.GetByBrand(ctx, brandParam.Brand)
	if err != nil {
		return nil, domain.Error{Type: "fetch_error", Status: http.StatusInternalServerError, Detail: err.Error()}
	}
	return devices, nil
}

func (s *Service) GetByState(ctx context.Context, param interface{}) (interface{}, error) {
	stateParam := param.(domain.GetByState)
	devices, err := s.repository.GetByState(ctx, stateParam.State)
	if err != nil {
		return nil, domain.Error{Type: "fetch_error", Status: http.StatusInternalServerError, Detail: err.Error()}
	}
	return devices, nil
}

func (s *Service) Delete(ctx context.Context, param interface{}) (interface{}, error) {
	deleteParam := param.(domain.Delete)
	existingDevice, err := s.repository.GetById(ctx, deleteParam.Id)
	if err != nil || existingDevice == nil {
		return nil, domain.Error{Type: "not_found", Status: http.StatusNotFound, Detail: "device not found"}
	}

	if existingDevice.State == domain.InUseState {
		return nil, domain.Error{
			Type:   "delete_error",
			Status: http.StatusForbidden,
			Detail: "cannot delete a device that is in use"}
	}

	if err := s.repository.Delete(ctx, deleteParam.Id); err != nil {
		return nil, domain.Error{Type: "delete_error", Status: http.StatusInternalServerError, Detail: err.Error()}
	}
	return nil, nil
}
