import { Warehouse, TrendingUp, Package } from 'lucide-react';

export default function Home() {
  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-7xl mx-auto px-4 py-12">
        <div className="text-center mb-16">
          <h1 className="text-4xl font-bold text-gray-900 mb-4">
            Система управления складом
          </h1>
          <p className="text-xl text-gray-600">
            Оптимизируйте складские операции с помощью нашей современной платформы
          </p>
        </div>

        <div className="grid md:grid-cols-3 gap-8">
          <div className="bg-white p-6 rounded-lg shadow-md">
            <div className="flex items-center justify-center w-12 h-12 bg-indigo-100 rounded-lg mb-4">
              <Warehouse className="h-6 w-6 text-indigo-600" />
            </div>
            <h3 className="text-xl font-semibold mb-2">Управление складом</h3>
            <p className="text-gray-600">
              Эффективное управление складскими помещениями и зонами хранения
            </p>
          </div>

          <div className="bg-white p-6 rounded-lg shadow-md">
            <div className="flex items-center justify-center w-12 h-12 bg-indigo-100 rounded-lg mb-4">
              <Package className="h-6 w-6 text-indigo-600" />
            </div>
            <h3 className="text-xl font-semibold mb-2">Учет товаров</h3>
            <p className="text-gray-600">
              Точный учет товаров с использованием QR-кодов и автоматизированной системы
            </p>
          </div>

          <div className="bg-white p-6 rounded-lg shadow-md">
            <div className="flex items-center justify-center w-12 h-12 bg-indigo-100 rounded-lg mb-4">
              <TrendingUp className="h-6 w-6 text-indigo-600" />
            </div>
            <h3 className="text-xl font-semibold mb-2">Аналитика</h3>
            <p className="text-gray-600">
              Подробная аналитика и отчеты о движении товаров
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}