package handler

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/ribeirosaimon/bootcamp/internal"
	"github.com/ribeirosaimon/bootcamp/web"
	"github.com/ribeirosaimon/bootcamp/web/apierr"
	"net/http"
	"strconv"
	"strings"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

func (v *VehicleJSON) ToDomain() internal.Vehicle {
	return internal.Vehicle{
		Id: v.ID,
		VehicleAttributes: internal.VehicleAttributes{
			Brand:           v.Brand,
			Model:           v.Model,
			Registration:    v.Registration,
			Color:           v.Color,
			FabricationYear: v.FabricationYear,
			Capacity:        v.Capacity,
			MaxSpeed:        v.MaxSpeed,
			FuelType:        v.FuelType,
			Transmission:    v.Transmission,
			Weight:          v.Weight,
			Dimensions: internal.Dimensions{
				Height: v.Height,
				Width:  v.Width,
				Length: v.Length,
			},
		},
	}
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

func (h *VehicleDefault) FindByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			apierr.NewApiErr(apierr.NewBadRequestApiErr(), w)
			return
		}
		byID, err := h.sv.FindByID(id)
		if err != nil {
			apierr.NewApiErr(err, w)
			return
		}
		web.NewResponse(
			web.WithData(byID),
			web.WithStatus(http.StatusOK),
		).Build(w)
	}
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			apierr.NewApiErr(err, w)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		web.NewResponse(
			web.WithData(data),
			web.WithStatus(http.StatusOK),
		).Build(w)
	}
}

// Save was to save one vehicle
func (h *VehicleDefault) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var vehicle VehicleJSON
		if err := json.NewDecoder(r.Body).Decode(&vehicle); err != nil {
			apierr.NewApiErr(apierr.NewBadRequestApiErr(), w)
			return
		}
		vehicleDomain := vehicle.ToDomain()
		if err := h.sv.Save(&vehicleDomain); err != nil {
			apierr.NewApiErr(err, w)
			return
		}
		web.NewResponse(
			web.WithData(&vehicleDomain),
			web.WithStatus(http.StatusCreated),
		).Build(w)
	}
}

// GetVehicleByColorAndYear get a vehicle by color and year
func (h *VehicleDefault) GetVehicleByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		yearString := r.URL.Query().Get("year")

		var (
			year int
			err  error
		)

		if yearString != "" {
			year, err = strconv.Atoi(yearString)
			if err != nil {
				apierr.NewApiErr(err, w)
			}
		}

		color := r.URL.Query().Get("color")
		res, err := h.sv.SearchVehiclesByColorAndYear(internal.VehicleFilter{
			Year:  year,
			Color: color,
		})
		if err != nil {
			apierr.NewApiErr(err, w)
			return
		}
		web.NewResponse(
			web.WithData(res),
			web.WithStatus(http.StatusOK),
		).Build(w)
	}
}

// GetVehicleByBrandAndYear get vehicle by brand and year
func (h *VehicleDefault) GetVehicleByBrandAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")
		if brand == "" {
			apierr.NewApiErr(errors.New("brand is required"), w)
		}
		startStringDate := chi.URLParam(r, "start_year")
		start, err := strconv.Atoi(startStringDate)
		if err != nil {
			apierr.NewApiErr(err, w)
		}
		endStringDate := chi.URLParam(r, "end_year")
		end, err := strconv.Atoi(endStringDate)
		if err != nil {
			apierr.NewApiErr(err, w)
		}

		if start > end {
			end, start = start, end
		}

		res, err := h.sv.SearchVehicleByBrandAndYear(internal.VehicleFilter{
			Brand: brand,
			PairDates: internal.PairDates{
				Start: start,
				End:   end,
			},
		})
		if err != nil {
			apierr.NewApiErr(err, w)
		}

		web.NewResponse(
			web.WithData(res),
			web.WithStatus(http.StatusOK),
		).Build(w)

	}
}

// GetAverageSpeed get average speed
func (h *VehicleDefault) GetAverageSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")
		res, err := h.sv.GetAverageSpeed(brand)
		if err != nil {
			apierr.NewApiErr(err, w)
		}
		web.NewResponse(
			web.WithData(struct {
				Average float64 `json:"average"`
			}{
				Average: res,
			}),
			web.WithStatus(http.StatusOK),
		).Build(w)
	}
}

// SaveMany save many vehicles in single time
func (h *VehicleDefault) SaveMany() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var vehicles []VehicleJSON
		if err := json.NewDecoder(r.Body).Decode(&vehicles); err != nil {
			apierr.NewApiErr(apierr.NewBadRequestApiErr(), w)
		}
		vehiclesDomain := make([]internal.Vehicle, 0, len(vehicles))
		for _, vehicle := range vehicles {
			vehiclesDomain = append(vehiclesDomain, vehicle.ToDomain())
		}

		if err := h.sv.SaveMany(vehiclesDomain); err != nil {
			apierr.NewApiErr(err, w)
			return
		}

		web.NewResponse(
			web.WithData(vehiclesDomain),
			web.WithStatus(http.StatusCreated),
		).Build(w)
	}
}

