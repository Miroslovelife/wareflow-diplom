import { useState } from 'react';
import { QRCodeSVG } from 'qrcode.react';
import { Download, Search } from 'lucide-react';

interface Product {
  id: string;
  name: string;
  category: string;
  quantity: number;
  location: string;
}

export default function Products() {
  const [searchTerm, setSearchTerm] = useState('');
  const [selectedProduct, setSelectedProduct] = useState<Product | null>(null);

  const products: Product[] = [
    { id: '001', name: 'Товар A', category: 'Категория 1', quantity: 100, location: 'Зона A' },
    { id: '002', name: 'Товар B', category: 'Категория 2', quantity: 75, location: 'Зона B' },
    { id: '003', name: 'Товар C', category: 'Категория 1', quantity: 50, location: 'Зона C' },
  ];

  const downloadQRCode = (productId: string) => {
    const canvas = document.getElementById('qr-code') as HTMLCanvasElement;
    if (canvas) {
      const pngUrl = canvas
        .toDataURL('image/png')
        .replace('image/png', 'image/octet-stream');
      const downloadLink = document.createElement('a');
      downloadLink.href = pngUrl;
      downloadLink.download = `product-${productId}.png`;
      document.body.appendChild(downloadLink);
      downloadLink.click();
      document.body.removeChild(downloadLink);
    }
  };

  const filteredProducts = products.filter(product =>
    product.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    product.category.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <div className="min-h-screen bg-gray-50 py-8">
      <div className="max-w-7xl mx-auto px-4">
        <div className="mb-8">
          <div className="relative">
            <input
              type="text"
              placeholder="Поиск товаров..."
              className="w-full pl-10 pr-4 py-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-indigo-500"
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
            />
            <Search className="absolute left-3 top-2.5 h-5 w-5 text-gray-400" />
          </div>
        </div>

        <div className="grid md:grid-cols-2 gap-6">
          <div className="bg-white rounded-lg shadow-md p-6">
            <h2 className="text-xl font-bold mb-4">Список товаров</h2>
            <div className="space-y-4">
              {filteredProducts.map((product) => (
                <div
                  key={product.id}
                  className="border rounded-lg p-4 cursor-pointer hover:bg-gray-50"
                  onClick={() => setSelectedProduct(product)}
                >
                  <h3 className="font-semibold">{product.name}</h3>
                  <p className="text-sm text-gray-600">Категория: {product.category}</p>
                  <p className="text-sm text-gray-600">Количество: {product.quantity}</p>
                  <p className="text-sm text-gray-600">Расположение: {product.location}</p>
                </div>
              ))}
            </div>
          </div>

          <div className="bg-white rounded-lg shadow-md p-6">
            <h2 className="text-xl font-bold mb-4">Информация о товаре</h2>
            {selectedProduct ? (
              <div>
                <div className="mb-4">
                  <h3 className="font-semibold text-lg mb-2">{selectedProduct.name}</h3>
                  <p className="text-gray-600">ID: {selectedProduct.id}</p>
                  <p className="text-gray-600">Категория: {selectedProduct.category}</p>
                  <p className="text-gray-600">Количество: {selectedProduct.quantity}</p>
                  <p className="text-gray-600">Расположение: {selectedProduct.location}</p>
                </div>
                <div className="flex flex-col items-center">
                  <QRCodeSVG
                    id="qr-code"
                    value={JSON.stringify(selectedProduct)}
                    size={200}
                    level="H"
                  />
                  <button
                    onClick={() => downloadQRCode(selectedProduct.id)}
                    className="mt-4 flex items-center px-4 py-2 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700"
                  >
                    <Download className="h-5 w-5 mr-2" />
                    Скачать QR-код
                  </button>
                </div>
              </div>
            ) : (
              <p className="text-gray-500 text-center">Выберите товар для просмотра информации</p>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}