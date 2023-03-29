package orders

type Order struct {
	OrderNumber int    `json:"order_number"`
	OrderType   string `json:"order_type"`
	StockName   string `json:"stock_name"`
	Quantity    int16  `json:"quantity"`
	Datetime    string `json:"datetime"`
}

type DayTrade struct {
	OperationNumber int     `json:"operation_number"`
	StockName       string  `json:"stock_name,omitempty"`
	Quantity        int     `json:"quantity,omitempty"`
	BuyPrice        float64 `json:"buy_price,omitempty"`
	SellPrice       float64 `json:"sell_price,omitempty"`
	Datetime        string  `json:"datetime,omitempty"`
}

type OrderStore struct {
	db []Order
}

type DayTradeStore struct {
	db []DayTrade
}
