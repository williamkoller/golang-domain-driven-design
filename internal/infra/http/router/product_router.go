package product_router

import (
	"github.com/gin-gonic/gin"
	product_handlers "github.com/williamkoller/golang-domain-driven-design/internal/infra/http/handlers"
)

func SetupProductRouter(productHandler *product_handlers.ProductHandler) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/products", productHandler.Create)
		v1.GET("/products", productHandler.FindAll)
		v1.GET("/products/:name", productHandler.FindOne)
	}

	return r
}
