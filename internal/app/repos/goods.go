package repos

import (
	"database/sql"
)

type Goods struct {
	db *sql.DB
}

func NewGoods(db *sql.DB) *Goods {
	return &Goods{db}
}
