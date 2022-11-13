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

func (r *ProductRepo) FindAll(
	limit string,
	offset string,
	category string,
) ([]*models.Product, error) {
	//var queryParams []string
	rowSql := r.getBaseSelect()

	if category != "" {
		rowSql += "WHERE p_category_id = " + category
		// queryParams = append(queryParams, category)
	}
	if limit != "" && offset != "" {
		rowSql += "LIMIT " + offset + "," + limit + " "
		// queryParams = append(queryParams, offset, limit)
	}
	rows, err := r.db.Query(rowSql)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	var products []*models.Product
	var ids []int
	for rows.Next() {
		p := models.Product{}
		err := rows.Scan(
			&p.Id, &p.Name, &p.Description, &p.Brand, &p.Image, &p.Price, &p.Quantity, &p.Sku, &p.Barcode,
			&p.Category.Id,
			&p.Category.Name,
			&p.Package.Id, &p.Package.Type, &p.Package.Material, &p.Package.Weight, &p.Package.Length, &p.Package.Width,
			&p.Package.Height, &p.Package.Price,
		)
		if err != nil {
			log.Fatalln(err)
		}
		ids = append(ids, int(p.Id))
		products = append(products, &p)
	}

	sets := r.getSet(ids)
	properties := r.getProperties(ids)
	gallery := r.getGalleryImage(ids)

	for _, product := range products {
		product.Set = sets[product.Id]
		product.Property = properties[product.Id]
		product.Gallery = gallery[product.Id]
	}

	return products, nil
}

func (r *ProductRepo) FindOne(productId int32) (*models.Product, error) {
	rowSql := r.getBaseSelect()
	rowSql += "WHERE p_id = ? " +
		"ORDER BY p_id "

	p := &models.Product{}
	row := r.db.QueryRowx(rowSql, productId)
	err := row.StructScan(p)
	if err != nil {
		return nil, err
	}

	ids := []int{int(p.Id)}
	sets := r.getSet(ids)
	p.Set = sets[productId]
	if p.Set == nil {
		p.Set = make([]*string, 0)
	}

	galleryImage := r.getGalleryImage(ids)
	galleryImageProduct := galleryImage[productId]
	if galleryImageProduct == nil {
		p.Gallery = make([]*string, 0)
	} else {
		p.Gallery = galleryImageProduct
	}

	properties := r.getProperties(ids)
	propertiesProductId := properties[productId]
	if propertiesProductId == nil {
		p.Property = make([]*map[string]models.Property, 0)
	} else {
		p.Property = properties[productId]
	}

	return p, nil
}

func (r *ProductRepo) getBaseSelect() string {
	sqlBase := "SELECT " +
		"p_id, p_name, p_description, p_brand, p_image, p_price, p_quantity, p_sku, p_barcode, " +
		"p_category_id AS `category.id`, p_category_name AS `category.name`, " +
		"p_package_id AS `package.id`, " +
		"p_package_type AS `package.type`, " +
		"p_package_material AS `package.material`, " +
		"p_package_weight AS `package.weight`, " +
		"p_package_length AS `package.length`, " +
		"p_package_width AS `package.width`, " +
		"p_package_height AS `package.height`, " +
		"p_package_price AS `package.price` " +
		"FROM p_product " +
		"LEFT JOIN p_category ON p_category_id = p_category " +
		"LEFT JOIN p_package ON p_package_id = p_package "

	return sqlBase
}

func (r *ProductRepo) getGalleryImage(ids []int) map[int32][]*string {
	sqlGallery := "SELECT " +
		"p_gallery_image_link, p_gallery_image_product_id " +
		"FROM p_gallery_image " +
		"WHERE p_gallery_image_product_id IN (?) " +
		"order by p_gallery_image_position ASC "

	q, args, _ := sqlx.In(sqlGallery, ids)
	rows, err := r.db.Queryx(q, args...)
	if err != nil {
		log.Println("error get galleryImage query", err.Error())
	}
	defer rows.Close()

	gallery := make(map[int32][]*string)
	for rows.Next() {
		var productId int32
		var imageUrl string
		rows.Scan(&imageUrl, &productId)
		gallery[productId] = append(gallery[productId], &imageUrl)
	}
	return gallery
}

func (r *ProductRepo) getSet(ids []int) map[int32][]*string {
	sqlSet := "SELECT " +
		"p_set_name, p_set_product_product_id " +
		"FROM p_set_product " +
		"LEFT JOIN p_set ps ON ps.p_set_id = p_set_product_set_id " +
		"WHERE " +
		"p_set_product_product_id IN (?) " +
		"AND p_set_status = 1"

	q, args, _ := sqlx.In(sqlSet, ids)
	rows, err := r.db.Queryx(q, args...)
	if err != nil {
		log.Println("error get set query", err.Error())
	}
	defer rows.Close()

	sets := make(map[int32][]*string)
	for rows.Next() {
		var productId int32
		var setName string
		rows.Scan(&setName, &productId)
		sets[productId] = append(sets[productId], &setName)
	}
	return sets
}

func (r *ProductRepo) getProperties(ids []int) map[int32][]*map[string]models.Property {
	sqlProperty := "SELECT " +
		"p_property_product_product_id, " +
		"p_property_code, p_property_name, p_property_value_value, p_property_measure " +
		"FROM p_property_product " +
		"LEFT JOIN p_property ON p_property_id = p_property_product_property_id " +
		"LEFT JOIN p_property_value ON p_property_value_property_product_id = p_property_product_id " +
		"WHERE p_property_product_product_id IN (?);"

	q, args, _ := sqlx.In(sqlProperty, ids)
	rows, err := r.db.Queryx(q, args...)
	if err != nil {
		log.Println("error get property query", err)
	}
	defer rows.Close()

	props := make(map[int32][]*map[string]models.Property)
	for rows.Next() {
		var productId int32
		var productFieldCode, name, value, measure string
		rows.Scan(&productId, &productFieldCode, &name, &value, &measure)

		prop := make(map[string]models.Property)
		prop[productFieldCode] = models.Property{
			Name:    name,
			Value:   value,
			Measure: measure,
		}
		props[productId] = append(props[productId], &prop)
	}
	return props
}
