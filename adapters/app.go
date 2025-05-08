package adapters

import (
	"fmt"
	mysqldriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"vehicles_api/adapters/mysql"
	"vehicles_api/domain"
)

func BuildApp(config *Config) *domain.App {
	db := buildDB(config)
	vehicleRepo := mysql.NewVehicleRepository(db)

	return &domain.App{
		VehicleRepo: vehicleRepo,
	}
}

func buildDB(config *Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open(mysqldriver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	return db
}
