package db

import (
	"database/sql"
	"github.com/MarceloAvanzi/appproduct-hexagonal-arch/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (s *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Preprare("select id, name, price, status from product where id=?")
	if err != nil {
		return nil, err
	}
	
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}

	return &product, nill
}