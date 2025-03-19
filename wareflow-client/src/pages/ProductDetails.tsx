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
  const { role, isAuthenticated } = useAuth();
  const { productId } = useParams();  // Извлекаем ID товара из URL
  const [product, setProduct] = useState<ProductDetail | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchProductDetail = async () => {
      if (!productId) {
        setError('Ошибка: не указан ID товара');
        setIsLoading(false);
        return;
      }

      try {
        const response = await api.get(`/api/v1/${role}/product/${productId}`);
        setProduct(response.data);
      } catch (err) {
        setError('Ошибка загрузки данных о товаре');
        console.error(err);
      } finally {
        setIsLoading(false);
      }
    };

    fetchProductDetail();
  }, [productId]);

  if (isLoading) return <div className="text-center py-6">Загрузка...</div>;
  if (error) return <div className="text-center py-6 text-red-600">{error}</div>;

  // Извлекаем имя файла без префикса './qr_storage/'
  const imageFileName = product?.qr_path ? product.qr_path.replace('./qr_storage/', '') : '';
  const imagePath = imageFileName ? `http://localhost:8089/qr_storage/${imageFileName}` : '';

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
