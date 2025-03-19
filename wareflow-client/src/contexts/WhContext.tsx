// contexts/WarehouseContext.tsx
import { createContext, useContext, useState, ReactNode } from 'react';

interface WarehouseContextType {
    warehouseId: number | null;
    setWarehouseId: (id: number) => void;
}

const WarehouseContext = createContext<WarehouseContextType | undefined>(undefined);

export const WarehouseProvider = ({ children }: { children: ReactNode }) => {
    const [warehouseId, setWarehouseId] = useState<number | null>(null);

    return (
        <WarehouseContext.Provider value={{ warehouseId, setWarehouseId }}>
            {children}
        </WarehouseContext.Provider>
    );
};

export const useWarehouse = () => {
    const context = useContext(WarehouseContext);
    if (!context) {
        throw new Error('useWarehouse must be used within a WarehouseProvider');
    }
    return context;
};
