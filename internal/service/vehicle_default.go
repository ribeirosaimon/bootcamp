package service

import (
	"github.com/ribeirosaimon/bootcamp/internal"
	"github.com/ribeirosaimon/bootcamp/web/apierr"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindByID find one vehicle
func (s *VehicleDefault) FindByID(id int) (internal.Vehicle, error) {
	return s.rp.FindByID(id)
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) Save(v *internal.Vehicle) error {
	return s.rp.Save(v)
}

func (s *VehicleDefault) SearchVehiclesByColorAndYear(filter internal.VehicleFilter) ([]internal.Vehicle, error) {
	return s.rp.SearchVehiclesByColorAndYear(filter)
}

func (s *VehicleDefault) SearchVehicleByBrandAndYear(d internal.VehicleFilter) ([]internal.Vehicle, error) {
	return s.rp.SearchVehicleByBrandAndYear(d)
}

func (s *VehicleDefault) GetAverageSpeed(brand string) (float64, error) {
	brandVehicles, err := s.rp.SearchVehicleByBrandAndYear(
		internal.VehicleFilter{
			Brand: brand,
		})
	if err != nil {
		return 0, err
	}
	total := 0.0
	for _, vehicle := range brandVehicles {
		total += vehicle.MaxSpeed
	}
	return total / float64(len(brandVehicles)), nil
}

func (s *VehicleDefault) SaveMany(vehicles []internal.Vehicle) error {
	// não é o ideal, porém como é um file,
	// na camada de persistência faria a mesma coisa
	for index := range vehicles {
		if err := s.rp.Save(&vehicles[index]); err != nil {
			return err
		}
	}
	return nil
}

func (s *VehicleDefault) UpdateSpeed(id int, speed float64) error {
	vehicle, err := s.rp.FindByID(id)
	if err != nil {
		return err
	}
	if speed > 200.0 {
		return apierr.NewBadRequestApiErr(
			apierr.WithMessage("speed must be greater than 200.0"),
		)
	}
	vehicle.MaxSpeed = speed

	if err = s.rp.Update(&vehicle); err != nil {
		return err
	}
	return nil
}

func (s *VehicleDefault) GetVehicleByFuelType(fuelType string) ([]internal.Vehicle, error) {
	byFuelType, err := s.rp.SearchVehicleByFuelType(fuelType)
	if err != nil {
		return nil, err
	}
	return byFuelType, nil
}

func (s *VehicleDefault) DeleteVehicle(id int) error {
	return s.rp.DeleteVehicle(id)
}

func (s *VehicleDefault) SearchVehiclesByTransmissionType(transmissionType string) ([]internal.Vehicle, error) {
	return s.rp.SearchVehiclesByTransmissionType(transmissionType)
}

func (s *VehicleDefault) UpdateFuel(id int, fuel string) error {
	vehicle, err := s.rp.FindByID(id)
	if err != nil {
		return err
	}
	vehicle.FuelType = fuel
	if err = s.rp.Update(&vehicle); err != nil {
		return err
	}
	return nil
}

func (s *VehicleDefault) FindAverageCapacity(brand string) (float64, error) {
	carByBrand, err := s.SearchVehicleByBrandAndYear(internal.VehicleFilter{
		Brand: brand,
	})
	if err != nil {
		return 0, err
	}

	var res int
	for _, car := range carByBrand {
		res += car.Capacity
	}
	return float64(res / len(carByBrand)), nil
}

func (s *VehicleDefault) SearchVehicleByDimensions(minLen, maxLen, minWidth, maxWidth float64) ([]internal.Vehicle, error) {
	return s.rp.SearchVehicleByDimensions(minLen, maxLen, minWidth, maxWidth)
}

func (s *VehicleDefault) SearchVehiclesByWeight(v float64, v2 float64) ([]internal.Vehicle, error) {
	return s.rp.SearchVehiclesByWeight(v, v2)
}
