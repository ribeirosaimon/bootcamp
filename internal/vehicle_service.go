package internal

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	FindByID(int) (Vehicle, error)
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	Save(v *Vehicle) error
	SearchVehiclesByColorAndYear(VehicleFilter) ([]Vehicle, error)
	SearchVehicleByBrandAndYear(VehicleFilter) ([]Vehicle, error)
	GetAverageSpeed(string) (float64, error)
	SaveMany(vehicles []Vehicle) error
	UpdateSpeed(int, float64) error
	GetVehicleByFuelType(fuelType string) ([]Vehicle, error)
	DeleteVehicle(id int) error
	SearchVehiclesByTransmissionType(transmissionType string) ([]Vehicle, error)
	UpdateFuel(id int, fuel string) error
	FindAverageCapacity(brand string) (float64, error)
	SearchVehicleByDimensions(minLen, maxLen, minWidth, maxWidth float64) ([]Vehicle, error)
	SearchVehiclesByWeight(v float64, v2 float64) ([]Vehicle, error)
}
