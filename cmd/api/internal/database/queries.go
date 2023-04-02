package database

var DBqueries = []string{
	"CREATE TABLE IF NOT EXISTS ORDERS (id INTEGER PRIMARY KEY AUTOINCREMENT, order_type TEXT, stock_name TEXT, quantity INTEGER, datetime TEXT)",
	"CREATE TABLE IF NOT EXISTS TRADES (id INTEGER PRIMARY KEY AUTOINCREMENT, stock_name TEXT, quantity INTEGER, buy_price FLOAT, sell_price FLOAT, datetime TEXT)",
}
