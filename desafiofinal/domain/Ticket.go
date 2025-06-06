package domain

import (
	"github.com/ribeirosaimon/bootcamp/desafiofinal/util"
	"strconv"
	"time"
)

type Ticket struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Destination string    `json:"destination"`
	Arrival     time.Time `json:"arrival"`
	Price       int64     `json:"price"`
}

func (t *Ticket) NormalizedTicket(s []string) error {
	id, err := strconv.Atoi(s[0])
	if err != nil {
		return err
	}
	t.Id = int64(id)
	t.Name = s[1]
	t.Email = s[2]
	t.Destination = s[3]
	val, err := util.NormalizeTime(s[4])
	if err != nil {
		return err
	}
	t.Arrival = val

	price, err := strconv.ParseInt(s[5], 10, 64)
	if err != nil {
		return err
	}
	t.Price = price
	return nil
}
