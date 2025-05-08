package resolver

import (
	"context"
	"vehicles_api/domain"
)

type VehicleResolver struct {
	vehicleUseCase domain.VehicleUseCase
}

func NewVehicleResolver(vehicleUseCase domain.VehicleUseCase) *VehicleResolver {
	return &VehicleResolver{
		vehicleUseCase: vehicleUseCase,
	}
}

func (r *VehicleResolver) Vehicles(ctx context.Context, first *int, after *string) (*domain.VehicleConnection, error) {
	if first == nil {
		defaultFirst := 10
		first = &defaultFirst
	}
	return r.vehicleUseCase.GetVehicles(ctx, *first, after)
} 