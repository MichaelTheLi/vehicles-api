package usecase

import (
	"context"
	"vehicles_api/domain"
)

type vehicleUseCase struct {
	vehicleRepo domain.VehicleRepository
}

func NewVehicleUseCase(vehicleRepo domain.VehicleRepository) domain.VehicleUseCase {
	return &vehicleUseCase{
		vehicleRepo: vehicleRepo,
	}
}

func (uc *vehicleUseCase) GetVehicles(ctx context.Context, first int, after *string) (*domain.VehicleConnection, error) {
	return uc.vehicleRepo.FindAll(first, after)
} 