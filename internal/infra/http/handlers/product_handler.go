package product_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	product_entity "github.com/williamkoller/golang-domain-driven-design/internal/domain/product/entity"
	product_repository "github.com/williamkoller/golang-domain-driven-design/internal/domain/product/repository"
	shared_events "github.com/williamkoller/golang-domain-driven-design/internal/shared/domain/events"
)

type ProductHandler struct {
	repo       *product_repository.ProductRepository
	dispatcher *shared_events.EventDispatcher
}

func NewProductHandler(repo *product_repository.ProductRepository, dispatcher *shared_events.EventDispatcher) *ProductHandler {
	return &ProductHandler{repo, dispatcher}
}

// POST /products
func (h *ProductHandler) Create(c *gin.Context) {
	var input struct {
		Name       string   `json:"name"`
		Sku        int      `json:"sku"`
		Categories []string `json:"categories"`
		Price      int      `json:"price"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, _, err := product_entity.NewProduct(input.Name, input.Sku, input.Categories, input.Price, h.dispatcher)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Add(*product); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) FindAll(c *gin.Context) {
	products, _ := h.repo.Find()
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) FindOne(c *gin.Context) {
	name := c.Param("name")
	product, err := h.repo.FindOne(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}
