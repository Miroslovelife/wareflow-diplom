import { Warehouse, Package, Home } from 'lucide-react';
import { Link } from 'react-router-dom';

export default function Navbar() {
  return (
    <nav className="bg-indigo-600 text-white">
      <div className="max-w-7xl mx-auto px-4">
        <div className="flex items-center justify-between h-16">
          <div className="flex items-center space-x-8">
            <Link to="/" className="flex items-center space-x-2 font-bold text-xl">
              <Warehouse className="h-6 w-6" />
              <span>WareFlow</span>
            </Link>
            
            <div className="flex space-x-4">
              <Link to="/" className="flex items-center space-x-1 hover:text-indigo-200">
                <Home className="h-5 w-5" />
                <span>Главная</span>
              </Link>
              <Link to="/warehouses" className="flex items-center space-x-1 hover:text-indigo-200">
                <Warehouse className="h-5 w-5" />
                <span>Склады</span>
              </Link>
              <Link to="/products" className="flex items-center space-x-1 hover:text-indigo-200">
                <Package className="h-5 w-5" />
                <span>Товары</span>
              </Link>
            </div>
          </div>
        </div>
      </div>
    </nav>
  );
}