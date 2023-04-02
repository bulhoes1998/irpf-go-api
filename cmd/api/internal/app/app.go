package app

import (
	"log"
	"net/http"

	// Driver for SQLite3 database
	_ "github.com/mattn/go-sqlite3"

	"github.com/bulhoes1998/irpf-go-api/cmd/api/internal/database"
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
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	ordersRepository := orders.NewOrderRepository(db)
	calculatorRepository := calculator.NewCalculatorRepository(ordersRepository, db)

	ctrl := &Handlers{
		orders:     orders.NewOrderHandler(ordersRepository),
		calculator: calculator.NewCalculatorHandler(calculatorRepository),
	}

	routes(r, ctrl)

	http.ListenAndServe(":8080", r)
}
