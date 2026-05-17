package product

import "time"

type Product struct {
	PKID       int64      `json:"pkid"`
	Name       string     `json:"name"`
	CreatedBy  int64      `json:"created_by"`
	ModifiedBy *int64     `json:"modified_by,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`

	Description *string `json:"description,omitempty"`

	Price           float64 `json:"price"`
	DiscountAmount  float64 `json:"discount_amount"`
	DiscountPercent float64 `json:"discount_percent"`

	SKUId   *int64  `json:"sku_id,omitempty"`
	SKUCode *string `json:"sku_code,omitempty"`

	ColorId   *int64  `json:"color_id,omitempty"`
	ColorName *string `json:"color_name,omitempty"`

	GenderId   *int64  `json:"gender_id,omitempty"`
	GenderName *string `json:"gender_name,omitempty"`

	SizeId   *int64  `json:"size_id,omitempty"`
	SizeName *string `json:"size_name,omitempty"`

	StockId  *int64 `json:"stock_id,omitempty"`
	StockQty int32  `json:"stock_qty"`

	TypeId   *int64  `json:"type_id,omitempty"`
	TypeName *string `json:"type_name,omitempty"`

	VoucherId   *int64  `json:"voucher_id,omitempty"`
	VoucherCode *string `json:"voucher_code,omitempty"`
}

type ListParams struct {
	Limit  int    `json:"limit" form:"limit" binding:"omitempty,min=1,max=100"`
	Offset int    `json:"offset" form:"offset" binding:"omitempty,min=0"`
	Search string `json:"search" form:"search" binding:"omitempty,max=255"`
}

type ProductResponse struct {
	Data   []Product `json:"data"`
	Total  int64     `json:"total"`
	Limit  int       `json:"limit"`
	Offset int       `json:"offset"`
}