func (h *VehicleDefault) UpdateSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vehicleId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(vehicleId)
		if err != nil {
			apierr.NewApiErr(apierr.NewBadRequestApiErr(), w)
			return
		}
		var speedDto struct {
			Speed float64 `json:"speed"`
		}

		if err = json.NewDecoder(r.Body).Decode(&speedDto); err != nil {
			apierr.NewApiErr(apierr.NewBadRequestApiErr(), w)
			return
		}
		if err = h.sv.UpdateSpeed(id, speedDto.Speed); err != nil {
			apierr.NewApiErr(err, w)
			return
		}

		web.NewResponse(
			web.WithStatus(http.StatusOK),
		).Build(w)
	}
}

func (h *VehicleDefault) GetVehicleByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fuelType := chi.URLParam(r, "type")
		if fuelType == "" {
			apierr.NewApiErr(apierr.NewBadRequestApiErr(), w)
			return
		}
		resp, err := h.sv.GetVehicleByFuelType(fuelType)
		if err != nil {
			apierr.NewApiErr(err, w)
			return
		}

		web.NewResponse(
			web.WithData(resp),
			web.WithStatus(http.StatusOK),
		).Build(w)
	}
}

func (h *VehicleDefault) DeleteVehicle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vehicleId := chi.URLParam(r, "id")
		id, err := strconv.Atoi(vehicleId)
		if err != nil {
			apierr.NewApiErr(apierr.NewBadRequestApiErr(), w)
			return
		}
		if err = h.sv.DeleteVehicle(id); err != nil {
			apierr.NewApiErr(err, w)
			return
		}
		web.NewResponse(
			web.WithStatus(http.StatusOK),
		).Build(w)
	}
}

func (h *VehicleDefault) GetVehicleByTransmission() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transmissionType := chi.URLParam(r, "type")
		byTransmissionType, err := h.sv.SearchVehiclesByTransmissionType(transmissionType)
		if err != nil {
			apierr.NewApiErr(err, w)
			return
		}
		web.NewResponse(
			web.WithData(byTransmissionType),
			web.WithStatus(http.StatusOK),
		).Build(w)
	}
}

func (h *VehicleDefault) UpdateFuel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			apierr.NewApiErr(apierr.NewBadRequestApiErr(), w)
			return
		}
		fuelDto := struct {
			Fuel string `json:"fuel"`
		}{}

		if err = json.NewDecoder(r.Body).Decode(&fuelDto); err != nil {
			apierr.NewApiErr(apierr.NewBadRequestApiErr(), w)
			return
		}

		if err = h.sv.UpdateFuel(id, fuelDto.Fuel); err != nil {
			apierr.NewApiErr(err, w)
			return
		}

		web.NewResponse(
			web.WithStatus(http.StatusOK),
		)
	}
}

func (h *VehicleDefault) GetAverageCapacity() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")
		if brand == "" {
			apierr.NewApiErr(apierr.NewBadRequestApiErr(), w)
			return
		}
		capacity, err := h.sv.FindAverageCapacity(brand)
		if err != nil {
			apierr.NewApiErr(err, w)
			return
		}

		web.NewResponse(
			web.WithData(struct {
				AverageCapacity float64 `json:"average"`
			}{
				AverageCapacity: capacity,
			}),
			web.WithStatus(http.StatusOK),
		).Build(w)
	}
}

func (h *VehicleDefault) SearchVehicleByDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var minWidth, maxWidth, minLen, maxLen float64 = 0, 0, 0, 0

		for _, st := range []struct {
			Field    string
			MinValue *float64
			MaxValue *float64
		}{
			{Field: "width", MinValue: &minWidth, MaxValue: &maxWidth},
			{Field: "height", MinValue: &minLen, MaxValue: &maxLen},
		} {
			if err := normalizeValue(r.URL.Query().Get(st.Field), st.MinValue, st.MaxValue); err != nil {
				apierr.NewApiErr(apierr.NewBadRequestApiErr(), w)
				return
			}
		}
		dimensions, err := h.sv.SearchVehicleByDimensions(minLen, maxLen, minWidth, maxWidth)
		if err != nil {
			apierr.NewApiErr(err, w)
			return
		}

		web.NewResponse(
			web.WithData(dimensions),
			web.WithStatus(http.StatusOK),
		).Build(w)
	}
}

func (h *VehicleDefault) SearchVehicleByWeight() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		minValueString := r.URL.Query().Get("min")
		minV, err := strconv.ParseFloat(minValueString, 64)
		if err != nil {
			apierr.NewApiErr(apierr.NewBadRequestApiErr(), w)
			return
		}

		maxValueString := r.URL.Query().Get("max")
		maxV, err := strconv.ParseFloat(maxValueString, 64)
		if err != nil {
			apierr.NewApiErr(apierr.NewBadRequestApiErr(), w)
			return
		}

		if minV > maxV {
			minV, maxV = maxV, minV
		}

		weight, err := h.sv.SearchVehiclesByWeight(minV, maxV)
		if err != nil {
			apierr.NewApiErr(err, w)
			return
		}
		web.NewResponse(
			web.WithData(weight),
			web.WithStatus(http.StatusOK),
		).Build(w)
	}
}

func normalizeValue(v string, minValue, maxValue *float64) error {
	var (
		minV, maxV float64
		err        error
	)
	splitedString := strings.Split(v, "-")
	if len(splitedString) != 2 {
		return errors.New("input deve conter exatamente um '-'")
	}

	minV, err = strconv.ParseFloat(splitedString[0], 64)
	if err != nil {
		return err
	}

	maxV, err = strconv.ParseFloat(splitedString[1], 64)
	if err != nil {
		return err
	}
	*minValue = minV
	*maxValue = maxV
	return nil
}
