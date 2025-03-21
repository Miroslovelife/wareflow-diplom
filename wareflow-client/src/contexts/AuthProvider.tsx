import { createContext, useContext, useEffect, useState } from 'react';
import { api } from '../utils/api';

interface AuthContextType {
    isAuthenticated: boolean;
    role: 'admin' | 'owner' | 'employer' | null;
    username: string | null;
    isLoading: boolean;
    permissions: string[];
    systemPermissions: string[];
    loginWithEmail: (email: string, password: string) => Promise<void>;
    loginWithPhone: (phone_number: string, password: string) => Promise<void>;
    logout: () => void;
    getPermissionsForWarehouse: (warehouseId: string | undefined, username: string | null) => Promise<void>;
    getSystemPermissions: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const useAuth = () => {
    const context = useContext(AuthContext);
    if (!context) {
        throw new Error('useAuth must be used within an AuthProvider');
    }
    return context;
};

// Функция парсинга JWT
const parseJwt = (token: string): { role?: string, username?: string, exp?: number } | null => {
    try {
        const base64Url = token.split('.')[1];
        const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
        const jsonPayload = decodeURIComponent(
            atob(base64)
                .split('')
                .map((c) => `%${('00' + c.charCodeAt(0).toString(16)).slice(-2)}`)
                .join('')
        );
        return JSON.parse(jsonPayload);
    } catch (error) {
        return null;
    }
};

// Проверка, истек ли токен
const isTokenExpired = (token: string | null): boolean => {
    if (!token) return true;
    const decoded = parseJwt(token);
    return !decoded?.exp || decoded.exp * 1000 < Date.now();
};

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [role, setRole] = useState<'admin' | 'owner' | 'employer' | null>(null);
    const [username, setUsername] = useState<string | null>(null);
    const [isLoading, setIsLoading] = useState(true);
    const [permissions, setPermissions] = useState<string[]>([]);
    const [systemPermissions, setSystemPermissions] = useState<string[]>([]);
    const [permissionsLoaded, setPermissionsLoaded] = useState(false); // Флаг загрузки прав

    useEffect(() => {
        const checkAuth = async () => {
            setIsLoading(true);
            const storedToken = localStorage.getItem('accessToken');

            if (storedToken && !isTokenExpired(storedToken)) {
                const decoded = parseJwt(storedToken);
                setIsAuthenticated(true);
                setRole(decoded?.role as 'admin' | 'owner' | 'employer' | null);
                setUsername(decoded?.username || null);
            } else {
                try {
                    const response = await api.get('/api/v1/auth/refresh', { withCredentials: true });

                    if (response.data.accessToken) {
                        localStorage.setItem('accessToken', response.data.accessToken);
                        const decoded = parseJwt(response.data.accessToken);
                        setIsAuthenticated(true);
                        setRole(decoded?.role as 'admin' | 'owner' | 'employer' | null);
                        setUsername(decoded?.username || null);
                    } else {
                        setIsAuthenticated(false);
                        setRole(null);
                        setUsername(null);
                    }
                } catch (error) {
                    console.error('Ошибка при обновлении токена:', error);
                    setIsAuthenticated(false);
                    setRole(null);
                    setUsername(null);
                }
            }
            setIsLoading(false);
        };

        checkAuth();
    }, []);

    const getPermissionsForWarehouse = async (warehouseId: string, username: string) => {
        if (role !== 'employer' || permissionsLoaded) return; // Запрос только если ещё не загружено

        try {
            const response = await api.post(`/api/v1/employer/permission/${warehouseId}/get_my_permissions`, { username });

            setPermissions(response.data.permissions || []);
            setPermissionsLoaded(true); // Устанавливаем, что права загружены
        } catch (error) {
            console.error('Ошибка получения прав на склад:', error);
        }
    };

    const getSystemPermissions = async () => {
        try {
            const response = await api.get(`/api/v1/${role}/permission`);
            setSystemPermissions(response.data || []);
        } catch (error) {
            console.error('Ошибка получения прав', error);
        }
    };

    const loginWithEmail = async (email: string, password: string) => {
        try {
            const response = await api.post('/api/v1/auth/sign-in-email', { email, password });

            if (response.data.accessToken) {
                localStorage.setItem('accessToken', response.data.accessToken);
                const decoded = parseJwt(response.data.accessToken);
                setIsAuthenticated(true);
                setRole(decoded?.role as 'admin' | 'owner' | 'employer' | null);
                setUsername(decoded?.username || null);
                setPermissionsLoaded(false); // Сбрасываем флаг загрузки прав при новом входе
            }
        } catch (error) {
            console.error('Ошибка входа по email:', error);
            throw new Error('Неверные учетные данные');
        }
    };

    const loginWithPhone = async (phone_number: string, password: string) => {
        try {
            const response = await api.post('/api/v1/auth/sign-in-phone', { phone_number, password }, { withCredentials: true });

            if (response.data.accessToken) {
                localStorage.setItem('accessToken', response.data.accessToken);
                const decoded = parseJwt(response.data.accessToken);
                setIsAuthenticated(true);
                setRole(decoded?.role as 'admin' | 'owner' | 'employer' | null);
                setUsername(decoded?.username || null);
                setPermissionsLoaded(false); // Сбрасываем флаг загрузки прав при новом входе
            }
        } catch (error) {
            console.error('Ошибка входа по телефону:', error);
            throw new Error('Неверный код подтверждения');
        }
    };

    const logout = async () => {
        try {
            await api.get('/api/v1/logout', { withCredentials: true });
        } catch (error) {
            console.error('Ошибка при выходе:', error);
        }
        localStorage.removeItem('accessToken');
        setIsAuthenticated(false);
        setRole(null);
        setUsername(null);
        setPermissions([]);
        setPermissionsLoaded(false); // Сбрасываем флаг загрузки прав
    };

    return (
        <AuthContext.Provider value={{ isAuthenticated, role, username, isLoading, permissions, systemPermissions, loginWithEmail, loginWithPhone, logout, getPermissionsForWarehouse, getSystemPermissions }}>
            {!isLoading && children}
        </AuthContext.Provider>
    );
};
