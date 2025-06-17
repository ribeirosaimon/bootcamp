package domain

import (
	"github.com/ribeirosaimon/bootcamp/util"
	"strconv"
	"time"
)

// Ticket is a struct that represents a ticket
type Ticket struct {
	Id int `json:"id"`
	// Name represents the name of the owner of the ticket
	Name string `json:"name"`
	// Email represents the email of the owner of the ticket
	Email string `json:"email"`
	// Country represents the destination country of the ticket
	Country string `json:"country"`
	// Hour represents the hour of the ticket
	Hour time.Time `json:"hour"`
	// Price represents the price of the ticket
	Price float64 `json:"price"`
}

func (t *Ticket) NormalizedTicket(s []string) error {
	id, err := strconv.Atoi(s[0])
	if err != nil {
		return err
	}
	t.Id = id
	t.Name = s[1]
	t.Email = s[2]
	t.Country = s[3]
	val, err := util.NormalizeTime(s[4])
	if err != nil {
		return err
	}
	t.Hour = val

	price, err := strconv.ParseFloat(s[5], 64)
	if err != nil {
		return err
	}
	t.Price = price
	return nil
}
