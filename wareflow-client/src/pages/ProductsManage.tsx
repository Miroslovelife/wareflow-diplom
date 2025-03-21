import { useEffect, useState } from 'react';
import { api } from '../utils/api';
import { PlusCircle, Trash2, Edit } from 'lucide-react';
import { useAuth } from '../contexts/AuthProvider';
import {Link, useParams} from "react-router-dom";

interface Product {
    uuid: number;
    title: string;
    count: number;
    description: string;
    zone_id: number;
}

interface Warehouse {
    id: number;
    name: string;
}

interface Zone {
    id: number;
    name: string;
}

export default function ProductManage() {
    const { role, permissions, isAuthenticated, getPermissionsForWarehouse, username } = useAuth();
    const [products, setProducts] = useState<Product[]>([]);
    const [newProductName, setNewProductName] = useState('');
    const [newProductDescription, setNewProductDescription] = useState('');
    const [newProductQuantity, setNewProductQuantity] = useState<number>(0);
    const [newProductZone, setNewProductZone] = useState<number>(0);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [error, setError] = useState<string | null>(null);
    const [isLoading, setIsLoading] = useState(true);
    const [isEditModalOpen, setIsEditModalOpen] = useState(false);
    const [editProductName, setEditProductName] = useState('');
    const [editProductDescription, setEditProductDescription] = useState('');
    const [editProductQuantity, setEditProductQuantity] = useState<number>(0);
    const { warehouseId } = useParams();
    const [warehouses, setWarehouses] = useState<Warehouse[]>([]);
    const [zones, setZones] = useState<Zone[]>([]);
    const [selectedWarehouse, setSelectedWarehouse] = useState<string>('');
    const [selectedZone, setSelectedZone] = useState<number>(0);

    useEffect(() => {
        const fetchWarehouseData = async () => {
            if (!isAuthenticated || !role) {
                setError('Ошибка: пользователь не авторизован или роль не определена.');
                setIsLoading(false);
                return;
            }

            if (!warehouseId) {
                setError('Ошибка: не указан ID склада');
                setIsLoading(false);
                return;
            }

            try {
                if (role === 'employer' && permissions.length === 0) {
                    await getPermissionsForWarehouse(warehouseId, username);
                }

                if (role === 'owner' || permissions.some(permission => permission.name === 'product_manage')) {
                    const warehouseResponse = await api.get(
                        role === 'employer'
                            ? `/api/v1/${role}/warehouse/${warehouseId}/product/product_manage`
                            : `/api/v1/${role}/warehouse/${warehouseId}/product`
                    );
                    setProducts(warehouseResponse.data);
                }
            } catch (err) {
                setError('Ошибка загрузки данных о складе');
                console.error(err);
            } finally {
                setIsLoading(false);
            }
        };

        fetchWarehouseData();
    }, [warehouseId, role, isAuthenticated, permissions]);

    useEffect(() => {
        if (!isAuthenticated || !role) {
            setError('Ошибка: пользователь не авторизован или роль не определена.');
            setIsLoading(false);
            return;
        }

        const fetchWarehouses = async () => {
            try {
                    const warehouseResponse = await api.get(
                        role === 'employer'
                            ? `/api/v1/${role}/warehouse`
                            : `/api/v1/${role}/warehouse`
                    );
                setWarehouses(warehouseResponse.data.warehouses);
                console.log(warehouseResponse.data.warehouses);
            } catch (err) {
                setError('Ошибка загрузки данных о складе');
                console.error(err);
            } finally {
                setIsLoading(false);
            }
        };

        fetchWarehouses();
    }, [role, isAuthenticated]);

    useEffect(() => {
        const fetchZones = async () => {
            if (!isAuthenticated || !role) {
                setError('Ошибка: пользователь не авторизован или роль не определена.');
                setIsLoading(false);
                return;
            }

            if (!selectedWarehouse) return;

            try {
                if (role === 'employer' && permissions.length === 0) {
                    await getPermissionsForWarehouse(warehouseId, username);
                }

                if (role === 'owner' || permissions.some(permission => permission.name === 'zone_manage')) {
                    const warehouseResponse = await api.get(
                        role === 'employer'
                            ? `/api/v1/${role}/warehouse/${warehouseId}/zone/zone_manage`
                            : `/api/v1/${role}/warehouse/${warehouseId}/zone`
                    );
                    setZones(warehouseResponse.data.zones);
                }
            } catch (err) {
                setError('Ошибка загрузки данных о складе');
                console.error(err);
            } finally {
                setIsLoading(false);
            }
        };

        fetchZones();
    }, [selectedWarehouse, role]);


    // Функция для открытия модального окна добавления
    const toggleModal = () => {
        setIsModalOpen(!isModalOpen);
    };

    // Функция для добавления нового продукта
    const addProduct = async () => {
        try {
            if (role === 'owner' || permissions.some(permission => permission.name === 'product_manage')) {
                const newProduct = { title: newProductName,  count: newProductQuantity, description: newProductDescription, zone_id:Number(selectedZone)};
                const warehouseResponse = await api.post(
                    role === 'employer'
                        ? `/api/v1/${role}/warehouse/${warehouseId}/product/product_manage`
                        : `/api/v1/${role}/warehouse/${warehouseId}/product`,
                    newProduct
                );
                setProducts([...products, newProduct]);
                toggleModal();
            }
        } catch (err) {
            setError('Ошибка добавления продукта');
            console.error(err);
        }
    };



    // Функция для редактирования продукта
    const toggleEditModal = (product: Product) => {
        setEditProductName(product.title);
        setEditProductDescription(product.description);
        setEditProductQuantity(product.count);
        setIsEditModalOpen(true);
    };

    // Функция для сохранения изменений продукта
    const editProduct = async (productId: number) => {
        const updatedProduct = {
            name: editProductName,
            description: editProductDescription,
            quantity: editProductQuantity
        };
        try {
            await api.put(`/api/v1/${role}/warehouse/${warehouseId}/product/${productId}`, updatedProduct);
            setProducts(products.map(product => (product.uuid === productId ? { ...product, ...updatedProduct } : product)));
            setIsEditModalOpen(false);
        } catch (err) {
            setError('Ошибка редактирования продукта');
            console.error(err);
        }
    };

    // Функция для удаления продукта
    const removeProduct = async (productId: number) => {
        try {
            await api.delete(`/api/v1/${role}/warehouse/${warehouseId}/product/${productId}`);
            setProducts(products.filter(product => product.uuid !== productId));
        } catch (err) {
            setError('Ошибка удаления продукта');
            console.error(err);
        }
    };

    return (
        <div className="max-w-4xl mx-auto p-6 bg-white shadow-lg rounded-lg mt-10">
            <h2 className="text-2xl font-semibold mb-4">Управление товарами</h2>

            {/* Кнопка для открытия модального окна добавления продукта */}
            {(role === 'owner' || permissions.some(permission => permission.name === 'product_manage')) && (
                <button
                    onClick={toggleModal}
                    className="flex items-center bg-indigo-600 text-white px-4 py-2 rounded-lg shadow hover:bg-indigo-700 mb-6"
                >
                    <PlusCircle className="h-5 w-5 mr-2" /> Добавить товар
                </button>
            )}

            {/* Модальное окно для добавления продукта */}

            {isModalOpen && (
                <div className="fixed inset-0 flex justify-center items-center bg-black bg-opacity-50">
                    <div className="bg-white p-6 rounded-lg shadow-lg max-w-md w-full">
                        <h3 className="text-xl font-semibold mb-4">Добавить товар</h3>

                        <div className="mb-4">
                            <select
                                value={selectedWarehouse}
                                onChange={(e) => setSelectedWarehouse(e.target.value)}
                                className="border p-2 rounded w-full"
                            >
                                <option value="">Выберите склад</option>
                                {warehouses.length === 0 ?  (
                                    <p value="" >
                                        Нет доступных складов.
                                    </p>) : (warehouses.map((warehouse) => (
                                    <option key={warehouse.id} value={warehouse.id}>{warehouse.name}</option>
                                )))}
                            </select>
                        </div>
                        <div className="mb-4">
                            <select
                                value={selectedZone}
                                onChange={(e) => setSelectedZone(Number(e.target.value))}
                                className="border p-2 rounded w-full"
                            >
                                <option value="">Выберите зону</option>
                                {zones === null ?  (
                                    <p value="" >
                                        Нет доступных зон.
                                    </p>
                                ) : ( (zones.map((zone) => (
                                    <option key={zone.id} value={zone.id}>{zone.name}</option>
                                ))))}
                            </select>
                        </div>
                        <div className="mb-4">
                            <input
                                type="text"
                                placeholder="Название продукта"
                                value={newProductName}
                                onChange={(e) => setNewProductName(e.target.value)}
                                className="border p-2 rounded w-full"
                            />
                        </div>
                        <div className="mb-4">
                            <textarea
                                placeholder="Описание продукта"
                                value={newProductDescription}
                                onChange={(e) => setNewProductDescription(e.target.value)}
                                className="border p-2 rounded w-full"
                            />
                        </div>
                        <div className="mb-4">
                            Количество
                            <input
                                type="number"
                                placeholder="Количество"
                                value={newProductQuantity}
                                onChange={(e) => setNewProductQuantity(Number(e.target.value))}
                                className="border p-2 rounded w-full"
                            />
                        </div>
                        <div className="flex justify-end">
                            <button onClick={addProduct} className="bg-blue-500 text-white px-4 py-2 rounded mr-2">
                                Добавить
                            </button>
                            <button onClick={toggleModal} className="bg-gray-500 text-white px-4 py-2 rounded">
                                Закрыть
                            </button>
                        </div>
                    </div>
                </div>
            )}

            {/* Модальное окно для редактирования продукта */}
            {isEditModalOpen && (
                <div className="fixed inset-0 flex justify-center items-center bg-black bg-opacity-50">
                    <div className="bg-white p-6 rounded-lg shadow-lg max-w-md w-full">
                        <h3 className="text-xl font-semibold mb-4">Изменить продукт</h3>

                        <div className="mb-4">
                            <input
                                type="text"
                                placeholder="Название продукта"
                                value={editProductName}
                                onChange={(e) => setEditProductName(e.target.value)}
                                className="border p-2 rounded w-full"
                            />
                        </div>
                        <div className="mb-4">
                            <textarea
                                placeholder="Описание продукта"
                                value={editProductDescription}
                                onChange={(e) => setEditProductDescription(e.target.value)}
                                className="border p-2 rounded w-full"
                            />
                        </div>
                        <div className="mb-4">
                            <input
                                type="number"
                                placeholder="Количество"
                                value={editProductQuantity}
                                onChange={(e) => setEditProductQuantity(Number(e.target.value))}
                                className="border p-2 rounded w-full"
                            />
                        </div>

                        <div className="flex justify-end">
                            <button
                                onClick={() => editProduct(products.find(product => product.title === editProductName)?.uuid || 0)}
                                className="bg-blue-500 text-white px-4 py-2 rounded mr-2"
                            >
                                Сохранить
                            </button>
                            <button
                                onClick={() => setIsEditModalOpen(false)}
                                className="bg-gray-500 text-white px-4 py-2 rounded"
                            >
                                Закрыть
                            </button>
                        </div>
                    </div>
                </div>
            )}

            {/* Таблица продуктов */}
            <table className="w-full border-collapse border border-gray-300">
                <thead>
                <tr className="bg-gray-100">
                    <th className="border p-2">Название</th>
                    <th className="border p-2">Описание</th>
                    <th className="border p-2">Количество</th>
                    <th className="border p-2">Действия</th>
                </tr>
                </thead>
                <tbody>
                {products === null ? (
                    <div className="col-span-2 text-center text-gray-600">
                        Товаров пока нет.
                    </div>
                ) : (
                    products.map((product) => (
                    <tr key={product.uuid} className="text-center">
                        <td className="border p-2">
                            <Link to={`/products/${product.uuid}`} zones={zones} className="text-indigo-600">{product.title}</Link>
                        </td>
                        <td className="border p-2">{product.description}</td>
                        <td className="border p-2">{product.count}</td>
                        <td className="border p-2">
                            {(role === 'owner' || permissions.some(permission => permission.name === 'product_manage')) && (
                                <>
                                    <button
                                        onClick={() => toggleEditModal(product)}
                                        className="bg-yellow-500 text-white px-3 py-1 rounded flex items-center mr-2"
                                    >
                                        <Edit className="h-5 w-5 mr-1" /> Изменить
                                    </button>
                                    <button
                                        onClick={() => removeProduct(product.uuid)}
                                        className="bg-red-500 text-white px-3 py-1 rounded flex items-center"
                                    >
                                        <Trash2 className="h-5 w-5 mr-1" /> Удалить
                                    </button>
                                </>
                            )}
                        </td>
                    </tr>
                )))}
                </tbody>
            </table>
        </div>
    );
}
