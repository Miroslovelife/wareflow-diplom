import {useEffect, useState} from 'react';
import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import * as yup from 'yup';
import { api } from '../utils/api';
import { useNavigate } from 'react-router-dom';

// ✅ Определяем схему валидации
const schema = yup.object().shape({
    username: yup.string().required('Имя пользователя обязательно'),
    email: yup.string().email('Некорректный email').required('Email обязателен'),
    password: yup.string().min(6, 'Пароль должен быть минимум 6 символов').required('Пароль обязателен'),
    first_name: yup.string().required('Имя обязательно'),
    last_name: yup.string().required('Фамилия обязательна'),
    surname: yup.string().required('Отчество обязательно'),
    phone_number: yup
        .string()
        .matches(/^\+?\d{10,15}$/, 'Некорректный номер телефона')
        .required('Номер телефона обязателен'),
    role: yup.string().oneOf(['employer', 'owner'], 'Выберите роль').required('Роль обязательна'),
});

// ✅ Определяем интерфейс данных формы
interface RegisterForm {
    phone_number: string;
    username: string;
    first_name: string;
    last_name: string;
    surname: string;
    email: string;
    password: string;
    role: 'employer' | 'owner';
}

export default function Register() {
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState(false);
    const navigate = useNavigate();

    const { register, handleSubmit, formState: { errors } } = useForm<RegisterForm>({
        resolver: yupResolver(schema),
    });

    useEffect(() => {
        const script = document.createElement("script");
        script.src = "../../app.js"; // Путь к файлу app.js
        script.async = true;

        document.body.appendChild(script);

        return () => {
            document.body.removeChild(script);
        };
    }, []);

    // ✅ Функция отправки данных на сервер
    const onSubmit = async (data: RegisterForm) => {
        setLoading(true);
        setError(null);

        try {
            const payload = {
                phone_number: data.phone_number,
                username: data.username,
                first_name: data.first_name,
                last_name: data.last_name,
                surname: data.surname,
                email: data.email,
                password: data.password,
                role: data.role,
            };

            console.log("Отправляемые данные:", JSON.stringify(payload, null, 2)); // Для отладки

            const response = await api.post('/api/v1/auth/sign-up', JSON.stringify(payload), {
                headers: {
                    'Content-Type': 'application/json'
                }
            });

            console.log("Ответ сервера:", response.data);
            navigate('/login');
        } catch (err) {
            setError('Ошибка регистрации, попробуйте снова.');
            console.error("Ошибка:", err);
        } finally {
            setLoading(false);
        }
    };


    return (
        <div className="relative">
            <div id="particles-js" className="absolute top-0 left-0 w-full h-full z-0"/>
            <div className="min-h-screen flex items-center justify-center bg-gray-100">
                <div className="relative bg-white p-8 rounded-lg shadow-md w-full max-w-md">
                    <h2 className="text-2xl font-bold mb-6 text-center">Регистрация</h2>

                    {error && <p className="text-red-600 text-center mb-4">{error}</p>}

                    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
                        {/* ✅ Имя пользователя */}
                        <div>
                            <label className="block text-sm font-medium">Имя пользователя</label>
                            <input {...register('username')} className="w-full border px-3 py-2 rounded"/>
                            <p className="text-red-500 text-sm">{errors.username?.message}</p>
                        </div>

                        {/* ✅ Email */}
                        <div>
                            <label className="block text-sm font-medium">Email</label>
                            <input {...register('email')} type="email" className="w-full border px-3 py-2 rounded"/>
                            <p className="text-red-500 text-sm">{errors.email?.message}</p>
                        </div>

                        {/* ✅ Пароль */}
                        <div>
                            <label className="block text-sm font-medium">Пароль</label>
                            <input {...register('password')} type="password"
                                   className="w-full border px-3 py-2 rounded"/>
                            <p className="text-red-500 text-sm">{errors.password?.message}</p>
                        </div>

                        {/* ✅ Имя */}
                        <div>
                            <label className="block text-sm font-medium">Имя</label>
                            <input {...register('first_name')} className="w-full border px-3 py-2 rounded"/>
                            <p className="text-red-500 text-sm">{errors.first_name?.message}</p>
                        </div>

                        {/* ✅ Фамилия */}
                        <div>
                            <label className="block text-sm font-medium">Фамилия</label>
                            <input {...register('last_name')} className="w-full border px-3 py-2 rounded"/>
                            <p className="text-red-500 text-sm">{errors.last_name?.message}</p>
                        </div>

                        {/* ✅ Отчество */}
                        <div>
                            <label className="block text-sm font-medium">Отчество</label>
                            <input {...register('surname')} className="w-full border px-3 py-2 rounded"/>
                            <p className="text-red-500 text-sm">{errors.surname?.message}</p>
                        </div>

                        {/* ✅ Номер телефона */}
                        <div>
                            <label className="block text-sm font-medium">Номер телефона</label>
                            <input {...register('phone_number')} className="w-full border px-3 py-2 rounded"/>
                            <p className="text-red-500 text-sm">{errors.phone_number?.message}</p>
                        </div>

                        {/* ✅ Выбор роли */}
                        <div>
                            <label className="block text-sm font-medium">Роль</label>
                            <select {...register('role')} className="w-full border px-3 py-2 rounded">
                                <option value="">Выберите роль</option>
                                <option value="employer">Работник склада</option>
                                <option value="owner">Владелец склада</option>
                            </select>
                            <p className="text-red-500 text-sm">{errors.role?.message}</p>
                        </div>

                        {/* ✅ Кнопка регистрации */}
                        <button type="submit"
                                className="w-full bg-indigo-600 text-white py-2 rounded hover:bg-indigo-700"
                                disabled={loading}>
                            {loading ? 'Регистрация...' : 'Зарегистрироваться'}
                        </button>
                    </form>
                </div>
            </div>
        </div>

    );
}
