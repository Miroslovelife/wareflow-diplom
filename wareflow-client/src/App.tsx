import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Navbar from './components/Navbar';
import Home from './pages/Home';
import Warehouses from './pages/Warehouses';
import WarehouseDetails from './pages/WarehouseDetails';
import Products from './pages/Products';

export default function App() {
  return (
    <Router>
      <div className="min-h-screen bg-gray-100">
        <Navbar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/warehouses" element={<Warehouses />} />
          <Route path="/warehouse/:id" element={<WarehouseDetails />} />
          <Route path="/products" element={<Products />} />
        </Routes>
      </div>
    </Router>
  );
}