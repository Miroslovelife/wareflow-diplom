// src/hooks/useRole.ts
import { useCheckAuth } from './useCheckAuth.tsx';

export const useRole = () => {
    const { state } = useCheckAuth();

    const hasRole = (requiredRoles: string[]) => {
        if (!state.user) return false;
        return requiredRoles.includes(state.user.role);
    };

    return {
        isAdmin: () => hasRole(['admin']),
        isOwner: () => hasRole(['owner']),
        isEmployer: () => hasRole(['employer']),
        isAuthenticated: () => !!state.user
    };
};