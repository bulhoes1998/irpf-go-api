package calculator

import (
	"net/http"

	"github.com/bulhoes1998/irpf-go-api/utils"
	"github.com/go-chi/chi"
)

type CalculatorHandler interface {
	CalculateDarfByMonth(w http.ResponseWriter, r *http.Request)
}

type calculatorHandler struct {
	calculatorRepository *calculatorRepository
}

func NewCalculatorHandler(calculatorRepository *calculatorRepository) *calculatorHandler {
	return &calculatorHandler{
		calculatorRepository: calculatorRepository,
	}
}

func (ch *calculatorHandler) CalculateDarfByMonth(w http.ResponseWriter, r *http.Request) {
	month := chi.URLParam(r, "month")
	year := chi.URLParam(r, "year")

	value, err := ch.calculatorRepository.CalculateDarfByMonth(month, year)
	if err != nil {
		utils.Encoder(w, nil, http.StatusInternalServerError)
		return
	}

	utils.Encoder(w, value, http.StatusOK)
}
