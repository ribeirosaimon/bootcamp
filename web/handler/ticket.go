package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/ribeirosaimon/bootcamp/internal"
	"github.com/ribeirosaimon/bootcamp/web/response"
	"net/http"
)

type ticket struct {
	group    string
	service  internal.ServiceTicket
	handlers []BasicHandler
}

func (h *ticket) GetHandlers() []BasicHandler {
	return h.handlers
}

func (h *ticket) GetGroup() string {
	return h.group
}

func NewTicket(sv internal.ServiceTicket) *ticket {
	h := ticket{
		group:   "/ticket",
		service: sv,
	}
	h.handlers = []BasicHandler{
		{Method: http.MethodGet, Path: "/get_by_country/{dest}", Handler: h.GetTotalAmountTickets()},
		{Method: http.MethodGet, Path: "/get_average/{dest}", Handler: h.GetPercentageTicketsByDestinationCountry()},
	}
	return &h
}

func (h *ticket) GetTotalAmountTickets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			total   int
			country string
			err     error
		)

		country = chi.URLParam(r, "dest")
		if country == "" {
			total, err = h.service.GetTotalAmountTickets(r.Context())
			if err != nil {
				response.Error(w, http.StatusBadRequest, "error getting total amount tickets")
				return
			}
			response.JSON(w, http.StatusOK, total)
			return
		}

		total, err = h.service.GetTicketsAmountByDestinationCountry(r.Context(), country)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "error getting total amount tickets")
			return
		}
		response.JSON(w, http.StatusOK, total)
		return
	}
}

func (h *ticket) GetPercentageTicketsByDestinationCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		country := chi.URLParam(r, "dest")
		if country == "" {
			response.Error(w, http.StatusBadRequest, "country is required")
			return
		}
		total, err := h.service.GetPercentageTicketsByDestinationCountry(r.Context(), country)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "error getting total amount tickets")
			return
		}
		response.JSON(w, http.StatusOK, total)
	}
}
