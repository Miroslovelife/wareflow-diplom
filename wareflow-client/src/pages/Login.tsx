import { useEffect, useState } from 'react';
import { useAuth } from '../contexts/AuthProvider';
import { useNavigate } from 'react-router-dom';

export const Login: React.FC = () => {
    const { loginWithEmail, loginWithPhone, isAuthenticated } = useAuth();
    const navigate = useNavigate();

    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [phone, setPhone] = useState('');
    const [currentForm, setCurrentForm] = useState<'email' | 'phone'>('email');
    const [errorMessage, setErrorMessage] = useState('');

    // **Если пользователь уже залогинен — отправляем на главную**
    useEffect(() => {
        if (isAuthenticated) {
            navigate('/');
        }
    }, [isAuthenticated, navigate]);

    const handleEmailLogin = async (e: React.FormEvent) => {
        e.preventDefault();
        setErrorMessage('');

        try {
            await loginWithEmail(email, password);
            navigate('/'); // Успешный вход -> редирект
        } catch (error) {
            setErrorMessage('Такого пользователя не найдено');
        }
    };

    const handlePhoneLogin = async (e: React.FormEvent) => {
        e.preventDefault();
        setErrorMessage('');

        try {
            await loginWithPhone(phone, password);
            navigate('/');
        } catch (error) {
            setErrorMessage('Такого пользователя не найдено');
        }
    };

    return (
        <div className="min-h-screen bg-gray-50 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
            <div className="max-w-md w-full space-y-8">
                <h2 className="text-center text-3xl font-extrabold text-gray-900">
                    Вход в систему
                </h2>

                <div className="bg-white shadow rounded-lg p-8 space-y-6">
                    <div className="flex justify-center space-x-4 mb-4">
                        <button
                            type="button"
                            onClick={() => setCurrentForm('email')}
                            className={`px-4 py-2 rounded-md text-sm font-medium ${
                                currentForm === 'email'
                                    ? 'bg-indigo-100 text-indigo-800'
                                    : 'text-gray-700 hover:bg-gray-50'
                            }`}
                        >
                            По email
                        </button>
                        <button
                            type="button"
                            onClick={() => setCurrentForm('phone')}
                            className={`px-4 py-2 rounded-md text-sm font-medium ${
                                currentForm === 'phone'
                                    ? 'bg-indigo-100 text-indigo-800'
                                    : 'text-gray-700 hover:bg-gray-50'
                            }`}
                        >
                            По телефону
                        </button>
                    </div>

                    {errorMessage && (
                        <div className="text-red-500 text-sm">{errorMessage}</div>
                    )}

                    {currentForm === 'email' && (
                        <form onSubmit={handleEmailLogin} className="space-y-4">
                            <input
                                type="email"
                                value={email}
                                onChange={(e) => setEmail(e.target.value)}
                                placeholder="Email"
                                required
                                className="w-full border px-3 py-2 rounded"
                            />
                            <input
                                type="password"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                                placeholder="Пароль"
                                required
                                className="w-full border px-3 py-2 rounded"
                            />
                            <button
                                type="submit"
                                className="w-full bg-indigo-600 text-white py-2 rounded"
                            >
                                Войти
                            </button>
                        </form>
                    )}

                    {currentForm === 'phone' && (
                        <form onSubmit={handlePhoneLogin} className="space-y-4">
                            <input
                                type="text"
                                value={phone}
                                onChange={(e) => setPhone(e.target.value)}
                                placeholder="Телефон"
                                required
                                className="w-full border px-3 py-2 rounded"
                            />
                            <input
                                type="password"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                                placeholder="Пароль"
                                required
                                className="w-full border px-3 py-2 rounded"
                            />
                            <button
                                type="submit"
                                className="w-full bg-indigo-600 text-white py-2 rounded"
                            >
                                Войти
                            </button>
                        </form>
                    )}
                </div>
            </div>
        </div>
    );
};
