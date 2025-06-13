package internal

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	FindByID(int) (Vehicle, error)
	Save(v *Vehicle) error
	Update(v *Vehicle) error
	SearchVehiclesByColorAndYear(d VehicleFilter) ([]Vehicle, error)
	SearchVehicleByBrandAndYear(d VehicleFilter) ([]Vehicle, error)
	SearchVehicleByFuelType(d string) ([]Vehicle, error)
	DeleteVehicle(id int) error
	SearchVehiclesByTransmissionType(transmissionType string) ([]Vehicle, error)
	SearchVehicleByDimensions(minLen, maxLen, minWidth, maxWidth float64) ([]Vehicle, error)
	SearchVehiclesByWeight(v float64, v2 float64) ([]Vehicle, error)
}
