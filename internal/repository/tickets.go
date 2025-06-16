package repository

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"github.com/ribeirosaimon/bootcamp/internal/domain"
	"os"
)

var (
	fileName = "/docs/db/tickets.csv"
)

type Ticket interface {
	Get(ctx context.Context) (t map[int]domain.Ticket, err error)
	GetTicketsByDestinationCountry(ctx context.Context, country string) (t map[int]domain.Ticket, err error)
}
type repository struct {
	filePath string
	data     map[int]domain.Ticket
}

type ticketsOpt func(*repository)

func WithFilePath(filePath string) ticketsOpt {
	return func(r *repository) {
		r.filePath = filePath
	}
}

func NewRepository(opt ...ticketsOpt) *repository {
	defaultRepository := repository{
		filePath: fileName,
	}
	for _, o := range opt {
		o(&defaultRepository)
	}

	file, err := os.ReadFile(defaultRepository.filePath)
	if err != nil {
		return nil
	}
	buf := bytes.NewBuffer(file)
	reader := csv.NewReader(buf)

	basicData := make(map[int]domain.Ticket)
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Erro ao ler linha:", err)
			break
		}
		var ticket domain.Ticket

		if err = ticket.NormalizedTicket(record); err != nil {
			return nil
		}
		basicData[ticket.Id] = ticket

	}

	defaultRepository.data = basicData
	return &defaultRepository
}
