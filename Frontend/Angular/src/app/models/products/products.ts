export interface Product {
  pkid: number;
  name: string;
  created_by: number;
  modified_by?: number | null;
  created_at: string;
  modified_at?: string | null;
  description?: string | null;
  price: number;
  discount_amount: number;
  discount_percent: number;
  sku_id?: number | null;
  sku_code?: string | null;
  color_id?: number | null;
  color_name?: string | null;
  gender_id?: number | null;
  gender_name?: string | null;
  size_id?: number | null;
  size_name?: string | null;
  stock_id?: number | null;
  stock_qty: number;
  type_id?: number | null;
  type_name?: string | null;
  voucher_id?: number | null;
  voucher_code?: string | null;
}

export interface ProductListResponse {
  data: Product[];
  total: number;
  limit: number;
  offset: number;
}
