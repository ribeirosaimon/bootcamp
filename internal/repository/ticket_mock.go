package repository

import (
	"context"
	"github.com/ribeirosaimon/bootcamp/internal/domain"
)

// NewRepositoryTicketMock creates a new repository for tickets in a map
func NewRepositoryTicketMock() *repositoryTicketMock {
	return &repositoryTicketMock{}
}

// RepositoryTicketMock implements the repository interface for tickets
type repositoryTicketMock struct {
	// FuncGet represents the mock for the Get function
	FuncGet func() (t map[int]domain.Ticket, err error)
	// FuncGetTicketsByDestinationCountry
	FuncGetTicketsByDestinationCountry       func(ctx context.Context, country string) (t map[int]domain.Ticket, err error)
	FuncGetTicketsAmountByDestinationCountry func(context.Context, string) (int, error)

	// Spy verifies if the methods were called
	Spy struct {
		// Get represents the spy for the Get function
		Get int
		// GetTicketsByDestinationCountry represents the spy for the GetTicketsByDestinationCountry function
		GetTicketsByDestinationCountry       int
		GetTicketsAmountByDestinationCountry int
	}
}

// Get returns all the tickets
func (r *repositoryTicketMock) Get(ctx context.Context) (t map[int]domain.Ticket, err error) {
	// spy
	r.Spy.Get++

	// mock
	t, err = r.FuncGet()
	return
}

// GetTicketsByDestinationCountry returns the tickets filtered by destination country
func (r *repositoryTicketMock) GetTicketsByDestinationCountry(ctx context.Context, country string) (
	t map[int]domain.Ticket, err error,
) {
	// spy
	r.Spy.GetTicketsByDestinationCountry++

	// mock
	t, err = r.FuncGetTicketsByDestinationCountry(ctx, country)
	return
}

// GetTicketsAmountByDestinationCountry returns total of value
func (r *repositoryTicketMock) GetTicketsAmountByDestinationCountry(ctx context.Context, country string) (total int, err error) {
	r.Spy.GetTicketsAmountByDestinationCountry++
	// mock
	return r.FuncGetTicketsAmountByDestinationCountry(ctx, country)
}
