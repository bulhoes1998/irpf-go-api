package orders

import (
	"database/sql"
	"fmt"
)

type OrderRepository interface {
	CreateSwingTradeOrder(order *Order) error
	CreateDayTradeOrder(trade *DayTrade) error
	GetDayTradeListByMonth(month string, year string) ([]DayTrade, error)
	GetOrderById(id int) (*Order, error)
}

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *orderRepository {
	return &orderRepository{
		db: db,
	}
}

func (or *orderRepository) CreateSwingTradeOrder(order *Order) error {
	query := `INSERT INTO ORDERS(order_type, stock_name, quantity, datetime) VALUES (?, ?, ?, ?)`
	statement, err := or.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := statement.Exec(&order.OrderType, order.StockName, order.Quantity, order.Datetime)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil

}

func (or *orderRepository) CreateDayTradeOrder(trade *DayTrade) error {
	row, err := or.db.Prepare(`INSERT INTO TRADES(stock_name, quantity, buy_price, sell_price, datetime) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	res, err := row.Exec(&trade.StockName, &trade.Quantity, &trade.BuyPrice, &trade.SellPrice, &trade.Datetime)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func (or *orderRepository) GetDayTradeListByMonth(month string, year string) ([]DayTrade, error) {
	rows, err := or.db.Query(`SELECT * FROM TRADES WHERE datetime LIKE ?`, year+"-"+month+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tradeList []DayTrade
	for rows.Next() {
		i := DayTrade{}
		err := rows.Scan(&i.OperationNumber, &i.StockName, &i.Quantity, &i.BuyPrice, &i.SellPrice, &i.Datetime)
		if err != nil {
			return nil, err
		}
		tradeList = append(tradeList, i)
	}

	return tradeList, nil
}

func (or *orderRepository) GetOrderById(id int) (*Order, error) {
	query := or.db.QueryRow(`SELECT * FROM ORDERS WHERE id=?`, id)
	order := Order{}

	if err := query.Scan(&order.OrderNumber, &order.OrderType, &order.StockName, &order.Quantity, &order.Datetime); err != nil {
		return &Order{}, err
	}

	return &order, nil
}
