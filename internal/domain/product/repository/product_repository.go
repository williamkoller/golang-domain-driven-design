package product_repository

import (
	"errors"
	"sync"

	product_entity "github.com/williamkoller/golang-domain-driven-design/internal/domain/product/entity"
)

type IProductRepository interface {
	Add(product product_entity.Product) error
	Find() ([]product_entity.Product, error)
	FindOne(name string) (product_entity.Product, error)
}

type ProductRepository struct {
	data map[string]product_entity.Product
	mu   sync.RWMutex
}

func NewRepository() *ProductRepository {
	return &ProductRepository{
		data: make(map[string]product_entity.Product),
	}
}

func (r *ProductRepository) Add(product product_entity.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.data[product.Name]; exists {
		return errors.New("product already exists")
	}

	r.data[product.Name] = product

	return nil
}

func (r *ProductRepository) Find() ([]product_entity.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	products := make([]product_entity.Product, 0, len(r.data))

	for _, p := range r.data {
		products = append(products, p)
	}

	return products, nil
}

func (r *ProductRepository) FindOne(name string) (product_entity.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	product, exists := r.data[name]

	if !exists {
		return product_entity.Product{}, errors.New("product not found")
	}

	return product, nil
}
