package main

import "github.com/ribeirosaimon/bootcamp/desafiofinal/tickets"

type desafioFinal struct {
	repository tickets.Ticket
}

func main() {
	v := desafioFinal{
		repository: tickets.NewRepository(),
	}

	// aqui posso fazer um looping infinito recebendo o valor de um cli
	v.repository.GetTotalTickets("brazil")
	v.repository.AverageDestination("brazil", 123)
	v.repository.AverageDestination("brazil", 1)
}
