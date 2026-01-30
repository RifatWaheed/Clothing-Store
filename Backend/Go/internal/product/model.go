package product

import (
	"database/sql"
	"time"
)

type Product struct {
	// Base fields
	PKID       int64         `json:"pkid"`
	Name       string        `json:"name"`
	CreatedBy  int64         `json:"created_by"`
	ModifiedBy sql.NullInt64 `json:"modified_by,omitempty"`
	CreatedAt  time.Time     `json:"created_at"`
	ModifiedAt sql.NullTime  `json:"modified_at,omitempty"`

	// Product description
	Description sql.NullString `json:"description,omitempty"`

	// Pricing
	Price           float64 `json:"price"`
	DiscountAmount  float64 `json:"discount_amount"`
	DiscountPercent float64 `json:"discount_percent"`

	// SKU
	SKUId   sql.NullInt64  `json:"sku_id,omitempty"`
	SKUCode sql.NullString `json:"sku_code,omitempty"`

	// Color
	ColorId   sql.NullInt64  `json:"color_id,omitempty"`
	ColorName sql.NullString `json:"color_name,omitempty"`

	// Gender
	GenderId   sql.NullInt64  `json:"gender_id,omitempty"`
	GenderName sql.NullString `json:"gender_name,omitempty"`

	// Size
	SizeId   sql.NullInt64  `json:"size_id,omitempty"`
	SizeName sql.NullString `json:"size_name,omitempty"`

	// Stock
	StockId  sql.NullInt64 `json:"stock_id,omitempty"`
	StockQty int32         `json:"stock_qty"`

	// Product type
	TypeId   sql.NullInt64  `json:"type_id,omitempty"`
	TypeName sql.NullString `json:"type_name,omitempty"`

	// Voucher
	VoucherId   sql.NullInt64  `json:"voucher_id,omitempty"`
	VoucherCode sql.NullString `json:"voucher_code,omitempty"`
}

// ListParams defines optional query parameters for listing products
type ListParams struct {
	Limit  int    `json:"limit" form:"limit" binding:"omitempty,min=1,max=100"`
	Offset int    `json:"offset" form:"offset" binding:"omitempty,min=0"`
	Search string `json:"search" form:"search" binding:"omitempty,max=255"`
}

// ProductResponse wraps products with pagination metadata
type ProductResponse struct {
	Data   []Product `json:"data"`
	Total  int64     `json:"total"`
	Limit  int       `json:"limit"`
	Offset int       `json:"offset"`
}
