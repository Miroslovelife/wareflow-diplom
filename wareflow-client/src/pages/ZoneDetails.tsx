import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { api } from '../utils/api';
import { useAuth } from '../contexts/AuthProvider';


interface Zone {
    id: number;
    name: string;
    capacity: number;
    warehouseName: string;
    warehouseAddress: string;
    // Другие данные о зоне, которые могут быть полезны
}

export default function ZoneDetails() {
    const { role, isAuthenticated } = useAuth();
    const { zoneId, warehouseId } = useParams(); // Получаем id зоны и склада из URL
    const [zone, setZone] = useState<Zone | null>(null);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const fetchZoneData = async () => {
            if (!isAuthenticated || !role || !zoneId || !warehouseId) {
                setError('Ошибка: пользователь не авторизован или данные не указаны.');
                setIsLoading(false);
                return;
            }

            try {
                // Получаем информацию о зоне с использованием эндпоинта
                const response = await api.get(`/api/v1/${role}/warehouse/${warehouseId}/zone/${zoneId}`);
                setZone(response.data.zone); // Устанавливаем данные о зоне в состояние
            } catch (err) {
                setError('Ошибка загрузки данных о зоне');
                console.error(err);
            } finally {
                setIsLoading(false);
            }
        };

        fetchZoneData();
    }, [role, isAuthenticated, zoneId, warehouseId]);

    // Если данные все еще загружаются, отображаем индикатор загрузки
    if (isLoading) return <div className="text-center py-6">Загрузка...</div>;

    // Если произошла ошибка при загрузке данных
    if (error) return <div className="text-center py-6 text-red-600">{error}</div>;

    return (
        <div className="min-h-screen bg-gray-50 py-8">
            <div className="max-w-7xl mx-auto px-4">
                {zone ? (
                    <>
                        <h2 className="text-3xl font-bold mb-4">{zone.name}</h2>
                        <p className="text-gray-600 mb-6">Склад: {zone.warehouseName}</p>
                        <p className="text-gray-600 mb-6">Адрес склада: {zone.warehouseAddress}</p>
                        <p className="text-gray-600 mb-6">Вместимость: {zone.capacity}</p>

                        {/* Можно добавить дополнительные данные о зоне */}
                    </>
                ) : (
                    <p className="text-center text-gray-600">Данные о зоне не найдены.</p>
                )}
            </div>
        </div>
    );
}
