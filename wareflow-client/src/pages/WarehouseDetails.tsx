import { useParams } from 'react-router-dom';
import { warehouses } from '../data/warehouses';
import WarehouseStats from '../components/warehouse/WarehouseStats';
import ZoneList from '../components/warehouse/ZoneList';

export default function WarehouseDetails() {
  const { id } = useParams();
  const warehouse = warehouses.find(w => w.id === Number(id));

  if (!warehouse) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center">
          <h2 className="text-2xl font-bold text-gray-900">Склад не найден</h2>
          <p className="text-gray-600">Запрошенный склад не существует</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50 py-8">
      <div className="max-w-7xl mx-auto px-4">
        <WarehouseStats warehouse={warehouse} />
        <ZoneList zones={warehouse.zones} />
      </div>
    </div>
  );
}