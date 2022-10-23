package product

import (
	"demo/models"
	"encoding/json"
	"fmt"
	"net/http"
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
		products, _ := hp.productRepo.Search(
			vars.Get("limit"),
			vars.Get("offset"),
			vars.Get("category"),
		)

		byteData, _ := json.Marshal(products)
		w.Write(byteData)
	}
}
func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		name := r.URL.Query().Get("name")
		fmt.Fprintln(w, "Hello guy!", name)
	}
}
func Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		name := r.URL.Query().Get("name")
		fmt.Fprintln(w, "Hello guy!", name)
	}
}
func Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		name := r.URL.Query().Get("name")
		fmt.Fprintln(w, "Hello guy!", name)
	}
}
