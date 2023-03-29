package app

import (
	"net/http"

	"github.com/bulhoes1998/irpf-go-api/internal/calculator"
	"github.com/bulhoes1998/irpf-go-api/internal/orders"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Handlers struct {
	orders     orders.OrderHandler
	calculator calculator.CalculatorHandler
}

func BuildApplication() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	var (
		orderStore    orders.OrderStore
		dayTradeStore orders.DayTradeStore
	)
	ordersRepository := orders.NewOrderRepository(orderStore, dayTradeStore)
	calculatorRepository := calculator.NewCalculatorRepository(ordersRepository)

	ctrl := &Handlers{
		orders:     orders.NewOrderHandler(ordersRepository),
		calculator: calculator.NewCalculatorHandler(calculatorRepository),
	}

	routes(r, ctrl)

	http.ListenAndServe(":8080", r)
}
