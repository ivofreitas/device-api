package device

import (
	"context"
	"database/sql"
	"github.com/ivofreitas/device-api/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, device *domain.Device) (int, error)
	Update(ctx context.Context, device *domain.Device) error
	GetAll(ctx context.Context) ([]domain.Device, error)
	GetById(ctx context.Context, id int) (*domain.Device, error)
	GetByBrand(ctx context.Context, brand string) ([]domain.Device, error)
	GetByState(ctx context.Context, state *domain.State) ([]domain.Device, error)
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, device *domain.Device) (int, error) {
	query := `INSERT INTO devices_schema.devices (name, brand, state, creation_time) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := r.db.QueryRowContext(ctx, query, device.Name, device.Brand, device.State, device.CreationTime).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) Update(ctx context.Context, device *domain.Device) error {
	query := `UPDATE devices_schema.devices SET name = $1, brand = $2, state = $3, creation_time = $4 WHERE id = $5`
	_, err := r.db.ExecContext(ctx, query, device.Name, device.Brand, device.State, device.CreationTime, device.Id)
	return err
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Device, error) {
	query := `SELECT id, name, brand, state, creation_time FROM devices_schema.devices`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []domain.Device
	for rows.Next() {
		var device domain.Device
		if err := rows.Scan(&device.Id, &device.Name, &device.Brand, &device.State, &device.CreationTime); err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}

func (r *repository) GetById(ctx context.Context, id int) (*domain.Device, error) {
	query := `SELECT id, name, brand, state, creation_time FROM devices_schema.devices WHERE id = $1`
	var device domain.Device
	err := r.db.QueryRowContext(ctx, query, id).Scan(&device.Id, &device.Name, &device.Brand, &device.State, &device.CreationTime)
	return &device, err
}

func (r *repository) GetByBrand(ctx context.Context, brand string) ([]domain.Device, error) {
	query := `SELECT id, name, brand, state, creation_time FROM devices_schema.devices WHERE brand = $1`
	rows, err := r.db.QueryContext(ctx, query, brand)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []domain.Device
	for rows.Next() {
		var device domain.Device
		if err = rows.Scan(&device.Id, &device.Name, &device.Brand, &device.State, &device.CreationTime); err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}

func (r *repository) GetByState(ctx context.Context, state *domain.State) ([]domain.Device, error) {
	query := `SELECT id, name, brand, state, creation_time FROM devices_schema.devices WHERE state = $1`
	rows, err := r.db.QueryContext(ctx, query, state)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []domain.Device
	for rows.Next() {
		var device domain.Device
		if err = rows.Scan(&device.Id, &device.Name, &device.Brand, &device.State, &device.CreationTime); err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM devices_schema.devices WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
