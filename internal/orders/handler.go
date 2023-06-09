package orders

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/bulhoes1998/irpf-go-api/utils"
	"github.com/go-chi/chi"
)

type OrderHandler interface {
	NewOrder(w http.ResponseWriter, r *http.Request)
	NewDayTrade(w http.ResponseWriter, r *http.Request)
	GetOrderById(w http.ResponseWriter, r *http.Request)
}

type orderHandler struct {
	orderRepository OrderRepository
}

func NewOrderHandler(orderRepository OrderRepository) *orderHandler {
	return &orderHandler{
		orderRepository: orderRepository,
	}
}

func (oh *orderHandler) NewOrder(w http.ResponseWriter, r *http.Request) {
	var order *Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		utils.Encoder(w, nil, http.StatusBadRequest)
		return
	}

	if _, err := time.Parse(time.RFC3339, order.Datetime); err != nil {
		utils.Encoder(w, nil, http.StatusBadRequest)
		return
	}

	if err := oh.orderRepository.CreateSwingTradeOrder(order); err != nil {
		utils.Encoder(w, nil, http.StatusInternalServerError)
		return
	}

	utils.Encoder(w, order, http.StatusCreated)
}

func (oh *orderHandler) NewDayTrade(w http.ResponseWriter, r *http.Request) {
	var dayTrade *DayTrade
	if err := json.NewDecoder(r.Body).Decode(&dayTrade); err != nil {
		utils.Encoder(w, nil, http.StatusBadRequest)
		return
	}

	if err := oh.orderRepository.CreateDayTradeOrder(dayTrade); err != nil {
		utils.Encoder(w, nil, http.StatusInternalServerError)
		return
	}

	utils.Encoder(w, dayTrade, http.StatusCreated)
}

func (oh *orderHandler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	order, err := oh.orderRepository.GetOrderById(id)
	if err != nil {
		utils.Encoder(w, nil, http.StatusInternalServerError)
		return
	}

	utils.Encoder(w, order, http.StatusOK)
}
