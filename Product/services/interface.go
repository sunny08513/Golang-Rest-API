package services

import (
	m "Golang-Rest-API/product/models"
	"context"
)

type Product interface {
	GetProduct(ctx context.Context) ([]m.Product, error)
	CreateProduct(product m.Product) (*m.Product, error)
	GetProductById(id int) (*m.Product, error)
	UpdateProduct(id int, product *m.Product) (*m.Product, error)
	DeleteProduct(id int) (string, error)
}
