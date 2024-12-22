import React from 'react';
import { Navigate } from 'react-router-dom';
import { useAuth } from '../context/authContext'; // Убедитесь, что путь к authContext соответствует вашей структуре проекта

const ProtectedRoute = ({ children }) => {
    const { token } = useAuth();

    if (!token) {
        // Если токен отсутствует, редирект на страницу входа
        return <Navigate to="/login" />;
    }

    // Если токен присутствует, разрешаем доступ к защищенному маршруту
    return <>{children}</>;
};

export default ProtectedRoute;
