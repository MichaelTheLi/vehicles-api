package domain

import "time"

type Vehicle struct {
    ID            uint      `json:"id" gorm:"column:ID"`
    BaseVehicleID uint      `json:"baseVehicleId" gorm:"column:BaseVehicleID"`
    Make          string    `json:"make" gorm:"column:Make"`
    Model         string    `json:"model" gorm:"column:Model"`
    Year          int       `json:"year" gorm:"column:Year"`
    VIN           string    `json:"vin" gorm:"column:VIN"`
    CreatedAt     time.Time `json:"createdAt" gorm:"column:CreatedAt"`
    UpdatedAt     time.Time `json:"updatedAt" gorm:"column:UpdatedAt"`
}

type PageInfo struct {
    HasNextPage bool
    EndCursor   string
}

type VehicleConnection struct {
    Edges    []VehicleEdge
    PageInfo PageInfo
}

type VehicleEdge struct {
    Node   Vehicle
    Cursor string
}

type VehicleRepository interface {
    FindAll(first int, after *string) (*VehicleConnection, error)
} 