package tickets_test

import (
	"fmt"
	"github.com/ribeirosaimon/bootcamp/desafiofinal/tickets"
	"testing"
)

func TestGetTotalTickets(t *testing.T) {
	repository := tickets.NewRepository(
		tickets.WithFilePath("./csvTest.csv"),
	)
	for _, v := range []struct {
		name    string
		auxFunc func() error
		isErr   bool
	}{
		{
			name: "TestGetTotalTickets",
			auxFunc: func() error {
				totalTickets, err := repository.GetTotalTickets("China")
				if err != nil {
					return err
				}
				if totalTickets != 3 {
					return fmt.Errorf("expected 3 tickets, got %d", totalTickets)
				}
				return nil
			},
		},
		{
			name:  "Need return a error",
			isErr: true,
			auxFunc: func() error {
				_, err := repository.GetTotalTickets("1234")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			name: "TestGetCountryByPeriod",
			auxFunc: func() error {
				_, err := repository.GetCountryByPeriod("china")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			name:  "TestGetCountryByPeriod need return error",
			isErr: true,
			auxFunc: func() error {
				_, err := repository.GetCountryByPeriod("china")
				if err != nil {
					return err
				}
				return nil
			},
		},
	} {
		t.Run(v.name, func(t *testing.T) {
			err := v.auxFunc()
			if v.isErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
			}
		})
	}
}
