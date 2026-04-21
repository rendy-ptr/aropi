import type { PublicMenu } from './../types/menu';
import apiClient from '@/utils/apiClient';
import type { ApiResponse } from '@/types/ApiResponse';
import { API_PATHS } from '@/constants/apiPaths';

export const getPublicMenus = async (
  search?: string,
  category?: string
): Promise<PublicMenu[]> => {
  const { data } = await apiClient.get<ApiResponse<PublicMenu[]>>(
    API_PATHS.PUBLIC.PRODUCTS(),
    {
      params: {
        search,
        category_id: category && category !== 'semua' ? category : undefined,
      },
    }
  );
  return data.data;
};
