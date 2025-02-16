import { warehouses } from '../data/warehouses';
import WarehouseCard from '../components/warehouse/WarehouseCard';

export default function Warehouses() {
  return (
    <div className="min-h-screen bg-gray-50 py-8">
      <div className="max-w-7xl mx-auto px-4">
        <div className="mb-8">
          <h1 className="text-2xl font-bold text-gray-900">Управление складами</h1>
          <p className="text-gray-600">Обзор и управление всеми складскими помещениями</p>
        </div>

        <div className="grid md:grid-cols-2 gap-6">
          {warehouses.map((warehouse) => (
            <WarehouseCard key={warehouse.id} warehouse={warehouse} />
          ))}
        </div>
      </div>
    </div>
  );
}