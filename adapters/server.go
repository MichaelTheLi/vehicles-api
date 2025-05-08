package adapters

import (
	"vehicles_api/api"
	"vehicles_api/api/graph/resolvers"
	"vehicles_api/domain"
)

type ServerState struct {
	Resolver *resolvers.Resolver
}

func BuildServerState(app *domain.App) *api.ServerState {
	return &api.ServerState{
		VehicleRepo: app.VehicleRepo,
	}
}
