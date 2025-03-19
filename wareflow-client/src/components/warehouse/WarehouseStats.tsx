interface Warehouse {
  id: number;
  name: string;
  address: string;
}

export default function WarehouseStats({ warehouse }: { warehouse: Warehouse }) {
  if (!warehouse) return <p className="text-center">Нет данных о складе</p>;

  return (
      <div className="p-0 mb-6">
        <h3 className="text-3xl font-bold">Название: {warehouse.name}</h3>
        <p className="text-gray-600">Адрес: {warehouse.address}</p>
      </div>
  );
}
