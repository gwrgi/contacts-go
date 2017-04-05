package types

import (
	"time"
)

type Expense struct {
	Merchant   string    `json:"merchant"`
	Amount     float64   `json:"amount"`
	CategoryId int       `json:"categoryId"`
	Date       time.Time `json:"date"`
}
