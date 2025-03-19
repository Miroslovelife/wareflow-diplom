export interface User {
    exp: number;
    role: 'admin' | 'owner' | 'employer';
    userId: string;
}

export interface AuthState {
    user: User | null;
    access-token: string | null;
}
