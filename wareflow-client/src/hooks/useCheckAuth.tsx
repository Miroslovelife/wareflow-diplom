// src/hooks/useAuth.ts
import { useContext } from 'react';
import { AuthProvider } from '../contexts/AuthProvider.tsx';

export const useAuth = () => {
    return useContext(AuthProvider);
};

