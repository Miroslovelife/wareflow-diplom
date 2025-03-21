import { useEffect, useState } from 'react';
import { User, Mail, Shield, Phone } from 'lucide-react';
import { api } from '../utils/api';
import { useAuth } from '../contexts/AuthProvider';

interface UserProfile {
    phone_number: string;
    username: string;
    first_name: string;
    last_name: string;
    surname: string;
    email: string;
    role: string;
}

export default function Profile() {
    const { role, isAuthenticated } = useAuth();
    const [profile, setProfile] = useState<UserProfile | null>(null);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchProfile = async () => {
            if (!isAuthenticated) {
                setError('Ошибка: пользователь не авторизован.');
                setIsLoading(false);
                return;
            }

            try {
                const response = await api.get(`/api/v1/profile`);
                setProfile(response.data.profile);
            } catch (err) {
                setError('Ошибка загрузки профиля');
                console.error(err);
            } finally {
                setIsLoading(false);
            }
        };

        fetchProfile();
    }, [role, isAuthenticated]);

    if (isLoading) return <div className="text-center py-6">Загрузка...</div>;
    if (error) return <div className="text-center py-6 text-red-600">{error}</div>;

    return (
        <div className="min-h-screen bg-gray-50 py-8">
            <div className="max-w-3xl mx-auto px-4">
                <div className="bg-white rounded-lg shadow-md p-6">
                    <h2 className="text-2xl font-bold mb-4">Профиль пользователя</h2>
                    <div className="flex items-center space-x-4 mb-4">
                        <User className="h-10 w-10 text-indigo-600" />
                        <div>
                            <p className="text-xl font-semibold">{profile?.first_name} {profile?.last_name} {profile?.surname}</p>
                            <p className="text-gray-500">@{profile?.username}</p>
                        </div>
                    </div>
                    <div className="flex items-center space-x-4 mb-4">
                        <Mail className="h-6 w-6 text-indigo-600" />
                        <p className="text-gray-700">{profile?.email}</p>
                    </div>
                    <div className="flex items-center space-x-4 mb-4">
                        <Phone className="h-6 w-6 text-indigo-600" />
                        <p className="text-gray-700">{profile?.phone_number}</p>
                    </div>
                    <div className="flex items-center space-x-4">
                        <Shield className="h-6 w-6 text-indigo-600" />
                        <p className="text-gray-700">
                            Роль: {profile?.role === 'owner' ? 'Владелец складов' : profile?.role === 'employer' ? 'Сотрудник складов' : ''}
                        </p>
                    </div>
                </div>
            </div>
        </div>
    );
}
