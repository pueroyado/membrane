package models

type Category struct {
	Id   *int32  `json:"id"`
	Name *string `json:"name"`
}

type Property struct {
	Barcode *string `json:"barcode"`
	Weight  *int32  `json:"weight"`
	Height  *int32  `json:"height"`
	Color   *string `json:"color"`
	Vat     *string `json:"vat"`
}

type Product struct {
	Id          int32    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Brand       string   `json:"brand"`
	Preview     string   `json:"preview"`
	Price       int32    `json:"price"`
	Category    Category `json:"category"`
	Property    Property `json:"property"`
}

type ProductRepository interface {
	FindAll(limit string, offset string, category string) ([]*Product, error)
	FindOne(productId int32) (*Product, error)
}
