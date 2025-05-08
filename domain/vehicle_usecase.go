package domain

import "context"

type VehicleUseCase interface {
	GetVehicles(ctx context.Context, first int, after *string) (*VehicleConnection, error)
} 