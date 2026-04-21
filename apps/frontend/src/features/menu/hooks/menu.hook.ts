import { useQuery } from '@tanstack/react-query';
import { getPublicMenus } from '../services/menu.service';
import type { PublicMenu } from '../types/menu';

export const usePublicMenus = (search?: string, category?: string) => {
  return useQuery<PublicMenu[], Error>({
    queryKey: ['publicMenus', search, category],
    queryFn: () => getPublicMenus(search, category),
    staleTime: 5000,
    refetchOnWindowFocus: false,
  });
};
