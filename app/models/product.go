package models

type Category struct {
	Id   *int32  `json:"id"`
	Name *string `json:"name"`
}

type Package struct {
	Id       int32  `json:"id"`
	Type     string `json:"type"`
	Material string `json:"material"`
	Weight   int32  `json:"weight"`
	Length   int32  `json:"length"`
	Width    int32  `json:"width"`
	Height   int32  `json:"height"`
	Price    int32  `json:"price"`
}

type Property struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Measure string `json:"measure"`
}

type Product struct {
	Id          int32                  `json:"id" db:"p_id"`
	Name        string                 `json:"name" db:"p_name"`
	Description string                 `json:"description" db:"p_description"`
	Brand       string                 `json:"brand" db:"p_brand"`
	Price       int32                  `json:"price" db:"p_price"`
	Image       string                 `json:"image" db:"p_image"`
	Sku         string                 `json:"sku" db:"p_sku"`
	Quantity    int32                  `json:"quantity" db:"p_quantity"`
	Barcode     string                 `json:"barcode" db:"p_barcode"`
	Category    Category               `json:"category"`
	Package     Package                `json:"package"`
	Set         []*string              `json:"set"`
	Gallery     []*string              `json:"gallery"`
	Property    []*map[string]Property `json:"property"`
}

type ProductRepository interface {
	FindAll(limit string, offset string, category string) ([]*Product, error)
	FindOne(productId int32) (*Product, error)
}
