import { useEffect, useState } from 'react';
import { useParams, Link } from 'react-router-dom';
import { api } from '../utils/api';
import { useAuth } from '../contexts/AuthProvider';
import WarehouseStats from '../components/warehouse/WarehouseStats';
import ZoneList from '../components/warehouse/ZoneList';

interface Warehouse {
    id: number;
    name: string;
    address: string;
}

interface Zone {
    id: number;
    name: string;
    capacity: number;
}

interface Employer {
    phone_number: string;
    username: string;
    first_name: string;
    last_name: string;
    surname: string;
    email: string;
}

interface Product {
    uuid: string;
    title: string;
    count: number;
    zone_id: number;
}

export default function WarehouseDetails() {
    const { role, permissions, isAuthenticated, getPermissionsForWarehouse, username } = useAuth();
    const { warehouseId } = useParams();
    const [warehouse, setWarehouse] = useState<Warehouse | null>(null);
    const [zones, setZones] = useState<Zone[] | null>(null);
    const [employers, setEmployers] = useState<Employer[] | null>(null);
    const [products, setProducts] = useState<Product[] | null>(null);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const [showProducts, setShowProducts] = useState(false); // Управление видимостью товаров
    const [showPermissions, setShowPermissions] = useState(false); // Управление видимостью разрешений на склад

    const permissionTranslations: { [key: string]: string } = {
        "create_role": "Создание роли",
        "change_role": "Изменение роли",
        "get_user_role": "Получение роли пользователя",
        "get_all_user_role": "Получение всех ролей пользователей",
        "remove_user_role": "Удаление роли пользователя",
        "zone_manage": "Управление зонами",
        "product_manage": "Управление товарами",
        "role_manage": "Управление ролями",
        "get_my_permissions": "Получение моих разрешений",
        "warehouse_manage": "Управление складом"
    };

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

                if (role === 'owner' || permissions.some(permission => permission.name === 'warehouse_manage')) {
                    const warehouseResponse = await api.get(
                        role === 'employer'
                            ? `/api/v1/${role}/warehouse/global/${warehouseId}/warehouse_manage`
                            : `/api/v1/${role}/warehouse/${warehouseId}`
                    );
                    setWarehouse(warehouseResponse.data);
                }

                if (role === 'owner' || permissions.some(permission => permission.name === 'zone_manage')) {
                    const zonesResponse = await api.get(
                        role === 'employer'
                            ? `/api/v1/${role}/warehouse/${warehouseId}/zone/zone_manage`
                            : `/api/v1/${role}/warehouse/${warehouseId}/zone`
                    );
                    setZones(zonesResponse.data?.zones || []);
                }

                if (role === 'owner' || permissions.some(permission => permission.name === 'role_manage')) {
                    const employersResponse = await api.get(
                        role === 'employer'
                            ? `/api/v1/employer/warehouse/role/${warehouseId}/role_manage/employer`
                            : `/api/v1/owner/warehouse/${warehouseId}/employer`
                    );
                    setEmployers(employersResponse.data || []);
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

    const fetchProducts = async () => {
        if (!warehouseId) {
            setError('Ошибка: не указан ID склада');
            return;
        }

        if (role !== 'owner' && !permissions.some(permission => permission.name === 'product_manage')) {
            setError('Ошибка: недостаточно прав для просмотра товаров');
            return;
        }

        try {
            const productsResponse = await api.get(
                role === 'employer'
                    ? `/api/v1/employer/warehouse/${warehouseId}/product/product_manage`
                    : `/api/v1/owner/warehouse/${warehouseId}/product`
            );
            setProducts(productsResponse.data || []);
        } catch (err) {
            setError('Ошибка загрузки данных о товарах');
            console.error(err);
        }
    };

    if (isLoading) return <div className="text-center py-6">Загрузка...</div>;
    if (error) return <div className="text-center py-6 text-red-600">{error}</div>;

    return (
        <div className="min-h-screen bg-gray-50 py-8">
            <div className="max-w-7xl mx-auto px-4">
                {warehouse ? (
                    <>
                        <div className="flex justify-between items-center mb-4">
                            <WarehouseStats warehouse={warehouse} />
                            {role === 'employer' && (
                                <button
                                    onClick={() => setShowPermissions(!showPermissions)}
                                    className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700"
                                >
                                    {showPermissions ? 'Скрыть разрешения' : 'Показать разрешения'}
                                </button>
                            )}
                        </div>

                        {/* Список разрешений */}
                        {showPermissions && role === 'employer' && (
                            <div className="bg-gray-100 p-4 rounded-lg shadow-lg mt-4">
                                <h3 className="text-xl font-semibold mb-4">Разрешения на склад:</h3>
                                <ul>
                                    {permissions.map((permission) => (
                                        <li key={permission.name} className="text-gray-600 mb-2">
                                            {permissionTranslations[permission.name] || permission.name}
                                        </li>
                                    ))}
                                </ul>
                            </div>
                        )}
                    </>
                ) : (
                    <p className="text-center text-gray-600">Данные о складе не найдены.</p>
                )}

                {/* Зоны склада */}
                {(permissions.some(permission => permission.name === 'zone_manage') || role === 'owner') && (
                    <div className="mt-8">
                        <h3 className="text-xl font-semibold mb-4">Зоны склада</h3>
                        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                            {zones?.map((zone) => (
                                <div key={zone.id} className="bg-white p-6 rounded-lg shadow-lg">
                                    <h4 className="text-lg font-semibold mb-4">{zone.name}</h4>
                                    <p className="text-gray-600">Вместимость: {zone.capacity}</p>
                                    <Link
                                        to={`/zones/${zone.id}`}
                                        className="mt-4 inline-block text-indigo-600 hover:text-indigo-800"
                                    >
                                        Подробнее
                                    </Link>
                                </div>
                            ))}
                        </div>
                    </div>
                )}

                {/* Работники */}
                {(permissions.some(permission => permission.name === 'role_manage') || role === 'owner') && (
                    <div className="mt-8">
                        <h3 className="text-xl font-semibold mb-4">
                            Работники склада
                            <span className="text-sm text-gray-500 ml-4">
                                <Link to={`/employees/manage/${warehouseId}`} className="text-indigo-600">Управление сотрудниками</Link>
                            </span>
                        </h3>
                        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                            {employers?.map((employer) => (
                                <div key={employer.username} className="bg-white p-6 rounded-lg shadow-lg">
                                    <h4 className="text-lg font-semibold mb-4">
                                        {employer.first_name} {employer.last_name}
                                    </h4>
                                    <p className="text-gray-600">Телефон: {employer.phone_number}</p>
                                    <p className="text-gray-600">Email: {employer.email}</p>
                                    <p className="text-gray-600">Username: {employer.username}</p>
                                </div>
                            ))}
                        </div>
                    </div>
                )}

                {/* Управление товарами */}
                {(permissions.some(permission => permission.name === 'product_manage') || role === 'owner') && (
                    <div className="mt-8">
                        <h3 className="text-xl font-semibold mb-4">
                            Скрыть товары
                            <span className="text-sm text-gray-500 ml-4">
                                <Link to={`/products/manage/${warehouseId}`} zones={zones} className="text-indigo-600">Управление товарами</Link>
                            </span>
                        </h3>
                        <button
                            onClick={() => {
                                if (!showProducts) {
                                    fetchProducts(); // Загружаем товары при первом показе
                                }
                                setShowProducts(!showProducts);
                            }}
                            className="mt-4 px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700"
                        >
                            {showProducts ? 'Скрыть товары' : 'Показать товары'}
                        </button>

                        {/* Карточки товаров */}
                        {showProducts && products && products.length > 0 && (
                            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 mt-6">
                                {products.map((product) => (
                                    <div key={product.uuid} className="bg-white p-6 rounded-lg shadow-lg">
                                        <h4 className="text-lg font-semibold mb-4">{product.title}</h4>
                                        <p className="text-gray-600">Количество: {product.count}</p>
                                        <p className="text-gray-600">Зона ID: {product.zone_id}</p>
                                        <Link
                                            to={`/products/${product.uuid}`}
                                            className="mt-4 inline-block text-indigo-600 hover:text-indigo-800"
                                        >
                                            Подробнее
                                        </Link>
                                    </div>
                                ))}
                            </div>
                        )}
                    </div>
                )}
            </div>
        </div>
    );
}
