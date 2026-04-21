import apiClient from '@/utils/apiClient';
import { API_PATHS } from '@/constants/apiPaths';

export interface AuthMeResponse {
  success: boolean;
  data: {
    id: string;
    email: string;
    role: 'MEMBER' | 'CASHIER' | 'ADMIN';
  };
}

export const getMe = async (): Promise<AuthMeResponse> => {
  const res = await apiClient.get<AuthMeResponse>(API_PATHS.AUTH.ME);
  return res.data;
};

export interface LoginResponse {
  success: boolean;
  message: string;
  data: {
    role: string;
  };
}

export const login = async (payload: { email: string; password: string }) => {
  const res = await apiClient.post<LoginResponse>(
    API_PATHS.AUTH.LOGIN,
    payload
  );
  return res.data;
};
