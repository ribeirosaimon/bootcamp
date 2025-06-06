package tickets

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/ribeirosaimon/bootcamp/desafiofinal/domain"
	"os"
	"strings"
	"time"
)

var (
	fileName = "/desafiofinal/tickets.csv"
)

type Ticket interface {
	GetTotalTickets(destination string) (int, error)
	GetCountryByPeriod(country string) (string, error)
	AverageDestination(destination string, total int) (int, error)
}
type repository struct {
	filePath string
	data     map[string][]domain.Ticket
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

	basicData := make(map[string][]domain.Ticket)
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
		basicData[strings.ToLower(ticket.Destination)] = append(basicData[strings.ToLower(ticket.Destination)], ticket)

	}

	defaultRepository.data = basicData
	return &defaultRepository
}

func (t *repository) GetTotalTickets(destination string) (int, error) {
	if _, ok := t.data[strings.ToLower(destination)]; !ok {
		return 0, fmt.Errorf("no tickets found for destination %s", destination)
	}
	return len(t.data[strings.ToLower(destination)]), nil

}

func (t *repository) GetCountryByPeriod(country string) (string, error) {
	var (
		sunset                                           = time.Date(0, 1, 1, 6, 0, 0, 0, time.Local)
		noon                                             = time.Date(0, 1, 1, 12, 0, 0, 0, time.Local)
		afternoon                                        = time.Date(0, 1, 1, 19, 0, 0, 0, time.Local)
		startMorn, mornCount, afternoonCount, nightCount = 0, 0, 0, 0
	)

	for _, ticket := range t.data[strings.ToLower(country)] {
		switch {
		case ticket.Arrival.Before(sunset):
			startMorn++
		case ticket.Arrival.After(sunset) && ticket.Arrival.Before(noon):
			mornCount++
		case ticket.Arrival.After(noon) && ticket.Arrival.Before(afternoon):
			afternoonCount++
		case ticket.Arrival.After(afternoon):
			nightCount++
		default:
			return "", errors.New("not found period")
		}
	}

	return fmt.Sprintf(`
        Início da Manhã: %d
        Manhã: %d
        Tarde: %d
        Noite: %d
    `, startMorn, mornCount, afternoonCount, nightCount), nil
}

func (t *repository) AverageDestination(destination string, total int) (int, error) {
	period, err := t.GetTotalTickets(destination)
	if err != nil {
		return 0, err
	}
	result := float64(period) / float64(total)
	return int(result * 100), nil
}
