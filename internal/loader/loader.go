package loader

import (
	"encoding/csv"
	"fmt"
	"github.com/ribeirosaimon/bootcamp/internal/domain"
	"io"
	"os"
)

// NewLoaderTicketCSV creates a new ticket loader from a CSV file
func NewLoaderTicketCSV(filePath string) *loaderTicketCSV {
	return &loaderTicketCSV{
		filePath: filePath,
	}
}

type TicketCSV interface {
	Load() (t map[int]domain.Ticket, err error)
}

// LoaderTicketCSV represents a ticket loader from a CSV file
type loaderTicketCSV struct {
	filePath string
}

// Load loads the tickets from the CSV file
func (l *loaderTicketCSV) Load() (map[int]domain.Ticket, error) {
	// open the file
	f, err := os.Open(l.filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	// read the file
	r := csv.NewReader(f)

	// read the records
	v := make(map[int]domain.Ticket)
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, fmt.Errorf("error reading record: %v", err)
		}

		// serialize the record
		var domainTicket domain.Ticket
		if err = domainTicket.NormalizedTicket(record); err != nil {
			return nil, err
		}
		v[domainTicket.Id] = domainTicket

	}

	return v, nil
}
