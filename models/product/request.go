package product

import "github.com/guregu/null"

type ProductRequest struct {
	Id          int        `json:"id,omitempty"`
	ProductName string     `json:"product_name"  validate:"required"`
	Description string     `json:"description"`
	DispCode    string     `json:"disp_code"  validate:"required"`
	InStock     int        `json:"in_stock"`
	ImageKey    string     `json:"image_key"`
	Cost        null.Float `json:"cost"`
	Price       null.Float `json:"price"`
}

type ProductQueryParams struct {
	Page     *int    `query:"page" json:"page,omitempty"`
	PageSize *int    `query:"pageSize" json:"pageSize,omitempty"`
	Q        *string `query:"q" json:"q,omitempty"`
	Count    bool
}

func (p ProductRequest) ToProductSchema() ProductSchema {
	return ProductSchema{
		Id:          p.Id,
		ProductName: p.ProductName,
		Description: p.Description,
		DispCode:    p.DispCode,
		InStock:     p.InStock,
		ImageKey:    p.ImageKey,
		Cost:        p.Cost,
		Price:       p.Price,
	}
}
