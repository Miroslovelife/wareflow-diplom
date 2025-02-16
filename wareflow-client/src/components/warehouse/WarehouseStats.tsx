import { Warehouse } from '../../types/warehouse';

interface WarehouseStatsProps {
  warehouse: Warehouse;
}

export default function WarehouseStats({ warehouse }: WarehouseStatsProps) {
  const capacityPercentage = (warehouse.usedCapacity / warehouse.totalCapacity) * 100;

  return (
    <div className="bg-white rounded-lg shadow-md p-6 mb-8">
      <div className="flex justify-between items-center mb-6">
        <div>
          <h2 className="text-2xl font-bold">{warehouse.name}</h2>
          <p className="text-gray-600">{warehouse.location}</p>
        </div>
      </div>
      
      <div className="grid grid-cols-3 gap-4">
        <div className="bg-indigo-50 p-4 rounded-lg">
          <p className="text-sm text-indigo-600 mb-1">Общая загруженность</p>
          <p className="text-2xl font-bold">{capacityPercentage.toFixed(1)}%</p>
        </div>
        <div className="bg-indigo-50 p-4 rounded-lg">
          <p className="text-sm text-indigo-600 mb-1">Всего товаров</p>
          <p className="text-2xl font-bold">
            {warehouse.zones.reduce((sum, zone) => sum + zone.items, 0)}
          </p>
        </div>
        <div className="bg-indigo-50 p-4 rounded-lg">
          <p className="text-sm text-indigo-600 mb-1">Активные зоны</p>
          <p className="text-2xl font-bold">{warehouse.activeZones}/{warehouse.totalZones}</p>
        </div>
      </div>
    </div>
  );
}