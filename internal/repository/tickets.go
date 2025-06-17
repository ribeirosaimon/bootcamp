package repository

import (
	"context"
	"github.com/ribeirosaimon/bootcamp/internal/domain"
	"github.com/ribeirosaimon/bootcamp/internal/loader"
)

// NewTicket creates a new repository for tickets in a map
func NewTicket(loader loader.TicketCSV) *ticketMap {
	load, err := loader.Load()
	if err != nil {
		panic(err)
	}

	return &ticketMap{
		db:     load,
		lastId: len(load),
	}
}

// RepositoryTicketMap implements the repository interface for tickets in a map
type ticketMap struct {
	// db represents the database in a map
	// - key: id of the ticket
	// - value: ticket
	db map[int]domain.Ticket

	// lastId represents the last id of the ticket
	lastId int
}

// Get returns all the tickets
func (r *ticketMap) Get(ctx context.Context) (t map[int]domain.Ticket, err error) {
	// create a copy of the map
	t = make(map[int]domain.Ticket, len(r.db))
	for k, v := range r.db {
		t[k] = v
	}

	return
}

// GetTicketsByDestinationCountry returns the tickets filtered by destination country
func (r *ticketMap) GetTicketsByDestinationCountry(ctx context.Context, country string) (t map[int]domain.Ticket, err error) {
	// create a copy of the map
	t = make(map[int]domain.Ticket)
	for k, v := range r.db {
		if v.Country == country {
			t[k] = v
		}
	}

	return
}

func (r *ticketMap) GetTicketsAmountByDestinationCountry(ctx context.Context, country string) (int, error) {
	t, err := r.GetTicketsByDestinationCountry(ctx, country)
	if err != nil {
		return 0, err
	}
	return len(t), nil
}
