package product

import (
	m "Golang-Rest-API/product/models"
	productStore "Golang-Rest-API/product/stores"
	"context"
	"errors"
)

type product struct {
	pStore productStore.Product
}

func NewProduct(pStore productStore.Product) *product {
	return &product{
		pStore: pStore,
	}
}

func (p *product) GetProduct(ctx context.Context) ([]m.Product, error) {
	// products := []m.Product{
	// 	{
	// 		Id:    1,
	// 		Name:  "Milk",
	// 		Price: 26.5,
	// 	},
	// 	{
	// 		Id:    1,
	// 		Name:  "Grapes",
	// 		Price: 70,
	// 	},
	// }

	products, err := p.pStore.GetProduct(ctx)
	return products, err
}

func (p *product) CreateProduct(product m.Product) (*m.Product, error) {
	return p.pStore.CreateProduct(product)
}

func (p *product) GetProductById(id int) (*m.Product, error) {
	return p.pStore.GetProductById(id)
}

func (p *product) UpdateProduct(id int, product *m.Product) (*m.Product, error) {
	_, err := p.GetProductById(id)
	if err != nil {
		return nil, errors.New("product doesn't exist")
	}
	return p.pStore.UpdateProduct(id, product)
}

func (p *product) DeleteProduct(id int) (string, error) {
	_, err := p.GetProductById(id)
	if err != nil {
		return "", errors.New("product doesn't exist")
	}
	return p.pStore.DeleteProduct(id)
}
