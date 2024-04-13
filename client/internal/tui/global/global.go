package global

import (
	"github.com/ElrohirGT/Ratatouille/internal/db"
)

var (
	Id   int
	Role string

	Driver *db.Queries
)

type ErrorDB struct {
	Description string
}
type SuccesDB struct {
	Description string
	Value       SuccessValue
}

type PaymentSuccess struct {
	Amount float64
}

type SuccessValue interface{}
