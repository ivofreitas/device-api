package device

import (
	"context"
	"database/sql"
	"errors"
	mocks "github.com/ivofreitas/device-api/internal/api/device/mock"
	"github.com/ivofreitas/device-api/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

type testCase struct {
	name        string
	input       interface{}
	expected    interface{}
	expectedErr *domain.Error
	mockSetup   func(*mocks.Repository, context.Context)
}

func TestServiceMethods(t *testing.T) {
	testCases := []testCase{
		{
			name:     "Create Device - Success",
			input:    &domain.Device{Name: "Test Device", Brand: "Test Brand"},
			expected: &domain.Device{Id: 1, Name: "Test Device", Brand: "Test Brand"},
			mockSetup: func(m *mocks.Repository, ctx context.Context) {
				m.On("Create", ctx, mock.Anything).Return(&domain.Device{Id: 1, Name: "Test Device", Brand: "Test Brand"}, nil)
			},
		},
		{
			name:        "Create Device - Failure",
			input:       &domain.Device{Name: "Test Device", Brand: "Test Brand"},
			expectedErr: &domain.Error{Type: "create_error", Status: http.StatusInternalServerError},
			mockSetup: func(m *mocks.Repository, ctx context.Context) {
				m.On("Create", ctx, mock.Anything).Return(nil, errors.New("DB error"))
			},
		},
		{
			name:     "Update Device - Success",
			input:    &domain.Update{Id: 1, Device: &domain.Device{Name: "Updated Name", Brand: "Same Brand"}},
			expected: &domain.Device{Id: 1, Name: "Updated Name", Brand: "Same Brand"},
			mockSetup: func(m *mocks.Repository, ctx context.Context) {
				m.On("GetById", ctx, 1).Return(&domain.Device{Id: 1, Name: "Old Name", Brand: "Same Brand"}, nil)
				m.On("Update", ctx, mock.Anything).Return(nil)
			},
		},
		{
			name:        "Update Device - Not Found",
			input:       &domain.Update{Id: 1, Device: &domain.Device{Name: "Updated Name"}},
			expectedErr: &domain.Error{Type: "not_found", Status: http.StatusNotFound},
			mockSetup: func(m *mocks.Repository, ctx context.Context) {
				m.On("GetById", ctx, 1).Return((*domain.Device)(nil), sql.ErrNoRows)
			},
		},
		{
			name:        "Update Device - In Use State, Forbidden Fields",
			input:       &domain.Update{Id: 1, Device: &domain.Device{Name: "New Name", Brand: "New Brand"}},
			expectedErr: &domain.Error{Type: "update_error", Status: http.StatusForbidden},
			mockSetup: func(m *mocks.Repository, ctx context.Context) {
				m.On("GetById", ctx, 1).Return(&domain.Device{Id: 1, State: domain.InUseState, Name: "Old Name", Brand: "Old Brand"}, nil)
			},
		},
		{
			name:  "GetAll - Success",
			input: nil,
			expected: []domain.Device{
				{Id: 1, Name: "Device1"},
				{Id: 2, Name: "Device2"},
			},
			mockSetup: func(m *mocks.Repository, ctx context.Context) {
				m.On("GetAll", ctx).Return([]domain.Device{
					{Id: 1, Name: "Device1"},
					{Id: 2, Name: "Device2"},
				}, nil)
			},
		},
		{
			name:        "GetAll - Failure",
			input:       nil,
			expectedErr: &domain.Error{Type: "fetch_error", Status: http.StatusInternalServerError},
			mockSetup: func(m *mocks.Repository, ctx context.Context) {
				m.On("GetAll", ctx).Return(nil, errors.New("DB error"))
			},
		},
		{
			name:  "GetByBrand - Success",
			input: &domain.GetByBrand{Brand: "Test Brand"},
			expected: []domain.Device{
				{Id: 1, Name: "Device1", Brand: "Test Brand"},
			},
			mockSetup: func(m *mocks.Repository, ctx context.Context) {
				m.On("GetByBrand", ctx, "Test Brand").Return([]domain.Device{
					{Id: 1, Name: "Device1", Brand: "Test Brand"},
				}, nil)
			},
		},
		{
			name:        "GetByBrand - Failure",
			input:       &domain.GetByBrand{Brand: "Unknown Brand"},
			expectedErr: &domain.Error{Type: "fetch_error", Status: http.StatusInternalServerError},
			mockSetup: func(m *mocks.Repository, ctx context.Context) {
				m.On("GetByBrand", ctx, "Unknown Brand").Return(nil, errors.New("DB error"))
			},
		},
		{
			name:  "GetByState - Success",
			input: &domain.GetByState{State: domain.AvailableState},
			expected: []domain.Device{
				{Id: 1, Name: "Device1", State: domain.AvailableState},
			},
			mockSetup: func(m *mocks.Repository, ctx context.Context) {
				m.On("GetByState", ctx, domain.AvailableState).Return(
					[]domain.Device{{Id: 1, Name: "Device1", State: domain.AvailableState}}, nil)
			},
		},
		{
			name:        "GetByState - Failure",
			input:       &domain.GetByState{State: domain.State(-1)},
			expectedErr: &domain.Error{Type: "fetch_error", Status: http.StatusInternalServerError},
			mockSetup: func(m *mocks.Repository, ctx context.Context) {
				m.On("GetByState", ctx, domain.State(-1)).
					Return(nil, errors.New("DB error"))
			},
		},
		{
			name:     "GetById - Device Exists",
			input:    &domain.GetById{Id: 1},
			expected: &domain.Device{Id: 1, Name: "Device1"},
			mockSetup: func(m *mocks.Repository, ctx context.Context) {
				m.On("GetById", ctx, 1).Return(&domain.Device{Id: 1, Name: "Device1"}, nil)
			},
		},
		{
			name:        "GetById - Not Found",
			input:       &domain.GetById{Id: 1000},
			expectedErr: &domain.Error{Type: "not_found", Status: http.StatusNotFound},
			mockSetup: func(m *mocks.Repository, ctx context.Context) {
				m.On("GetById", ctx, 1000).Return(nil, sql.ErrNoRows)
			},
		},
		{
			name:        "GetById - Failure",
			input:       &domain.GetById{Id: 999},
			expectedErr: &domain.Error{Type: "get_by_id_error", Status: http.StatusInternalServerError},
			mockSetup: func(m *mocks.Repository, ctx context.Context) {
				m.On("GetById", ctx, 999).Return(nil, errors.New("DB error"))
			},
		},
		{
			name:     "Delete Device - Success",
			input:    &domain.Delete{Id: 1},
			expected: nil,
			mockSetup: func(m *mocks.Repository, ctx context.Context) {
				m.On("GetById", ctx, 1).Return(&domain.Device{Id: 1, State: domain.AvailableState}, nil)
				m.On("Delete", ctx, 1).Return(nil)
			},
		},
		{
			name:        "Delete Device - Not Found",
			input:       &domain.Delete{Id: 999},
			expectedErr: &domain.Error{Type: "not_found", Status: http.StatusNotFound},
			mockSetup: func(m *mocks.Repository, ctx context.Context) {
				m.On("GetById", ctx, 999).Return(nil, sql.ErrNoRows)
			},
		},
		{
			name:        "Delete Device - In Use",
			input:       &domain.Delete{Id: 2},
			expectedErr: &domain.Error{Type: "delete_error", Status: http.StatusForbidden},
			mockSetup: func(m *mocks.Repository, ctx context.Context) {
				m.On("GetById", ctx, 2).Return(&domain.Device{Id: 2, State: domain.InUseState}, nil)
			},
		},
		{
			name:        "Delete Device - Failure",
			input:       &domain.Delete{Id: 3},
			expectedErr: &domain.Error{Type: "delete_error", Status: http.StatusInternalServerError},
			mockSetup: func(m *mocks.Repository, ctx context.Context) {
				m.On("GetById", ctx, 3).Return(&domain.Device{Id: 3, State: domain.AvailableState}, nil)
				m.On("Delete", ctx, 3).Return(errors.New("DB error"))
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(mocks.Repository)
			service := NewService(mockRepo)
			ctx := context.Background()

			if tc.mockSetup != nil {
				tc.mockSetup(mockRepo, ctx)
			}

			var result interface{}
			var err error
			switch v := tc.input.(type) {
			case *domain.Device:
				result, err = service.Create(ctx, v)
			case *domain.Update:
				result, err = service.Update(ctx, v)
			case *domain.GetById:
				result, err = service.GetById(ctx, v)
			case *domain.GetByState:
				result, err = service.GetByState(ctx, v)
			case *domain.GetByBrand:
				result, err = service.GetByBrand(ctx, v)
			case nil:
				result, err = service.GetAll(ctx, v)
			case *domain.Delete:
				result, err = service.Delete(ctx, v)
			}

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr.Type, err.(*domain.Error).Type)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
