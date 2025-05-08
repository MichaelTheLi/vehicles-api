package resolvers

import "vehicles_api/domain"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	vehicleRepo domain.VehicleRepository
}

func NewResolver(vehicleRepo domain.VehicleRepository) *Resolver {
	return &Resolver{
		vehicleRepo: vehicleRepo,
	}
}
