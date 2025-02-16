import { Package, ArrowRight } from 'lucide-react';
import { Zone } from '../../types/warehouse';

interface ZoneListProps {
  zones: Zone[];
}

export default function ZoneList({ zones }: ZoneListProps) {
  return (
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
  );
}