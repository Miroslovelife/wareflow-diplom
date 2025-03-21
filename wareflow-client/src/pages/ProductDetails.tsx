import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { api } from '../utils/api';
import { useAuth } from "../contexts/AuthProvider.tsx";

interface ProductDetail {
  uuid: string;
  title: string;
  count: number;
  qr_path: string;
  description: string;
  zone_id: number;
}

export default function ProductDetailPage() {
  const { role, permissions, isAuthenticated, getPermissionsForWarehouse, username } = useAuth();
  const { warehouseId, zoneId, productId } = useParams();
  const [product, setProduct] = useState<ProductDetail | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [permissionsLoaded, setPermissionsLoaded] = useState(false); // Флаг загрузки прав

  useEffect(() => {
    const fetchProductDetail = async () => {
      if (!warehouseId || !zoneId || !productId || !role) {
        console.error("Ошибка: один из параметров undefined!");
        setError("Ошибка: недостающие параметры в URL");
        setIsLoading(false);
        return;
      }

      try {
        setIsLoading(true);

        const url =
            role === "employer"
                ? `/api/v1/${role}/warehouse/${warehouseId}/zone/${zoneId}/product/product_manage/${productId}`
                : `/api/v1/${role}/warehouse/${warehouseId}/zone/${zoneId}/product/${productId}`;

        console.log("Отправляем запрос:", url);
        const response = await api.get(url);
        console.log("Ответ от API:", response.data);

        setProduct(response.data);
      } catch (err) {
        setError("Ошибка загрузки данных о товаре");
        console.error("Ошибка запроса:", err);
      } finally {
        setIsLoading(false);
      }
    };

    // Если роль - employer и права еще не загружены, загружаем их
    if (role === "employer" && !permissionsLoaded && username) {
      getPermissionsForWarehouse(warehouseId!, username).then(() => {
        setPermissionsLoaded(true);
      });
    }

    // Вызываем `fetchProductDetail`, только если:
    // - роль загружена
    // - либо это `owner`, либо у `employer` загружены права
    if (role && (role === "owner" || permissionsLoaded)) {
      fetchProductDetail();
    }
  }, [warehouseId, zoneId, productId, role, permissionsLoaded, username]); // Используем permissionsLoaded вместо permissions.length

  if (isLoading) return <div className="text-center py-6">Загрузка...</div>;
  if (error) return <div className="text-center py-6 text-red-600">{error}</div>;

  // Извлекаем имя файла без префикса './qr_storage/'
  const imageFileName = product?.qr_path ? product.qr_path.replace('./qr_storage/', '') : '';
  const imagePath = imageFileName ? `https://bebradomen.twc1.net:8443/qr_storage/${imageFileName}` : '';

  return (
      <div className="min-h-screen bg-gray-50 py-8">
        <div className="max-w-7xl mx-auto px-4">
          {product ? (
              <div className="bg-white p-6 rounded-lg shadow-lg">
                <h2 className="text-3xl font-bold mb-4">{product.title}</h2>
                <p className="text-gray-600">Описание: {product.description}</p>
                <p className="text-gray-600">Количество: {product.count}</p>
                <p className="text-gray-600">Зона ID: {product.zone_id}</p>
                {imagePath && (
                    <div className="mt-4">
                      <img src={imagePath} alt="QR-код" />
                    </div>
                )}
              </div>
          ) : (
              <p className="text-center text-gray-600">Данные о товаре не найдены.</p>
          )}
        </div>
      </div>
  );
}
