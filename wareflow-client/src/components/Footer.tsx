export default function Footer() {
    return (
        <footer className="h-auto bg-gray-900 text-white py-8">
            <div className="max-w-7xl mx-auto px-4">
                {/* Основной контент футера */}
                <div className="flex justify-between">
                    {/* Описание */}
                    <div className="w-64">
                        <h3 className="text-xl font-semibold mb-4">Система управления складом</h3>
                        <p className="text-gray-400">
                            Платформа для управления складами и аналитики товаров. Используйте для оптимизации складских операций.
                        </p>
                    </div>
                    {/* QR-код */}
                    <div className="text-center">
                        <h3 className="text-xl font-semibold mb-4">tg web_app</h3>
                        <a
                            href="https://t.me/wareflow_bot" // Замените на ваш настоящий Telegram-ссылку
                            target="_blank"
                            rel="noopener noreferrer"
                        >
                            <img
                                src="/wareflow.png" // Генерация QR-кода
                                alt="Telegram QR Code"
                                className="mx-auto"
                                width="150"
                                height="150"
                            />
                        </a>

                    </div>
                </div>

                {/* Нижняя часть футера */}
                <div className="mt-8 border-t border-gray-700 pt-4 text-center text-sm text-gray-400">
                    <p>&copy; {new Date().getFullYear()} wareflow. Все права защищены.</p>
                </div>
            </div>
        </footer>
    );
};
