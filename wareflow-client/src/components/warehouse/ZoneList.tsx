interface Zone {
  id: number;
  name: string;
  capacity: number;
}

export default function ZoneList({ zones }: { zones: Zone[] }) {
  if (!Array.isArray(zones)) return <p className="text-center">Нет данных о зонах</p>;

  return (
      <div className="bg-white shadow-md rounded-lg p-4">
        <h3 className="text-xl font-bold mb-4">Зоны склада</h3>
        <ul>
          {zones.map((zone) => (
              <li key={zone.id} className="border-b py-2">
                {zone.name} (Вместимость: {zone.capacity})
              </li>
          ))}
        </ul>
      </div>
  );
}
