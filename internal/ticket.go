package internal

import (
	"context"
	"github.com/ribeirosaimon/bootcamp/internal/domain"
)

// Ticket represents a ticket
type Ticket struct {
	// Id represents the id of the ticket
	Id int `json:"id"`
	// Attributes represents the attributes of the ticket
	Attributes domain.Ticket `json:"attributes"`
}

// RepositoryTicket represents the repository interface for tickets
type RepositoryTicket interface {
	// Get returns all the tickets
	Get(ctx context.Context) (t map[int]domain.Ticket, err error)

	// GetTicketsByDestinationCountry returns the tickets filtered by destination country
	GetTicketsByDestinationCountry(ctx context.Context, country string) (t map[int]domain.Ticket, err error)

	// GetTicketsAmountByDestinationCountry return amount ticket by destination
	GetTicketsAmountByDestinationCountry(context.Context, string) (int, error)
}

type ServiceTicket interface {
	// GetTotalAmountTickets returns the total amount of tickets
	GetTotalAmountTickets(ctx context.Context) (total int, err error)

	// GetTicketsAmountByDestinationCountry returns the amount of tickets filtered by destination country
	GetTicketsAmountByDestinationCountry(ctx context.Context, country string) (int, error)

	// GetPercentageTicketsByDestinationCountry returns the percentage of tickets filtered by destination country
	GetPercentageTicketsByDestinationCountry(ctx context.Context, country string) (float64, error)
}
