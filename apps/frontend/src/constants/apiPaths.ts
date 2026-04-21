export const API_PATHS = {
  AUTH: {
    LOGIN: '/public/users/login',
    REGISTER: '/public/users/register',
    ME: '/protected/users/me',
    LOGOUT: '/public/users/logout',
  },
  ADMIN: {
    DASHBOARD: '/dashboard/admin',
    SETTING: (id?: string) => (id ? `/admin/setting/${id}` : `/admin/setting`),
    KATEGORI: (id?: string) =>
      id ? `/admin/kategori/${id}` : `/admin/kategori`,
    MENU: (id?: string) => (id ? `/admin/menu/${id}` : `/admin/menu`),
    KASIR: (id?: string) => (id ? `/admin/kasir/${id}` : `/admin/kasir`),
    TABLE: (id?: string) => (id ? `/admin/table/${id}` : `/admin/table`),
    REWARD: (id?: string) => (id ? `/admin/reward/${id}` : `/admin/reward`),
  },
  KASIR: {
    DASHBOARD: '/dashboard/kasir',
    MENU: (id?: string) => (id ? `/kasir/menu/${id}` : `/kasir/menu`),
    CATEGORY: (id?: string) =>
      id ? `/kasir/category/${id}` : `/kasir/category`,
    TABLE: (id?: string) => (id ? `/kasir/table/${id}` : `/kasir/table`),
    MEMBER_ID: (id?: string) => (id ? `/kasir/member/${id}` : `/kasir/member`),
    ORDER: (id?: string) => (id ? `/kasir/order/${id}` : `/kasir/order`),
    ORDER_DETAIL: (id: string) => `/kasir/order/${id}/detail`,
  },
  PUBLIC: {
    UPLOAD_IMAGE: '/upload',
    PRODUCTS: () => '/public/products',
    CATEGORIES: () => '/public/categories',
  },
};
