package main

import "github.com/ribeirosaimon/bootcamp/desafiofinal/tickets"

type desafioFinalServer struct {
	repository tickets.Ticket
}

var server = desafioFinalServer{
	repository: tickets.NewRepository(),
}

func main() {
	// aqui posso fazer um looping infinito recebendo o valor como um cli
	country := "brazil"
	server.repository.GetTotalTickets(country)
	server.repository.GetCountryByPeriod(country)
	server.repository.AverageDestination(country, 1)
}
