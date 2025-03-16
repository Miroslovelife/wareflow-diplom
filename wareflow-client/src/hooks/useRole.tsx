// src/hooks/useRole.ts
import { useAuth } from './useAuth';

export const useRole = () => {
    const { state } = useAuth();

    const hasRole = (requiredRoles: string[]) => {
        if (!state.user) return false;
        return requiredRoles.includes(state.user.role);
    };

    return {
        isAdmin: () => hasRole(['admin']),
        isModerator: () => hasRole(['moderator', 'admin']),
        isAuthenticated: () => !!state.user
    };
};