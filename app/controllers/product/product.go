package product

import (
	"demo/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type HandlerProduct struct {
	productRepo models.ProductRepository
}

func NewHandlerProduct(productRepo models.ProductRepository) *HandlerProduct {
	return &HandlerProduct{
		productRepo: productRepo,
	}
}

// List @BasePath /api
// @Tags Product
// @Summary Product list
// @Produce json
// @Accept json
// @Schemes
// @Description Получение списка продуктов с возможным применением фильтров
// @Param category query int false "category id" Enums(1,2)
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Success 200 {array} models.Product "Successful operation"
// @Failure 404 {object} models.Error "Unexpected error"
// @Router /product [get]
func (hp *HandlerProduct) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := r.URL.Query()
		products, _ := hp.productRepo.FindAll(
			vars.Get("limit"),
			vars.Get("offset"),
			vars.Get("category"),
		)

		byteData, _ := json.Marshal(products)
		w.Write(byteData)
	}
}

// Detail @BasePath /api
// @Tags Product
// @Summary Product detail
// @Schemes
// @Description Получение детальной информации по товару
// @Produce json
// @Accept json
// @Success 200 {object} models.Product "Successful operation"
// @Failure 404 {object} models.Error "Unexpected error"
// @Router /product/{id} [get]
func (hp *HandlerProduct) Detail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)

		productId, err := strconv.Atoi(vars["id"])
		if err != nil || productId < 0 {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		product, error := hp.productRepo.FindOne(int32(productId))
		if error != nil {
			http.Error(w, "Product not found, detail: "+error.Error(), http.StatusNotFound)
			return
		}

		byteData, _ := json.Marshal(product)
		w.Write(byteData)
	}
}
