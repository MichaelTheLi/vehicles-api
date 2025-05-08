package mysql

import (
	"gorm.io/gorm"
	"vehicles_api/domain"
)

type vehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) domain.VehicleRepository {
	return &vehicleRepository{db: db}
}

func (r *vehicleRepository) FindAll() ([]domain.Vehicle, error) {
	var vehicles []domain.Vehicle
	if err := r.db.Table("Vehicle").
		Select("Vehicle.VehicleID as ID, Vehicle.BaseVehicleID, Make.MakeName as Make, Model.ModelName as Model, Year.YearID as Year").
		Joins("INNER JOIN BaseVehicle ON Vehicle.BaseVehicleID = BaseVehicle.BaseVehicleID").
		Joins("INNER JOIN Make ON BaseVehicle.MakeID = Make.MakeID").
		Joins("INNER JOIN Model ON BaseVehicle.ModelID = Model.ModelID").
		Joins("INNER JOIN Year ON BaseVehicle.YearID = Year.YearID").
		Scan(&vehicles).Error; err != nil {
		return nil, err
	}
	return vehicles, nil
} 