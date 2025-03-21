import { useState, useEffect, useRef } from 'react';
import { Warehouse, Home, LogIn, UserPlus, User, LogOut, Package, Menu, X, QrCode } from 'lucide-react';
import { Link } from 'react-router-dom';
import { useAuth } from '../contexts/AuthProvider';

export default function Navbar() {
    const { username, isAuthenticated, logout, role } = useAuth();
    const [menuOpen, setMenuOpen] = useState(false);
    const [burgerMenuOpen, setBurgerMenuOpen] = useState(false);
    const menuRef = useRef<HTMLDivElement | null>(null);

    // Обработчик клика вне меню
    useEffect(() => {
        const handleClickOutside = (event: MouseEvent) => {
            if (menuRef.current && !menuRef.current.contains(event.target as Node)) {
                setMenuOpen(false);
            }
        };

        if (menuOpen) {
            document.addEventListener('mousedown', handleClickOutside);
        } else {
            document.removeEventListener('mousedown', handleClickOutside);
        }

        return () => document.removeEventListener('mousedown', handleClickOutside);
    }, [menuOpen]);

    return (
        <nav className="bg-indigo-600 text-white">
            <div className="max-w-7xl mx-auto px-4">
                <div className="flex items-center justify-between h-16">
                    {/* Логотип */}
                    <Link to="/" className="flex items-center space-x-2 font-bold text-xl">
                        <Warehouse className="h-6 w-6" />
                        <span>Ware¯\_(ツ)_/¯Flow</span>
                    </Link>

                    {/* Бургер меню для маленьких экранов */}
                    {isAuthenticated ? (
                        <div className="block lg:hidden">
                            <button
                                className="flex items-center text-white"
                                onClick={() => setBurgerMenuOpen(!burgerMenuOpen)}
                            >
                                {burgerMenuOpen ? (
                                    <X className="h-6 w-6" />
                                ) : (
                                    <Menu className="h-6 w-6" />
                                )}
                            </button>
                        </div>) : null}

                    {/* Навигация для авторизованных пользователей */}
                    {isAuthenticated ? (
                        <div className="hidden lg:flex items-center space-x-6">
                            {role !== 'employer' && (
                                <Link to="/" className="flex items-center space-x-1 hover:text-indigo-200">
                                    <Home className="h-5 w-5" />
                                    <span>Главная</span>
                                </Link>
                            )}
                            <Link to="/warehouses" className="flex items-center space-x-1 hover:text-indigo-200">
                                <Warehouse className="h-5 w-5" />
                                <span>Склады</span>
                            </Link>
                            {role !== 'employer' && (
                                <Link to="/products" className="flex items-center space-x-1 hover:text-indigo-200">
                                    <Package className="h-5 w-5" />
                                    <span>Товары</span>
                                </Link>
                            )}
                            <Link to="/scanner" className="flex items-center space-x-1 hover:text-indigo-200">
                                <QrCode className="h-5 w-5" />
                                <span>Сканер</span>
                            </Link>

                            {/* Иконка пользователя с выпадающим меню */}
                            <div className="relative" ref={menuRef}>
                                <button
                                    className="flex items-center space-x-2 bg-indigo-500 px-3 py-2 rounded-full hover:bg-indigo-700"
                                    onClick={() => setMenuOpen(!menuOpen)}
                                >
                                    <User className="h-6 w-6" />
                                </button>

                                {menuOpen && (
                                    <div className="absolute right-0 mt-2 w-40 bg-white text-black rounded-md shadow-lg z-50">
                                        <Link to="/profile" className="block px-4 py-2 hover:bg-gray-100">
                                            {username}
                                        </Link>
                                        <button
                                            onClick={() => { logout(); setMenuOpen(false); }}
                                            className="flex w-full text-left px-4 py-2 hover:bg-gray-100"
                                        >
                                            <LogOut className="h-5 w-5 mr-2" /> Выйти
                                        </button>
                                    </div>
                                )}
                            </div>
                        </div>
                    ) : (
                        <div className="flex space-x-4">
                            <Link to="/login" className="flex items-center space-x-1 px-4 py-2 bg-white text-indigo-600 rounded-md hover:bg-indigo-100">
                                <LogIn className="h-5 w-5" />
                                <span>Войти</span>
                            </Link>
                            <Link to="/reg" className="flex items-center space-x-1 px-4 py-2 bg-indigo-500 rounded-md hover:bg-indigo-700">
                                <UserPlus className="h-5 w-5" />
                                <span>Регистрация</span>
                            </Link>
                        </div>
                    )}
                </div>
            </div>

            {/* Бургер меню */}
            {burgerMenuOpen && isAuthenticated && (
                <div className="lg:hidden bg-indigo-600 text-white py-4">
                    <div className="space-y-4 px-6">
                        {role !== 'employer' && (
                            <Link to="/" className="flex items-center space-x-1">
                                <Home className="h-5 w-5" />
                                <span>Главная</span>
                            </Link>
                        )}
                        <Link to="/warehouses" className="flex items-center space-x-1">
                            <Warehouse className="h-5 w-5" />
                            <span>Склады</span>
                        </Link>
                        {role !== 'employer' && (
                            <Link to="/products" className="flex items-center space-x-1">
                                <Package className="h-5 w-5" />
                                <span>Товары</span>
                            </Link>
                        )}
                        <Link to="/scanner" className="flex items-center space-x-1">
                            <QrCode className="h-5 w-5" />
                            <span>Сканер</span>
                        </Link>
                        <div className="flex items-center space-x-2 mt-4">
                            <Link to="/profile" className="flex items-center space-x-2">
                                <User className="h-5 w-5" />
                                <span>{username}</span>
                            </Link>
                            <button
                                onClick={() => { logout(); setBurgerMenuOpen(false); }}
                                className="flex items-center space-x-2 mt-4 px-4 py-2 text-white bg-indigo-500 rounded-md hover:bg-indigo-700"
                            >
                                <LogOut className="h-5 w-5" />
                                <span>Выйти</span>
                            </button>
                        </div>
                    </div>
                </div>
            )}
        </nav>
    );
}
