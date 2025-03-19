import axios from 'axios';

export const api = axios.create({
    baseURL: 'https://bebradomen.twc1.net:8443/',
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
