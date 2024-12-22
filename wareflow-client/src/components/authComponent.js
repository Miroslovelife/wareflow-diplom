import React, { useEffect } from 'react';
import { useAuth } from '../context/authContext';

const AuthComponent = ({ children }) => {
    const { token } = useAuth();

    useEffect(() => {
        if (!token) {
            // Редирект на страницу входа
            window.location.href = '/login';
        }
    }, [token]);

    return <>{children}</>;
};

export default AuthComponent;
