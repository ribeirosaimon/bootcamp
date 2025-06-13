package application

import (
	"github.com/ribeirosaimon/bootcamp/internal/handler"
	"github.com/ribeirosaimon/bootcamp/internal/loader"
	"github.com/ribeirosaimon/bootcamp/internal/repository"
	"github.com/ribeirosaimon/bootcamp/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// ConfigServerChi is a struct that represents the configuration for ServerChi
type ConfigServerChi struct {
	// ServerAddress is the address where the server will be listening
	ServerAddress string
	// LoaderFilePath is the path to the file that contains the vehicles
	LoaderFilePath string
}

// NewServerChi is a function that returns a new instance of ServerChi
func NewServerChi(cfg *ConfigServerChi) *ServerChi {
	// default values
	defaultConfig := &ConfigServerChi{
		ServerAddress: ":8080",
	}
	if cfg != nil {
		if cfg.ServerAddress != "" {
			defaultConfig.ServerAddress = cfg.ServerAddress
		}
		if cfg.LoaderFilePath != "" {
			defaultConfig.LoaderFilePath = cfg.LoaderFilePath
		}
	}

	return &ServerChi{
		serverAddress:  defaultConfig.ServerAddress,
		loaderFilePath: defaultConfig.LoaderFilePath,
	}
}

// ServerChi is a struct that implements the Application interface
type ServerChi struct {
	// serverAddress is the address where the server will be listening
	serverAddress string
	// loaderFilePath is the path to the file that contains the vehicles
	loaderFilePath string
}

// Run is a method that runs the application
func (a *ServerChi) Run() (err error) {
	// dependencies
	// - loader
	ld := loader.NewVehicleJSONFile(a.loaderFilePath)
	db, err := ld.Load()
	if err != nil {
		return
	}

	// - repository
	rp := repository.NewVehicleMap(db)
	// - service
	sv := service.NewVehicleDefault(rp)
	// - handler
	hd := handler.NewVehicleDefault(sv)
	// router
	rt := chi.NewRouter()
	// - middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)
	// - endpoints
	rt.Route("/vehicles", func(rt chi.Router) {
		// - GET /vehicles
		rt.Get("/{id}", hd.FindByID())
		//rt.Get("/", hd.GetAll())
		rt.Post("/", hd.Save())
		rt.Get("/", hd.GetVehicleByColorAndYear())
		rt.Get("/brand/{brand}/between/{start_year}/{end_year}", hd.GetVehicleByBrandAndYear())
		rt.Get("/average_speed/brand/{brand}", hd.GetAverageSpeed())
		rt.Post("/batch", hd.SaveMany())
		rt.Put("/{id}/update_speed", hd.UpdateSpeed())
		rt.Get("/fuel_type/{type}", hd.GetVehicleByFuelType())
		rt.Delete("/{id}", hd.DeleteVehicle())
		rt.Get("/transmission/{type}", hd.GetVehicleByTransmission())
		rt.Put("/{id}/update_fuel", hd.UpdateFuel())
		rt.Get("/average_capacity/brand/{brand}", hd.GetAverageCapacity())
		rt.Get("/dimensions", hd.SearchVehicleByDimensions())
		rt.Get("/weight", hd.SearchVehicleByWeight())
	})

	// run server
	err = http.ListenAndServe(a.serverAddress, rt)
	return
}
