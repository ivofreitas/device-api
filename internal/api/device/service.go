package device

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ivofreitas/device-api/internal/domain"
	"net/http"
	"time"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository}
}

// Create
// @Summary Create a new device
// @Description Adds a new device to the inventory
// @Tags Device
// @Accept  json
// @Produce  json
// @Param request body domain.Device true "Device details"
// @Success 201 {object} domain.Device "Created device"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/devices [post]
func (s *Service) Create(ctx context.Context, param interface{}) (interface{}, error) {
	device := param.(*domain.Device)
	createdDevice, err := s.repository.Create(ctx, device)
	if err != nil {
		return nil, &domain.Error{Type: "create_error", Status: http.StatusInternalServerError, Detail: err.Error()}
	}

	return createdDevice, nil
}

// Update
// @Summary Update an existing device
// @Description Updates device details if allowed. `PUT` requires a full update, while `PATCH` allows partial updates.
// @Tags Device
// @Accept  json
// @Produce  json
// @Param id path int true "Device ID"
// @Param request body domain.Update true "Device update details"
// @Success 200 {object} domain.Device "Updated device"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 403 {object} map[string]string "Forbidden update"
// @Failure 404 {object} map[string]string "Device not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/devices/{id} [put]
// @Router /v1/devices/{id} [patch]
func (s *Service) Update(ctx context.Context, param interface{}) (interface{}, error) {
	update := param.(*domain.Update)

	if update.CreationTime != (time.Time{}) {
		return nil, &domain.Error{
			Type:   "update_error",
			Status: http.StatusForbidden,
			Detail: "cannot update creation time of a device"}
	}

	existingDevice, err := s.repository.GetById(ctx, update.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &domain.Error{Type: "not_found", Status: http.StatusNotFound, Detail: "device not found"}
		}
		return nil, &domain.Error{Type: "update_error", Status: http.StatusInternalServerError, Detail: err.Error()}
	}

	if existingDevice.State == domain.InUseState &&
		(*update.Name != existingDevice.Name || *update.Brand != existingDevice.Brand) {
		return nil, &domain.Error{
			Type:   "update_error",
			Status: http.StatusForbidden,
			Detail: "cannot update name or brand of a device in use"}
	}

	existingDevice.Name = *update.Name
	existingDevice.Brand = *update.Brand
	existingDevice.State = *update.State

	if err = s.repository.Update(ctx, existingDevice); err != nil {
		return nil, &domain.Error{Type: "update_error", Status: http.StatusInternalServerError, Detail: err.Error()}
	}

	return existingDevice, nil
}

// Patch (PATCH)
// @Summary Partially update an existing device
// @Description Allows partial updates to a device. Only provided fields are modified.
// @Tags Device
// @Accept  json
// @Produce  json
// @Param id path int true "Device ID"
// @Param request body domain.Patch true "Partial device update details"
// @Success 200 {object} domain.Device "Updated device"
// @Failure 400 {object} map[string]string "Invalid request body"
// @Failure 403 {object} map[string]string "Forbidden update"
// @Failure 404 {object} map[string]string "Device not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/devices/{id} [patch]
func (s *Service) Patch(ctx context.Context, param interface{}) (interface{}, error) {
	patch := param.(*domain.Patch)

	if patch.CreationTime != (time.Time{}) {
		return nil, &domain.Error{
			Type:   "update_error",
			Status: http.StatusForbidden,
			Detail: "cannot update creation time of a device"}
	}

	existingDevice, err := s.repository.GetById(ctx, patch.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &domain.Error{Type: "not_found", Status: http.StatusNotFound, Detail: "device not found"}
		}
		return nil, &domain.Error{Type: "patch_error", Status: http.StatusInternalServerError, Detail: err.Error()}
	}

	if existingDevice.State == domain.InUseState &&
		((patch.Name != nil && *patch.Name != existingDevice.Name) ||
			(patch.Brand != nil && *patch.Brand != existingDevice.Brand)) {
		return nil, &domain.Error{
			Type:   "patch_error",
			Status: http.StatusForbidden,
			Detail: "cannot update name or brand of a device in use"}

	}

	if patch.Name != nil {
		existingDevice.Name = *patch.Name
	}
	if patch.Brand != nil {
		existingDevice.Brand = *patch.Brand
	}
	if patch.State != nil {
		existingDevice.State = *patch.State
	}

	if err = s.repository.Update(ctx, existingDevice); err != nil {
		return nil, &domain.Error{Type: "update_error", Status: http.StatusInternalServerError, Detail: err.Error()}
	}

	return existingDevice, nil
}

