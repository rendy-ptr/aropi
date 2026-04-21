export interface PublicMenu {
  id: string;
  product_image_file: string;
  name: string;
  price: number;
  stock: number;
  category: {
    id: string;
    name: string;
  };
}
