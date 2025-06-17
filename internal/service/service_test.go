package service_test

import (
	"context"
	"github.com/ribeirosaimon/bootcamp/internal/domain"
	"github.com/ribeirosaimon/bootcamp/internal/repository"
	"github.com/ribeirosaimon/bootcamp/internal/service"
	"testing"
	"time"
)

// Tests for ServiceTicketDefault.GetTotalAmountTickets
func TestServiceTicketDefault_GetTotalAmountTickets(t *testing.T) {
	ctx := context.Background()
	// arrange
	// - repository: mock
	rp := repository.NewRepositoryTicketMock()
	response := map[int]domain.Ticket{
		1: {
			Name:    "John",
			Email:   "johndoe@gmail.com",
			Country: "USA",
			Hour:    time.Now(),
			Price:   100,
		},
		3: {
			Name:    "John",
			Email:   "johndoe@gmail.com",
			Country: "Brazil",
			Hour:    time.Now(),
			Price:   100,
		},
	}

	// - repository: set-up
	rp.FuncGet = func() (t map[int]domain.Ticket, err error) {
		return response, nil
	}

	rp.FuncGetTicketsAmountByDestinationCountry = func(context.Context, string) (int, error) {
		return 2, nil
	}

	// - service
	sv := service.NewServiceTicketDefault(rp)

	for _, tt := range []struct {
		name    string
		want    any
		auxFunc func() (any, error)
	}{
		{
			name: "success to get total tickets",
			want: 2,
			auxFunc: func() (any, error) {
				return sv.GetTotalAmountTickets(ctx)
			},
		},
		{
			name: "success to get total tickets",
			want: 1,
			auxFunc: func() (any, error) {
				_, err := sv.GetTicketsAmountByDestinationCountry(ctx, "Brazil")
				return rp.Spy.GetTicketsAmountByDestinationCountry, err
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.auxFunc()
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if res != tt.want {
				t.Errorf("got %v, want %v", res, tt.want)
			}
		})
	}
}
