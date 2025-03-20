import axios from 'axios';
import { useAuth } from '../contexts/AuthProvider';
export const api = axios.create({
    baseURL: 'http://localhost:8089/',
    withCredentials: true,
});

// Добавляем interceptor для автоматической подстановки токена
api.interceptors.request.use(
    (config) => {
        const accessToken = localStorage.getItem('accessToken');

        // Исключаем login, register, refresh
        if (
            accessToken &&
            !config.url?.includes('/auth/sign-in-email') &&
            !config.url?.includes('/auth/sign-in-phone') &&
            !config.url?.includes('/auth/register') &&
            !config.url?.includes('/auth/refresh')
        ) {
            config.headers.Authorization = `Bearer ${accessToken}`;
        }

        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

api.interceptors.response.use(
    (response) => response,
    async (error) => {
        const originalRequest = error.config;

        if (error.response?.status === 401 && !originalRequest._retry) {
            originalRequest._retry = true;
            try {
                const refreshResponse = await axios.get('/api/v1/auth/refresh', { withCredentials: true });
                if (refreshResponse.data.accessToken) {
                    localStorage.setItem('accessToken', refreshResponse.data.accessToken);
                    originalRequest.headers.Authorization = `Bearer ${refreshResponse.data.accessToken}`;
                    return api(originalRequest);
                }
            } catch (refreshError) {
                const { logout } = useAuth(); // Выход из системы, если рефреш не сработал
                logout();
                return Promise.reject(refreshError);
            }
        }
        return Promise.reject(error);
    }
);