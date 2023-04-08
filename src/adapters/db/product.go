package db

import (
	"database/sql"
	"log"

	"github.com/Lukiita/go-hexagonal/src/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db}
}

func (p *ProductDb) Get(id string) (application.IProduct, error) {
	var product application.Product
	stmt, err := p.db.Prepare("select id, name, price, status from products where id=?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDb) Save(product application.IProduct) (application.IProduct, error) {
	var rows int
	p.db.QueryRow("SELECT id FROM products WHERE id = ?", product.GetID()).Scan(&rows)
	if rows == 0 {
		_, err := p.create(product)
		if err != nil {
			log.Fatalf(err.Error())
			return nil, err
		}
	} else {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}

	return product, nil
}

func (p *ProductDb) create(product application.IProduct) (application.IProduct, error) {

	stmt, err := p.db.Prepare(`INSERT INTO products(id, name, price, status) VALUES (?, ?, ?, ?)`)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}

	_, err = stmt.Exec(
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)

	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product application.IProduct) (application.IProduct, error) {
	_, err := p.db.Exec(
		"UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?",
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
		product.GetStatus(),
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}
