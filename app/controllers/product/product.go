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
