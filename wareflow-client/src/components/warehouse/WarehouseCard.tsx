import { ArrowRight } from 'lucide-react';
import { Warehouse } from '../../types/warehouse';
import { Link } from 'react-router-dom';

interface WarehouseCardProps {
  warehouse: Warehouse;
}

export default function WarehouseCard({ warehouse }: WarehouseCardProps) {
  const capacityPercentage = (warehouse.usedCapacity / warehouse.totalCapacity) * 100;

  return (
    <div className="bg-white rounded-lg shadow-md p-6">
      <div className="flex justify-between items-start mb-4">
        <div>
          <h3 className="text-xl font-semibold">{warehouse.name}</h3>
          <p className="text-sm text-gray-600">{warehouse.location}</p>
        </div>
        <Link 
          to={`/warehouse/${warehouse.id}`}
          className="text-indigo-600 hover:text-indigo-800 flex items-center"
        >
          Подробнее <ArrowRight className="h-4 w-4 ml-1" />
        </Link>
      </div>

      <div className="space-y-4">
        <div>
          <div className="flex justify-between text-sm mb-1">
            <span>Загруженность</span>
            <span>{capacityPercentage.toFixed(1)}%</span>
          </div>
          <div className="w-full bg-gray-200 rounded-full h-2">
            <div 
              className="bg-indigo-600 rounded-full h-2" 
              style={{ width: `${capacityPercentage}%` }}
            />
          </div>
        </div>

        <div className="grid grid-cols-2 gap-4">
          <div className="bg-indigo-50 p-3 rounded-lg">
            <p className="text-sm text-indigo-600">Активные зоны</p>
            <p className="text-lg font-semibold">
              {warehouse.activeZones}/{warehouse.totalZones}
            </p>
          </div>
          <div className="bg-indigo-50 p-3 rounded-lg">
            <p className="text-sm text-indigo-600">Всего товаров</p>
            <p className="text-lg font-semibold">
              {warehouse.zones.reduce((sum, zone) => sum + zone.items, 0)}
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}