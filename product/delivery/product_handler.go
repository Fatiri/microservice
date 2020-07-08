package delivery

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/microservice/product/models"
	"github.com/microservice/product/service"
	"github.com/sirupsen/logrus"
)

//ProductHandler struct
type ProductHandler struct {
	ProductService service.Service
}

//NewProductHandler route endpoint with mux
func NewProductHandler(r *mux.Router, productService service.Service) {
	claimHandler := &ProductHandler{
		ProductService: productService,
	}

	v1 := r.PathPrefix("/v1/product").Subrouter()

	v1.Handle("", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(claimHandler.createDataProduct))).Methods(http.MethodPost)
}

//createDataProduct function create data product
func (c *ProductHandler) createDataProduct(w http.ResponseWriter, r *http.Request) {
	product := new(models.Product)

	productName := r.FormValue("product_name")
	price := r.FormValue("price")

	product.ProductName = productName
	product.Price = price

	response, err := c.ProductService.CreateNewProduct(product)
	if err != nil {
		logrus.Error(err)
		return
	}

	fmt.Println(response)
	return
}
