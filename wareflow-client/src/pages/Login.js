import React, { useState } from 'react';
import axios from 'axios';

const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            // Здесь вы должны отправить POST-запрос на ваш API для аутентификации
            const response = await axios.post('http://api.example.com/login', { email, password });

            // После успешной аутентификации, сохраняем токен в localStorage
            localStorage.setItem('token', response.data.token);

            // Здесь можно добавить редирект на главную страницу после успешного входа
            window.location.href = '/';
        } catch (error) {
            console.error('Ошибка при авторизации:', error);
            alert('Неправильный email или пароль');
        }
    };

    return (
        <div className="login-page">
            <h1>Login</h1>
            <form onSubmit={handleSubmit}>
                <input
                    type="email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    placeholder="Email"
                    required
                />
                <input
                    type="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    placeholder="Password"
                    required
                />
                <button type="submit">Login</button>
            </form>
        </div>
    );
};

export default Login;
