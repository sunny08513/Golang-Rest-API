package product

import (
	m "Product/models"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type product struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *product {
	return &product{
		db: db,
	}
}

func (p *product) GetProduct() ([]m.Product, error) {
	if err := p.db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to MySQL database!")

	rows, err := p.db.Query("SELECT id, name, price FROM products;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []m.Product{}
	for rows.Next() {
		p := m.Product{}

		if err := rows.Scan(&p.Id, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (p *product) CreateProduct(product m.Product) (*m.Product, error) {
	insertProduct := `INSERT INTO products (name, price) VALUES (?, ?);`
	stmt, err := p.db.Prepare(insertProduct)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the prepared statement with the variables
	res, err := stmt.Exec(product.Name, product.Price)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	fmt.Printf("User inserted successfully. Rows affected: %d\n", rowsAffected)

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return p.GetProductById(int(lastInsertID))
}

func (p *product) GetProductById(id int) (*m.Product, error) {
	query := "SELECT id, name, price FROM products where id = ?;"
	row := p.db.QueryRow(query, id)
	prod := m.Product{}

	if err := row.Scan(&prod.Id, &prod.Name, &prod.Price); err != nil {
		return nil, err
	}

	return &prod, nil
}

func (p *product) UpdateProduct(id int, product *m.Product) (*m.Product, error) {
	updateProduct := `UPDATE products set price = ? WHERE id = ?;`
	stmt, err := p.db.Prepare(updateProduct)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute the prepared statement with the variables
	res, err := stmt.Exec(product.Price, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	fmt.Printf("User inserted successfully. Rows affected: %d\n", rowsAffected)
	return p.GetProductById(id)
}

func (p *product) DeleteProduct(id int) (string, error) {
	deleteUser := `DELETE FROM products WHERE id = ?;`
	stmt, err := p.db.Prepare(deleteUser)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return "", err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return "", err
	}
	fmt.Printf("User inserted successfully. Rows affected: %d\n", rowsAffected)

	message := fmt.Sprintf("Product with id  %d , deleted successfully.", id)
	return message, nil
}
