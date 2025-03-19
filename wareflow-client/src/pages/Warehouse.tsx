import { useEffect, useState } from 'react';
import { Package, ArrowRight } from 'lucide-react';
import { api } from '../utils/api';
import { useAuth } from '../contexts/AuthProvider';

interface Warehouse {
  id: number;
  name: string;
  address: string;
}

export default function Warehouses() {
  const { role, isAuthenticated } = useAuth(); // Получаем роль и статус аутентификации
  const [warehouses, setWarehouses] = useState<Warehouse[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchWarehouses = async () => {
      if (!isAuthenticated || !role) {
        setError('Ошибка: пользователь не авторизован или роль не определена.');
        setIsLoading(false);
        return;
      }

      const endpoint = `/api/v1/${role}/warehouse`;

      try {
        const response = await api.get(endpoint);
        setWarehouses(response.data.warehouses); // Извлекаем из объекта "warehouses"
      } catch (err) {
        setError('Ошибка загрузки складов');
        console.error(err);
      } finally {
        setIsLoading(false);
      }
    };

    fetchWarehouses();
  }, [role, isAuthenticated]);

  if (isLoading) return <div className="text-center py-6">Загрузка...</div>;
  if (error) return <div className="text-center py-6 text-red-600">{error}</div>;

  return (
      <div className="min-h-screen bg-gray-50 py-8">
        <div className="max-w-7xl mx-auto px-4">
          <div className="bg-white rounded-lg shadow-md p-6 mb-8">
            <h2 className="text-2xl font-bold mb-4">Список складов</h2>
          </div>

          <div className="grid md:grid-cols-2 gap-6">
            {warehouses.map((warehouse) => (
                <div key={warehouse.id} className="bg-white rounded-lg shadow-md p-6">
                  <div className="flex justify-between items-center mb-4">
                    <h3 className="text-xl font-semibold">{warehouse.name.trim()}</h3>
                    <button className="text-indigo-600 hover:text-indigo-800 flex items-center">
                      Подробнее <ArrowRight className="h-4 w-4 ml-1" />
                    </button>
                  </div>
                  <div className="flex items-center space-x-4">
                    <Package className="h-8 w-8 text-indigo-600" />
                    <div>
                      <p className="text-sm text-gray-600">Адрес: {warehouse.id}</p>
                    </div>
                  </div>
                </div>
            ))}
          </div>
        </div>
      </div>
  );
}
