import { useEffect, useState } from 'react';
import { api } from '../utils/api';
import { UserPlus, Trash2, Edit } from 'lucide-react';
import { useAuth } from '../contexts/AuthProvider';
import { useParams } from "react-router-dom";

interface Employer {
    phone_number: string;
    username: string;
    first_name: string;
    last_name: string;
    surname: string;
    email: string;
}
interface Warehouse {
    id: number;
    name: string;
    address: string;
}


export default function WarehouseEmployees() {
    const { role, isAuthenticated } = useAuth();
    const [employees, setEmployees] = useState<Employer[]>([]);
    const [newEmployeeUsername, setNewEmployeeUsername] = useState('');
    const [roleName, setRoleName] = useState('');
    const [selectedPermissions, setSelectedPermissions] = useState<number[]>([]);
    const [systemPermissions, setSystemPermissions] = useState<{ id: number; name: string }[]>([]);
    const [loading, setLoading] = useState(false);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [isEditModalOpen, setIsEditModalOpen] = useState(false);
    const [editRoleName, setEditRoleName] = useState('');
    const { warehouseId } = useParams();
    const [ warehouse, setWarehouse] = useState<Warehouse[]>([]);
    // Получение прав для системы
    useEffect(() => {
        const fetchPermissions = async () => {
            if (!isAuthenticated || !role) return;

            try {
                const response = await api.get(`/api/v1/${role}/permission`);
                setSystemPermissions(response.data || []);
            } catch (error) {
                console.error('Ошибка получения прав:', error);
            }
        };

        fetchPermissions();
    }, [isAuthenticated, role]);


    useEffect(() => {
        fetchEmployees();
        getWH();
    }, []);

    const fetchEmployees = async () => {
        setLoading(true);
        try {
            const endpoint = role === 'owner'
                ? `/api/v1/${role}/warehouse/${warehouseId}/employer`
                : `/api/v1/${role}/warehouse/role/${warehouseId}/role_manage/employer`;
            const response = await api.get(endpoint);
            setEmployees(response.data);
        } catch (error) {
            console.error('Ошибка при получении сотрудников:', error);
        }
        setLoading(false);
    };


    const getWH = async () => {

        try {
            const endpoint = role === 'owner'
                ? `/api/v1/${role}/warehouse/${warehouseId}`
                : `/api/v1/${role}/global/${warehouseId}/warehouse_manage`;

            const response =  await api.get(endpoint);
            setWarehouse(response.data) // Обновляем список сотрудников
        } catch (error) {
            console.error('Ошибка при добавлении сотрудника:', error);
        }
    };

    const addEmployee = async () => {
        if (!newEmployeeUsername || !roleName || selectedPermissions.length === 0) return;

        const roleRequest = {
            name: roleName,
            username: newEmployeeUsername,
            permissions: selectedPermissions
        };

        try {
            const endpoint = role === 'owner'
                ? `/api/v1/owner/role/${warehouseId}`
                : `/api/v1/employer/warehouse/role/${warehouseId}/role_manage`;

            await api.post(endpoint, roleRequest);
            setIsModalOpen(false);
            fetchEmployees(); // Обновляем список сотрудников
        } catch (error) {
            console.error('Ошибка при добавлении сотрудника:', error);
        }
    };

    // Удалить сотрудника
    const removeEmployee = async (username: string) => {
        try {
            await api.delete(`/api/v1/employees/${username}`);
            fetchEmployees();
        } catch (error) {
            console.error('Ошибка при удалении сотрудника:', error);
        }
    };

    // Изменить роль сотрудника
    const editEmployeeRole = async (username: string) => {
        if (!editRoleName) return;

        try {
            const endpoint = `/api/v1/employees/${username}/role`;
            await api.put(endpoint, { role: editRoleName });
            setIsEditModalOpen(false);
            fetchEmployees();
        } catch (error) {
            console.error('Ошибка при изменении роли сотрудника:', error);
        }
    };

    // Открыть/закрыть модальное окно для добавления сотрудника
    const toggleModal = () => {
        setIsModalOpen(!isModalOpen);
    };

    // Открыть/закрыть модальное окно для редактирования роли
    const toggleEditModal = (username: string) => {
        setIsEditModalOpen(!isEditModalOpen);
        setEditRoleName(''); // Сбрасываем значение роли при открытии
    };

    return (
        <div className="max-w-4xl mx-auto p-6 bg-white shadow-lg rounded-lg">
            <h2 className="text-2xl font-semibold mb-4">Управление сотрудниками склада {warehouse.name}</h2>

            {/* Кнопка для открытия модального окна добавления сотрудника */}
            <button
                onClick={toggleModal}
                className="flex items-center bg-indigo-600 text-white px-4 py-2 rounded-lg shadow hover:bg-indigo-700 mb-6"
            >
                <UserPlus className="h-5 w-5 mr-2" /> Добавить сотрудника
            </button>

            {/* Модальное окно для добавления сотрудника */}
            {isModalOpen && (
                <div className="fixed inset-0 flex justify-center items-center bg-black bg-opacity-50">
                    <div className="bg-white p-6 rounded-lg shadow-lg max-w-md w-full">
                        <h3 className="text-xl font-semibold mb-4">Добавить сотрудника</h3>

                        <div className="mb-4">
                            <input
                                type="text"
                                placeholder="Username сотрудника"
                                value={newEmployeeUsername}
                                onChange={(e) => setNewEmployeeUsername(e.target.value)}
                                className="border p-2 rounded w-full"
                            />
                        </div>
                        <div className="mb-4">
                            <input
                                type="text"
                                placeholder="Название роли"
                                value={roleName}
                                onChange={(e) => setRoleName(e.target.value)}
                                className="border p-2 rounded w-full"
                            />
                        </div>

                        <div className="mb-4">
                            <h4 className="font-semibold">Разрешения</h4>

                            {systemPermissions.length === 0 ? (
                                    <div className="col-span-2 text-center text-gray-600">
                                        Ошибка при загрузки системных прав.
                                    </div>) : (
                                systemPermissions.map((permission) => (
                                <label key={permission.id} className="block">
                                    <input
                                        type="checkbox"
                                        value={permission.id}
                                        onChange={(e) => {
                                            const selected = [...selectedPermissions];
                                            if (e.target.checked) {
                                                selected.push(permission.id);
                                            } else {
                                                const index = selected.indexOf(permission.id);
                                                if (index !== -1) {
                                                    selected.splice(index, 1);
                                                }
                                            }
                                            setSelectedPermissions(selected);
                                        }}
                                    />
                                    {permission.name}
                                </label>
                            )))}
                        </div>

                        <div className="flex justify-end">
                            <button
                                onClick={addEmployee}
                                className="bg-blue-500 text-white px-4 py-2 rounded mr-2"
                            >
                                Добавить
                            </button>
                            <button
                                onClick={toggleModal}
                                className="bg-gray-500 text-white px-4 py-2 rounded"
                            >
                                Закрыть
                            </button>
                        </div>
                    </div>
                </div>
            )}

            {/* Модальное окно для редактирования роли сотрудника */}
            {isEditModalOpen && (
                <div className="fixed inset-0 flex justify-center items-center bg-black bg-opacity-50">
                    <div className="bg-white p-6 rounded-lg shadow-lg max-w-md w-full">
                        <h3 className="text-xl font-semibold mb-4">Изменить роль сотрудника</h3>

                        <div className="mb-4">
                            <input
                                type="text"
                                placeholder="Новая роль"
                                value={editRoleName}
                                onChange={(e) => setEditRoleName(e.target.value)}
                                className="border p-2 rounded w-full"
                            />
                        </div>

                        <div className="flex justify-end">
                            <button
                                onClick={editEmployeeRole}
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

            {/* Таблица сотрудников */}
            {loading ? (
                <p>Загрузка...</p>
            ) : (
                <table className="w-full border-collapse border border-gray-300">
                    <thead>
                    <tr className="bg-gray-100">
                        <th className="border p-2">Имя</th>
                        <th className="border p-2">Username</th>
                        <th className="border p-2">Email</th>
                        <th className="border p-2">Телефон</th>
                        <th className="border p-2">Действия</th>
                    </tr>
                    </thead>
                    <tbody>
                    {
                        employees === null ? (
                                <div className="col-span-2 text-gray-600 flex justify-center">
                                    Сотрудников на складе нет.
                                </div>) : (
                        employees.map((employee) => (
                        <tr key={employee.username} className="text-center">
                            <td className="border p-2">{employee.first_name} {employee.last_name}</td>
                            <td className="border p-2">{employee.username}</td>
                            <td className="border p-2">{employee.email}</td>
                            <td className="border p-2">{employee.phone_number}</td>
                            <td className="border p-2">
                                <button
                                    onClick={() => toggleEditModal(employee.username)}
                                    className="bg-yellow-500 text-white px-3 py-1 rounded flex items-center mr-2"
                                >
                                    <Edit className="h-5 w-5 mr-1" /> Изменить роль
                                </button>
                                <button
                                    onClick={() => removeEmployee(employee.username)}
                                    className="bg-red-500 text-white px-3 py-1 rounded flex items-center"
                                >
                                    <Trash2 className="h-5 w-5 mr-1" /> Удалить
                                </button>
                            </td>
                        </tr>
                    )))}
                    </tbody>
                </table>
            )}
        </div>
    );
}
