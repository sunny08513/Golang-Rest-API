package product

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	m "Product/models"

	"github.com/go-redis/redis/v8"
)

type productRedis struct {
	rdb *redis.Client
	ctx context.Context
}

func NewRedisStore(rdb *redis.Client) *productRedis {
	return &productRedis{
		rdb: rdb,
		ctx: context.Background(),
	}
}

func (p *productRedis) GetProduct() ([]m.Product, error) {
	keys, err := p.rdb.Keys(p.ctx, "product:*").Result()
	if err != nil {
		return nil, err
	}

	products := []m.Product{}
	for _, key := range keys {
		prod, err := p.GetProductById(key)
		if err != nil {
			return nil, err
		}
		products = append(products, *prod)
	}
	return products, nil
}

func (p *productRedis) CreateProduct(product m.Product) (*m.Product, error) {
	id, err := p.rdb.Incr(p.ctx, "product:next-id").Result()
	if err != nil {
		return nil, err
	}

	product.Id = int(id)
	key := fmt.Sprintf("product:%d", product.Id)

	prodBytes, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}

	err = p.rdb.Set(p.ctx, key, prodBytes, 0).Err()
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *productRedis) GetProductById(id string) (*m.Product, error) {
	key := fmt.Sprintf("product:%s", id)
	prodBytes, err := p.rdb.Get(p.ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	var prod m.Product
	err = json.Unmarshal(prodBytes, &prod)
	if err != nil {
		return nil, err
	}

	return &prod, nil
}

func (p *productRedis) UpdateProduct(id int, product *m.Product) (*m.Product, error) {
	key := fmt.Sprintf("product:%d", id)
	prodBytes, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}

	err = p.rdb.Set(p.ctx, key, prodBytes, 0).Err()
	if err != nil {
		return nil, err
	}

	return p.GetProductById(strconv.Itoa(id))
}

func (p *productRedis) DeleteProduct(id int) (string, error) {
	key := fmt.Sprintf("product:%d", id)
	err := p.rdb.Del(p.ctx, key).Err()
	if err != nil {
		return "", err
	}

	message := fmt.Sprintf("Product with id %d deleted successfully.", id)
	return message, nil
}
