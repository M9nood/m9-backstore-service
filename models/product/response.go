package product

import "github.com/guregu/null"

type ProductResponse struct {
	Id          int        `json:"id"`
	ProductName string     `json:"product_name"`
	Description string     `json:"description"`
	DispCode    string     `json:"disp_code"`
	InStock     int        `json:"in_stock"`
	ProductUuid *string    `json:"product_uuid,omitempty"`
	StoreId     int        `json:"store_id"`
	ImageKey    string     `json:"image_key"`
	Cost        null.Float `json:"cost"`
	Price       null.Float `json:"price"`
}

type ProductsWithCountResponse struct {
	Products  []ProductResponse `json:"products"`
	TotalRows uint              `json:"totalRows"`
}
