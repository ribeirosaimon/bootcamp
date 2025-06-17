package service

import (
	"context"
	"github.com/ribeirosaimon/bootcamp/internal"
)

// ServiceTicketDefault represents the default service of the tickets
type ticketDefault struct {
	// rp represents the repository of the tickets
	rp internal.RepositoryTicket
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp internal.RepositoryTicket) *ticketDefault {
	return &ticketDefault{
		rp: rp,
	}
}

// GetTotalAmountTickets returns the total number of tickets
func (s *ticketDefault) GetTotalAmountTickets(ctx context.Context) (total int, err error) {
	t, err := s.rp.Get(ctx)
	if err != nil {
		return 0, err
	}
	return len(t), nil
}

// GetTicketsAmountByDestinationCountry returns the total number of tickets by country
func (s *ticketDefault) GetTicketsAmountByDestinationCountry(ctx context.Context, country string) (int, error) {
	return s.rp.GetTicketsAmountByDestinationCountry(ctx, country)
}

func (s *ticketDefault) GetPercentageTicketsByDestinationCountry(ctx context.Context, country string) (float64, error) {
	destinationCountry, err := s.rp.GetTicketsAmountByDestinationCountry(ctx, country)
	if err != nil {
		return 0, err
	}
	total, err := s.GetTotalAmountTickets(ctx)
	return float64(destinationCountry) / float64(total) * 100, nil
}
