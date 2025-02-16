import { Package, ArrowRight } from 'lucide-react';

export default function Warehouse() {
  const zones = [
    { id: 1, name: 'Зона A', capacity: '85%', items: 150 },
    { id: 2, name: 'Зона B', capacity: '60%', items: 98 },
    { id: 3, name: 'Зона C', capacity: '45%', items: 76 },
    { id: 4, name: 'Зона D', capacity: '90%', items: 180 },
  ];

  return (
    <div className="min-h-screen bg-gray-50 py-8">
      <div className="max-w-7xl mx-auto px-4">
        <div className="bg-white rounded-lg shadow-md p-6 mb-8">
          <h2 className="text-2xl font-bold mb-4">Основной склад</h2>
          <div className="grid grid-cols-3 gap-4 mb-6">
            <div className="bg-indigo-50 p-4 rounded-lg">
              <p className="text-sm text-indigo-600 mb-1">Общая загруженность</p>
              <p className="text-2xl font-bold">70%</p>
            </div>
            <div className="bg-indigo-50 p-4 rounded-lg">
              <p className="text-sm text-indigo-600 mb-1">Всего товаров</p>
              <p className="text-2xl font-bold">504</p>
            </div>
            <div className="bg-indigo-50 p-4 rounded-lg">
              <p className="text-sm text-indigo-600 mb-1">Активные зоны</p>
              <p className="text-2xl font-bold">4/5</p>
            </div>
          </div>
        </div>

        <div className="grid md:grid-cols-2 gap-6">
          {zones.map((zone) => (
            <div key={zone.id} className="bg-white rounded-lg shadow-md p-6">
              <div className="flex justify-between items-center mb-4">
                <h3 className="text-xl font-semibold">{zone.name}</h3>
                <button className="text-indigo-600 hover:text-indigo-800 flex items-center">
                  Подробнее <ArrowRight className="h-4 w-4 ml-1" />
                </button>
              </div>
              <div className="flex items-center space-x-4">
                <Package className="h-8 w-8 text-indigo-600" />
                <div>
                  <p className="text-sm text-gray-600">Загруженность: {zone.capacity}</p>
                  <p className="text-sm text-gray-600">Количество товаров: {zone.items}</p>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}