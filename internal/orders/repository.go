package orders

import (
	"strconv"
	"time"
)

type OrderRepository interface {
	CreateSwingTradeOrder(order *Order) error
	CreateDayTradeOrder(trade *DayTrade) error
	GetDayTradeListByMonth(month string, year string) ([]DayTrade, error)
}

type orderRepository struct {
	orderStory    *OrderStore
	dayTradeStore *DayTradeStore
}

func NewOrderRepository(orderStore OrderStore, dayTradeStore DayTradeStore) *orderRepository {
	return &orderRepository{
		orderStory:    &orderStore,
		dayTradeStore: &dayTradeStore,
	}
}

func (or *orderRepository) CreateSwingTradeOrder(order *Order) error {
	order.OrderNumber = len(or.orderStory.db)
	or.orderStory.db = append(or.orderStory.db, *order)
	return nil
}

func (or *orderRepository) CreateDayTradeOrder(trade *DayTrade) error {
	trade.OperationNumber = len(or.dayTradeStore.db)
	or.dayTradeStore.db = append(or.dayTradeStore.db, *trade)
	return nil
}

func (or *orderRepository) GetDayTradeListByMonth(month string, year string) ([]DayTrade, error) {
	var tradeList []DayTrade
	for _, s := range or.dayTradeStore.db {

		date, _ := time.Parse(time.RFC3339, s.Datetime)
		numericYear, _ := strconv.Atoi(year)
		numericMonth, _ := strconv.Atoi(month)

		if date.Year() == numericYear && int(date.Month()) == numericMonth {
			tradeList = append(tradeList, s)
		}
	}

	return tradeList, nil
}
