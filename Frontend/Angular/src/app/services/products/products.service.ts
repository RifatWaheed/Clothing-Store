import { Injectable, inject } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Product, ProductListResponse } from '../../models/products/products';

@Injectable({
  providedIn: 'root',
})
export class ProductService {
  private http = inject(HttpClient);

  getProducts(params: { limit?: number; offset?: number; search?: string } = {}) {
    const { limit = 20, offset = 0, search = '' } = params;
    let httpParams = new HttpParams()
      .set('limit', limit)
      .set('offset', offset);
    if (search) httpParams = httpParams.set('search', search);
    return this.http.get<ProductListResponse>('/api/public/products', { params: httpParams });
  }

  getProduct(id: number) {
    return this.http.get<Product>(`/api/public/products/${id}`);
  }
}
