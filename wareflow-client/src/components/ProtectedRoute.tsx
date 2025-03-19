import { Navigate } from 'react-router-dom';
import { useAuth } from '../contexts/AuthProvider';

export const ProtectedRoute: React.FC<{ children: React.ReactNode }> = ({ children }) => {
    const { isAuthenticated } = useAuth();
    console.log('isAuthenticated:', isAuthenticated);

    if (!isAuthenticated) {
        return <Navigate to="/login" replace />;
    }

    return <>{children}</>; // ✅ Теперь рендерит переданный компонент
};