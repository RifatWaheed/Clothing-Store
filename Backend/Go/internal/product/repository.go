package product

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	DB *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{DB: db}
}

// GetProducts retrieves a paginated list of products with optional search filtering
func (r *Repository) GetProducts(ctx context.Context, limit, offset int, search string) ([]Product, int64, error) {
	// Build the query with optional search filter
	baseQuery := `
		SELECT 
			pkid, name, created_by, modified_by, created_date, modified_date,
			description, price, discount_amount, discount_percent,
			sku_id, sku_code, color_id, color_name, gender_id, gender_name,
			size_id, size_name, stock_id, stock_qty, type_id, type_name,
			voucher_id, voucher_code
		FROM products
		WHERE 1=1
	`

	args := []interface{}{}
	argIndex := 1

	// Apply search filter if provided
	if search != "" {
		baseQuery += ` AND (name ILIKE $` + string(rune(argIndex)) + ` OR description ILIKE $` + string(rune(argIndex)) + `)`
		args = append(args, "%"+search+"%")
		argIndex++
	}

	// Count total records with filter
	countQuery := `SELECT COUNT(*) FROM products WHERE 1=1`
	if search != "" {
		countQuery += ` AND (name ILIKE $1 OR description ILIKE $1)`
	}

	var total int64
	err := r.DB.QueryRow(ctx, countQuery, args[:len(args)-1]...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Apply ordering and pagination
	baseQuery += ` ORDER BY created_date DESC LIMIT $` + string(rune(argIndex)) + ` OFFSET $` + string(rune(argIndex+1))
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, baseQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(
			&p.PKID, &p.Name, &p.CreatedBy, &p.ModifiedBy, &p.CreatedAt, &p.ModifiedAt,
			&p.Description, &p.Price, &p.DiscountAmount, &p.DiscountPercent,
			&p.SKUId, &p.SKUCode, &p.ColorId, &p.ColorName, &p.GenderId, &p.GenderName,
			&p.SizeId, &p.SizeName, &p.StockId, &p.StockQty, &p.TypeId, &p.TypeName,
			&p.VoucherId, &p.VoucherCode,
		)
		if err != nil {
			return nil, 0, err
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

// GetProductByID retrieves a single product by PKID
func (r *Repository) GetProductByID(ctx context.Context, pkid int64) (*Product, error) {
	query := `
		SELECT 
			pkid, name, created_by, modified_by, created_date, modified_date,
			description, price, discount_amount, discount_percent,
			sku_id, sku_code, color_id, color_name, gender_id, gender_name,
			size_id, size_name, stock_id, stock_qty, type_id, type_name,
			voucher_id, voucher_code
		FROM products
		WHERE pkid = $1
	`

	var p Product
	err := r.DB.QueryRow(ctx, query, pkid).Scan(
		&p.PKID, &p.Name, &p.CreatedBy, &p.ModifiedBy, &p.CreatedAt, &p.ModifiedAt,
		&p.Description, &p.Price, &p.DiscountAmount, &p.DiscountPercent,
		&p.SKUId, &p.SKUCode, &p.ColorId, &p.ColorName, &p.GenderId, &p.GenderName,
		&p.SizeId, &p.SizeName, &p.StockId, &p.StockQty, &p.TypeId, &p.TypeName,
		&p.VoucherId, &p.VoucherCode,
	)

	if err != nil {
		return nil, err
	}

	return &p, nil
}
