package app

import (
	"github.com/go-chi/chi"
)

func routes(r *chi.Mux, handlers *Handlers) {
	r.Post("/swing_trade/new", handlers.orders.NewOrder)
	r.Post("/day_trade/new", handlers.orders.NewDayTrade)
	r.Get("/darf/{month}/{year}", handlers.calculator.CalculateDarfByMonth)
}
