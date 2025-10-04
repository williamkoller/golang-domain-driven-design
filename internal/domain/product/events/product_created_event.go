package product_events

import (
	"time"
)

type ProductCreatedEvent struct {
	Name       string
	Sku        int
	Categories []string
	Price      int
	Timestamp  time.Time
}

func NewProductCreatedEvent(name string, sku int, categories []string, price int) *ProductCreatedEvent {
	return &ProductCreatedEvent{Name: name, Sku: sku, Categories: categories, Price: price, Timestamp: time.Now()}
}

func (e *ProductCreatedEvent) EventName() string {
	return "product.created"
}
