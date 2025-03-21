import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { AuthProvider } from './contexts/AuthProvider.tsx';
import { Login } from './pages/Login';
import { Dashboard } from './pages/Dashboard';
import { Home } from './pages/Home';
import { ProtectedRoute } from './components/ProtectedRoute';
import Navbar from "./components/Navbar.tsx";
import Footer from "./components/Footer.tsx";
import Warehouses from "./pages/Warehouses.tsx";
import Register from "./pages/Register.tsx";
import WarehouseDetails from "./pages/WarehouseDetails.tsx";
import {Scanner} from "./pages/QrScaner.tsx";
import ZoneDetails from "./pages/ZoneDetails.tsx";
import ProductDetails from "./pages/ProductDetails.tsx";
import Employers from "./pages/Employers.tsx";
import ProductManage from "./pages/ProductsManage.tsx";
import Profile from "./pages/Profile.tsx";
function App() {
    return (

        <AuthProvider>
            <BrowserRouter>
                <div className="flex flex-col min-h-screen">
                    <Navbar/>
                    <main className="flex-grow">
                        <Routes>
                            <Route path="/login" element={<Login/>}/>
                            <Route path="/reg" element={<Register/>}/>
                            <Route
                                path="/dashboard"
                                element={
                                    <ProtectedRoute>
                                        <Dashboard/>
                                    </ProtectedRoute>
                                }
                            />
                            <Route
                                path="/"
                                element={
                                    <Home/>
                                }
                            />
                            <Route
                                path="/profile"
                                element={
                                    <ProtectedRoute>
                                        <Profile/>
                                    </ProtectedRoute>
                                }
                            />
                            <Route
                                path="/warehouses"
                                element={
                                    <ProtectedRoute>
                                        <Warehouses/>
                                    </ProtectedRoute>
                                }
                            />
                            <Route
                                path="/warehouses/:warehouseId" // <-- Добавлен маршрут для отдельного склада
                                element={
                                    <ProtectedRoute>
                                        <WarehouseDetails/>
                                    </ProtectedRoute>
                                }
                            />
                            <Route
                                path="warehouses/:warehouseId/zones/:zoneId" // <-- Добавлен маршрут для отдельного склада
                                element={
                                    <ProtectedRoute>
                                        <ZoneDetails/>
                                    </ProtectedRoute>
                                }
                            />
                            <Route
                                path="/scanner" // <-- Добавлен маршрут для отдельного склада
                                element={
                                    <ProtectedRoute>
                                        <Scanner/>
                                    </ProtectedRoute>
                                }
                            />
                            <Route
                                path="/warehouses/:warehouseId/zone/:zoneId" // <-- Добавлен маршрут для отдельного склада
                                element={
                                    <ProtectedRoute>
                                        <ZoneDetails/>
                                    </ProtectedRoute>
                                }
                            />
                            <Route
                                path="/:zoneId/:warehouseId/products/:productId" // <-- Добавлен маршрут для отдельного склада
                                element={
                                    <ProtectedRoute>
                                        <ProductDetails/>
                                    </ProtectedRoute>
                                }
                            />
                            <Route
                                path="/employees/manage/:warehouseId" // <-- Добавлен маршрут для отдельного склада
                                element={
                                    <ProtectedRoute>
                                        <Employers/>
                                    </ProtectedRoute>
                                }
                            />
                            <Route
                                path="/products/manage/:warehouseId" // <-- Добавлен маршрут для отдельного склада
                                element={
                                    <ProtectedRoute>
                                        <ProductManage/>
                                    </ProtectedRoute>
                                }
                            />
                        </Routes>

                    </main>
                    <Footer/>
                </div>

            </BrowserRouter>
        </AuthProvider>
    );
}

export default App;
