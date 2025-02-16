export interface Zone {
  id: number;
  name: string;
  capacity: string;
  items: number;
}

export interface Warehouse {
  id: number;
  name: string;
  location: string;
  totalCapacity: number;
  usedCapacity: number;
  zones: Zone[];
  activeZones: number;
  totalZones: number;
}