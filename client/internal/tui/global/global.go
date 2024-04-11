package global

import "github.com/ElrohirGT/Ratatouille/internal/db"

var (
	Id       int
	Username string
	Role     int

	Driver *db.Queries
)

type ErrorDB struct {
	Description string
}
type SuccesDB struct {
	Description string
}
