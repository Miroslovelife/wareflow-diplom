import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import { Package, ArrowRight, Plus } from 'lucide-react';
import { api } from '../utils/api';
import { useAuth } from '../contexts/AuthProvider';

interface Warehouse {
  id: number;
  name: string;
  address: string;
}

export default function Warehouses() {
  const { role, isAuthenticated } = useAuth();
  const [warehouses, setWarehouses] = useState<Warehouse[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [newWarehouse, setNewWarehouse] = useState({ name: '', address: '' });
  const [isSubmitting, setIsSubmitting] = useState(false);

  useEffect(() => {
    const fetchWarehouses = async () => {
      if (!isAuthenticated || !role) {
        setError('Ошибка: пользователь не авторизован или роль не определена.');
        setIsLoading(false);
        return;
      }

      setIsLoading(true);
      const endpoint = `/api/v1/${role}/warehouse`;

      try {
        const response = await api.get(endpoint);
        setWarehouses(response.data.warehouses || []);
      } catch (err) {
        setError('Ошибка загрузки складов');
        console.error(err);
      } finally {
        setIsLoading(false);
      }
    };

    fetchWarehouses();
  }, [role, isAuthenticated]);

  const handleCreateWarehouse = async () => {
    if (!newWarehouse.name.trim() || !newWarehouse.address.trim()) {
      alert('Введите имя и адрес склада.');
      return;
    }
    console.log('Создаваемый склад:', newWarehouse);
    setIsSubmitting(true);

    try {
      const response = await api.post(`/api/v1/${role}/warehouse`, newWarehouse);
      setWarehouses([...warehouses, response.data]);
      setIsModalOpen(false);
      setNewWarehouse({ name: '', address: '' });
    } catch (err) {
      console.error('Ошибка при создании склада:', err);
      alert('Не удалось создать склад.');
    } finally {
      setIsSubmitting(false);
    }

    const endpoint = `/api/v1/${role}/warehouse`;

    try {
      const response = await api.get(endpoint);
      setWarehouses(response.data.warehouses || []);
    } catch (err) {
      setError('Ошибка загрузки складов');
      console.error(err);
    } finally {
      setIsLoading(false);
    }
  };

  if (isLoading) return <div className="text-center py-6">Загрузка...</div>;
  if (error) return <div className="text-center py-6 text-red-600">{error}</div>;

  return (
      <div className="min-h-screen bg-gray-50 py-8">
        <div className="max-w-7xl mx-auto px-4">
          <div className="bg-white rounded-lg shadow-md p-6 mb-8 flex justify-between items-center">
            <h2 className="text-2xl font-bold">Список складов</h2>
            {(role === 'admin' || role === 'owner') && (
                <button
                    onClick={() => setIsModalOpen(true)}
                    className="flex items-center bg-indigo-600 text-white px-4 py-2 rounded-lg shadow hover:bg-indigo-700"
                >
                  <Plus className="h-5 w-5 mr-2" /> Создать склад
                </button>
            )}
          </div>

          <div className="grid md:grid-cols-2 gap-6">
            {warehouses.length === 0 ? (
                <div className="col-span-2 text-center text-gray-600">
                  Нет доступных складов.
                </div>
            ) : (
                warehouses.map((warehouse) => (
                    <div key={warehouse.id} className="bg-white rounded-lg shadow-md p-6">
                      <div className="flex justify-between items-center mb-4">
                        <h3 className="text-xl font-semibold">
                          {warehouse.name ? warehouse.name.trim() : 'Без названия'}
                        </h3>
                        <Link
                            to={`/warehouses/${encodeURIComponent(warehouse.id)}`}
                            className="text-indigo-600 hover:text-indigo-800 flex items-center"
                        >
                          Подробнее <ArrowRight className="h-4 w-4 ml-1" />
                        </Link>
                      </div>
                      <div className="flex items-center space-x-4">
                        <Package className="h-8 w-8 text-indigo-600" />
                        <div>
                          <p className="text-sm text-gray-600">
                            Адрес: {warehouse.address ? warehouse.address.trim() : 'Не указан'}
                          </p>
                        </div>
                      </div>
                    </div>
                ))
            )}
          </div>
        </div>

        {/* Модальное окно */}
        {isModalOpen && (
            <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
              <div className="bg-white p-6 rounded-lg shadow-lg w-96">
                <h2 className="text-xl font-bold mb-4">Создать склад</h2>
                <label className="block text-sm font-medium text-gray-700">Название склада</label>
                <input
                    type="text"
                    value={newWarehouse.name}
                    onChange={(e) => setNewWarehouse({ ...newWarehouse, name: e.target.value })}
                    className="w-full p-2 border rounded-lg mb-3"
                />
                <label className="block text-sm font-medium text-gray-700">Адрес</label>
                <input
                    type="text"
                    value={newWarehouse.address}
                    onChange={(e) => setNewWarehouse({ ...newWarehouse, address: e.target.value })}
                    className="w-full p-2 border rounded-lg mb-4"
                />
                <div className="flex justify-end space-x-2">
                  <button
                      onClick={() => setIsModalOpen(false)}
                      className="px-4 py-2 bg-gray-300 text-gray-800 rounded-lg"
                  >
                    Отмена
                  </button>
                  <button
                      onClick={handleCreateWarehouse}
                      disabled={isSubmitting}
                      className="px-4 py-2 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700"
                  >
                    {isSubmitting ? 'Создание...' : 'Создать'}
                  </button>
                </div>
              </div>
            </div>
        )}
      </div>
  );
}
