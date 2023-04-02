package calculator

import (
	"database/sql"

	"github.com/bulhoes1998/irpf-go-api/internal/orders"
	"github.com/bulhoes1998/irpf-go-api/utils"
)

type CalculatorRepository interface {
	CalculateDarfByMonth(month string, year string) error
}

type calculatorRepository struct {
	orderRepository orders.OrderRepository
}

func NewCalculatorRepository(orderRepository orders.OrderRepository, db *sql.DB) *calculatorRepository {
	return &calculatorRepository{
		orderRepository: orderRepository,
	}
}

func (cr *calculatorRepository) CalculateDarfByMonth(month string, year string) (DARF, error) {
	tradeList, err := cr.orderRepository.GetDayTradeListByMonth(month, year)
	if err != nil {
		return DARF{}, err
	}
	var totalValue float64
	for _, s := range tradeList {
		totalValue += (s.SellPrice - s.BuyPrice) * float64(s.Quantity)
	}
	return DARF{
		TotalValue: totalValue,
		Tax:        totalValue * utils.DayTradeTax,
	}, nil
}
