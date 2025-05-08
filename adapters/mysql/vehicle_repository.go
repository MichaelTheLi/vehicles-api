package mysql

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"gorm.io/gorm"
	"vehicles_api/domain"
)

type vehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) domain.VehicleRepository {
	return &vehicleRepository{db: db}
}

func (r *vehicleRepository) FindAll(first int, after *string) (*domain.VehicleConnection, error) {
	if first <= 0 {
		first = 10 // default page size
	}

	query := r.db.Table("Vehicle").
		Select("Vehicle.VehicleID as ID, Vehicle.BaseVehicleID, Make.MakeName as Make, Model.ModelName as Model, Year.YearID as Year").
		Joins("INNER JOIN BaseVehicle ON Vehicle.BaseVehicleID = BaseVehicle.BaseVehicleID").
		Joins("INNER JOIN Make ON BaseVehicle.MakeID = Make.MakeID").
		Joins("INNER JOIN Model ON BaseVehicle.ModelID = Model.ModelID").
		Joins("INNER JOIN Year ON BaseVehicle.YearID = Year.YearID")

	if after != nil {
		cursor, err := decodeCursor(*after)
		if err != nil {
			return nil, err
		}
		query = query.Where("Vehicle.VehicleID > ?", cursor)
	}

	// Get one extra record to determine if there are more pages
	query = query.Order("Vehicle.VehicleID ASC").Limit(first + 1)

	var vehicles []domain.Vehicle
	if err := query.Scan(&vehicles).Error; err != nil {
		return nil, err
	}

	hasNextPage := len(vehicles) > first
	if hasNextPage {
		vehicles = vehicles[:first]
	}

	edges := make([]domain.VehicleEdge, len(vehicles))
	for i, v := range vehicles {
		edges[i] = domain.VehicleEdge{
			Node:   v,
			Cursor: encodeCursor(v.ID),
		}
	}

	var endCursor string
	if len(edges) > 0 {
		endCursor = encodeCursor(edges[len(edges)-1].Node.ID)
	}

	return &domain.VehicleConnection{
		Edges: edges,
		PageInfo: domain.PageInfo{
			HasNextPage: hasNextPage,
			EndCursor:   endCursor,
		},
	}, nil
}

func encodeCursor(id uint) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", id)))
}

func decodeCursor(cursor string) (uint, error) {
	bytes, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return 0, err
	}
	id, err := strconv.ParseUint(string(bytes), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
} 