// GetAll
// @Summary Get all devices
// @Description Retrieves a list of all devices
// @Tags Device
// @Produce json
// @Success 200 {array} domain.Device "List of devices"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/devices [get]
func (s *Service) GetAll(ctx context.Context, param interface{}) (interface{}, error) {
	devices, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, &domain.Error{Type: "fetch_error", Status: http.StatusInternalServerError, Detail: err.Error()}
	}
	return devices, nil
}

// GetById
// @Summary Get a device by ID
// @Description Retrieves a single device by its ID
// @Tags Device
// @Produce json
// @Param id path int true "Device ID"
// @Success 200 {object} domain.Device "Device details"
// @Failure 404 {object} map[string]string "Device not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/devices/{id} [get]
func (s *Service) GetById(ctx context.Context, param interface{}) (interface{}, error) {
	idParam := param.(*domain.GetById)
	device, err := s.repository.GetById(ctx, idParam.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &domain.Error{Type: "not_found", Status: http.StatusNotFound, Detail: "device not found"}
		}
		return nil, &domain.Error{Type: "get_by_id_error", Status: http.StatusInternalServerError, Detail: err.Error()}
	}
	return device, nil
}

// GetByBrand
// @Summary Get a device by brand
// @Description Retrieves a single device by its brand
// @Tags Device
// @Produce json
// @Param brand path int true "Device Brand"
// @Success 200 {object} domain.Device "Device details"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/devices/brand/{brand} [get]
func (s *Service) GetByBrand(ctx context.Context, param interface{}) (interface{}, error) {
	brandParam := param.(*domain.GetByBrand)
	devices, err := s.repository.GetByBrand(ctx, brandParam.Brand)
	if err != nil {
		return nil, &domain.Error{Type: "fetch_error", Status: http.StatusInternalServerError, Detail: err.Error()}
	}
	return devices, nil
}

// GetByState
// @Summary Get a device by state
// @Description Retrieves a single device by its state
// @Tags Device
// @Produce json
// @Param brand path int true "Device State"
// @Success 200 {object} domain.Device "Device details"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/devices/state/{state} [get]
func (s *Service) GetByState(ctx context.Context, param interface{}) (interface{}, error) {
	stateParam := param.(*domain.GetByState)
	devices, err := s.repository.GetByState(ctx, stateParam.State)
	if err != nil {
		return nil, &domain.Error{Type: "fetch_error", Status: http.StatusInternalServerError, Detail: err.Error()}
	}
	return devices, nil
}

// Delete
// @Summary Delete a device
// @Description Removes a device from the inventory
// @Tags Device
// @Param id path int true "Device ID"
// @Success 204 "No content"
// @Failure 403 {object} map[string]string "Cannot delete device in use"
// @Failure 404 {object} map[string]string "Device not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/devices/{id} [delete]
func (s *Service) Delete(ctx context.Context, param interface{}) (interface{}, error) {
	deleteParam := param.(*domain.Delete)
	existingDevice, err := s.repository.GetById(ctx, deleteParam.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &domain.Error{Type: "not_found", Status: http.StatusNotFound, Detail: "device not found"}
		}
		return nil, &domain.Error{Type: "delete_error", Status: http.StatusInternalServerError, Detail: err.Error()}
	}

	if existingDevice.State == domain.InUseState {
		return nil, &domain.Error{
			Type:   "delete_error",
			Status: http.StatusForbidden,
			Detail: "cannot delete a device that is in use"}
	}

	if err = s.repository.Delete(ctx, deleteParam.Id); err != nil {
		return nil, &domain.Error{Type: "delete_error", Status: http.StatusInternalServerError, Detail: err.Error()}
	}
	return nil, nil
}
