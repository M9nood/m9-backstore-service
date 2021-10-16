package product

import (
	"fmt"
	"time"

	"github.com/guregu/null"
)

type ProductSchema struct {
	Id          int        `gorm:"column:id" json:"id"`
	ProductName string     `gorm:"column:product_name" json:"product_name"`
	Description string     `gorm:"column:description" json:"description"`
	DispCode    string     `gorm:"column:disp_code" json:"disp_code"`
	InStock     int        `gorm:"column:in_stock" json:"in_stock"`
	ProductUuid *string    `gorm:"-;column:product_uuid" json:"product_uuid,omitempty"`
	StoreId     int        `gorm:"column:store_id" json:"store_id"`
	ImageKey    string     `gorm:"column:image_key" json:"image_key"`
	Cost        null.Float `gorm:"column:cost" json:"cost"`
	Price       null.Float `gorm:"column:price" json:"price"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeleteFlag  int        `gorm:"column:delete_flag" json:"delete_flag"`
}

type Products []ProductSchema

func ToLineMessage(p ProductSchema) string {
	return fmt.Sprintf("%s   x%d\n", p.ProductName, p.InStock)
}

func (p ProductSchema) Response() ProductResponse {
	return ProductResponse{
		Id:          p.Id,
		ProductName: p.ProductName,
		Description: p.Description,
		DispCode:    p.DispCode,
		InStock:     p.InStock,
		ProductUuid: p.ProductUuid,
		StoreId:     p.StoreId,
		ImageKey:    p.ImageKey,
		Cost:        p.Cost,
		Price:       p.Price,
	}
}

func (ps Products) Response() []ProductResponse {
	res := []ProductResponse{}
	list := []ProductSchema(ps)

	for _, item := range list {
		res = append(res, item.Response())
	}
	return res
}
