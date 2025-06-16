package repository_test

import (
	"errors"
	"fmt"
	"github.com/ribeirosaimon/bootcamp/internal/repository"
	"strconv"
	"testing"
)

func TestGetTotalTickets(t *testing.T) {
	repository := repository.NewRepository(
		repository.WithFilePath("./csvTest.csv"),
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
			name: "TestGetCountryByPeriod need return error",
			auxFunc: func() error {
				res, err := repository.GetCountryByPeriod("chasina")
				if err != nil {
					return err
				}
				if res != "\n        Início da Manhã: 0\n        Manhã: 0\n        Tarde: 0\n        Noite: 0\n    " {
					return fmt.Errorf("expected empty string, got %s", res)
				}
				return nil
			},
		},
		{
			name: "AverageDestination",
			auxFunc: func() error {
				res, err := repository.AverageDestination("china", 10)
				if err != nil {
					return err
				}
				if res != 30 {
					return errors.New("expected 30 tickets, got " + strconv.Itoa(res))
				}
				return nil
			},
		},
		{
			name:  "AverageDestination need return error",
			isErr: true,
			auxFunc: func() error {
				res, err := repository.AverageDestination("chasina", 10)
				if err != nil {
					return err
				}
				if res != 30 {
					return errors.New("expected 30 tickets, got " + strconv.Itoa(res))
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
