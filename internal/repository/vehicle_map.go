package repository

import (
	"github.com/ribeirosaimon/bootcamp/internal"
	"github.com/ribeirosaimon/bootcamp/web/apierr"
	"strings"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb, lastIndex: len(defaultDb)}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db        map[int]internal.Vehicle
	lastIndex int
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

func (r *VehicleMap) Save(v *internal.Vehicle) error {
	if v.Id == 0 {
		r.lastIndex++
		v.Id = r.lastIndex
	}

	if _, ok := r.db[v.Id]; !ok {
		r.db[v.Id] = *v
		return nil
	}

	return apierr.NewConflictApiErr()
}

func (r *VehicleMap) SearchVehiclesByColorAndYear(filter internal.VehicleFilter) ([]internal.Vehicle, error) {
	resp := make([]internal.Vehicle, 0)
	for _, value := range r.db {
		if filter.Color != "" && !strings.EqualFold(value.Color, filter.Color) {
			continue
		}

		if filter.Year != 0 && filter.Year != value.FabricationYear {
			continue
		}

		resp = append(resp, value)
	}
	return resp, nil
}

func (r *VehicleMap) SearchVehicleByBrandAndYear(filter internal.VehicleFilter) ([]internal.Vehicle, error) {
	resp := make([]internal.Vehicle, 0)
	for _, value := range r.db {
		if filter.Brand != "" && !strings.EqualFold(value.Brand, filter.Brand) {
			continue
		}

		if (filter.PairDates.Start != 0 && filter.PairDates.End != 0) &&
			(filter.PairDates.Start >= value.FabricationYear || filter.PairDates.End <= value.FabricationYear) {
			continue
		}
		resp = append(resp, value)
	}
	return resp, nil
}

func (r *VehicleMap) FindByID(id int) (internal.Vehicle, error) {
	if res, ok := r.db[id]; ok {
		return res, nil
	}
	return internal.Vehicle{}, apierr.NewNotFoundApiErr()
}

func (r *VehicleMap) Update(v *internal.Vehicle) error {
	if _, ok := r.db[v.Id]; !ok {
		return apierr.NewNotFoundApiErr()
	}
	r.db[v.Id] = *v
	return nil
}

func (r *VehicleMap) SearchVehicleByFuelType(fuelType string) ([]internal.Vehicle, error) {
	resp := make([]internal.Vehicle, 0)
	for _, value := range r.db {
		if strings.EqualFold(value.FuelType, fuelType) {
			resp = append(resp, value)
		}
	}
	return resp, nil
}

func (r *VehicleMap) DeleteVehicle(id int) error {
	if _, ok := r.db[id]; !ok {
		return apierr.NewNotFoundApiErr()
	}
	delete(r.db, id)
	return nil
}

func (r *VehicleMap) SearchVehiclesByTransmissionType(transmissionType string) ([]internal.Vehicle, error) {
	resp := make([]internal.Vehicle, 0)
	for _, value := range r.db {
		if strings.EqualFold(value.Transmission, transmissionType) {
			resp = append(resp, value)
		}
	}
	return resp, nil
}

func (r *VehicleMap) SearchVehicleByDimensions(minLen, maxLen, minWidth, maxWidth float64) ([]internal.Vehicle, error) {
	resp := make([]internal.Vehicle, 0)
	for _, v := range r.db {
		if v.Height <= minLen || v.Height >= maxLen {
			continue
		}
		if v.Width <= minWidth || v.Width >= maxWidth {
			continue
		}
		resp = append(resp, v)
	}
	return resp, nil
}

func (r *VehicleMap) SearchVehiclesByWeight(v float64, v2 float64) ([]internal.Vehicle, error) {
	resp := make([]internal.Vehicle, 0)
	for _, veh := range r.db {
		if veh.Weight <= v || veh.Weight >= v2 {
			continue
		}
		resp = append(resp, veh)
	}
	return resp, nil
}
