import { Base } from '../base/base';

export class Products extends Base {
  colorID: number = 0;
  colorName: string = '';
  description: string = '';
  price: number = 0;
  discountAmount : number = 0;
  discountPercent : number = 0;
  skuID: number = 0;
  skuCode: string = '';
  genderName: string = '';
  genderID: number = 0;
  sizeName: number = 0;
  sizeID: number = 0;
  stockQty: number = 0;
  stockID: number = 0;
  typeID: number = 0;
  typeName: string = '';
  voucherCode: string = '';
  voucherID: number = 0;
}
