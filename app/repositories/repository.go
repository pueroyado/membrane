package repositories

import (
	"demo/models"
	"github.com/jmoiron/sqlx"
	"log"
)

type ProductRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) Search(
	limit string,
	offset string,
	category string,
) ([]*models.Product, error) {
	rowSql := "select " +
		"p_id, p_name, p_description, p_brand, p_preview, p_price, " +
		"p_cat_id, p_cat_name, " +
		"p_prop_barcode, " +
		"p_prop_weight, " +
		"p_prop_height, " +
		"p_prop_color, " +
		"p_prop_vat " +
		"from product " +
		"LEFT JOIN product_category ON p_cat_id = p_category " +
		"LEFT JOIN product_property ON p_prop_product_id = p_id "

	if category != "" {
		rowSql += "WHERE p_category = " + category + " "
	}
	if limit != "" && offset != "" {
		rowSql += "LIMIT " + offset + "," + limit + " "
	}
	rows, err := r.db.Query(rowSql)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		p := models.Product{}
		err := rows.Scan(
			&p.Id,
			&p.Name,
			&p.Description,
			&p.Brand,
			&p.Preview,
			&p.Price,
			&p.Category.Id,
			&p.Category.Name,
			&p.Property.Barcode,
			&p.Property.Weight,
			&p.Property.Height,
			&p.Property.Color,
			&p.Property.Vat,
		)
		if err != nil {
			log.Fatalln(err)
		}
		products = append(products, &p)
	}
	return products, nil
}
