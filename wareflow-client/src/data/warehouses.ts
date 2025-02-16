import { Warehouse } from '../types/warehouse';

export const warehouses: Warehouse[] = [
  {
    id: 1,
    name: 'Центральный склад',
    location: 'Москва',
    totalCapacity: 1000,
    usedCapacity: 700,
    activeZones: 4,
    totalZones: 5,
    zones: [
      { id: 1, name: 'Зона A', capacity: '85%', items: 150 },
      { id: 2, name: 'Зона B', capacity: '60%', items: 98 },
      { id: 3, name: 'Зона C', capacity: '45%', items: 76 },
      { id: 4, name: 'Зона D', capacity: '90%', items: 180 },
    ]
  },
  {
    id: 2,
    name: 'Региональный склад',
    location: 'Санкт-Петербург',
    totalCapacity: 800,
    usedCapacity: 400,
    activeZones: 3,
    totalZones: 4,
    zones: [
      { id: 1, name: 'Зона A', capacity: '55%', items: 120 },
      { id: 2, name: 'Зона B', capacity: '40%', items: 85 },
      { id: 3, name: 'Зона C', capacity: '65%', items: 95 },
    ]
  }
];