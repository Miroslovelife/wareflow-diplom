// src/hooks/useCheckAuth.ts
import { useState, useEffect } from 'react';

export const useCheckAuth = () => {
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [loading, setLoading] = useState(true); // Новое состояние для отслеживания загрузки

    useEffect(() => {
        const accessToken = localStorage.getItem('accessToken');
        if (accessToken) {
            setIsAuthenticated(true);
        } else {
            setIsAuthenticated(false);
        }
        setLoading(false); // Завершаем загрузку
    }, []);

    return { isAuthenticated, loading };
};