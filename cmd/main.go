package main

import (
	"github.com/ribeirosaimon/bootcamp/internal/repository"
)

type desafioFinalServer struct {
	repository repository.Ticket
}

var server = desafioFinalServer{
	repository: repository.NewRepository(),
}

func main() {
	// aqui posso fazer um looping infinito recebendo o valor como um cli
	country := "brazil"
	server.repository.GetTotalTickets(country)
	server.repository.GetCountryByPeriod(country)
	server.repository.AverageDestination(country, 1)
}
