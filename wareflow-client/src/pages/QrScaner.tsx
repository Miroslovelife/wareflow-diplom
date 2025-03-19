import { Html5QrcodeScanner } from 'html5-qrcode';
import { useEffect, useState } from 'react';

export const Scanner: React.FC = () => {
    const [scanResult, setScanResult] = useState<string | null>(null);
    const [scanner, setScanner] = useState<Html5QrcodeScanner | null>(null);

    useEffect(() => {
        // Инициализируем сканер при монтировании компонента
        const newScanner = new Html5QrcodeScanner('reader', { fps: 10, qrbox: 350 });

        newScanner.render(
            (decodedText) => {
                setScanResult(decodedText);
                newScanner.clear();

                // Проверяем, является ли результат ссылкой
                if (isValidURL(decodedText)) {
                    window.location.href = decodedText; // Автоматический переход
                }
            },
            (errorMessage) => console.warn(errorMessage)
        );
        setScanner(newScanner);

        // Очищаем сканер при размонтировании компонента
        return () => {
            newScanner.clear();
            setScanner(null); // Обнуляем сканер, чтобы избежать попыток доступа к нему после размонтирования
        };
    }, []); // Эффект запускается только один раз при монтировании компонента

    const restartScan = () => {
        setScanResult(null); // Сброс результата
    };

    const isValidURL = (text: string) => {
        try {
            new URL(text);
            return true;
        } catch {
            return false;
        }
    };

    return (
        <div className="min-h-screen flex flex-col items-center justify-center bg-gray-100 p-6">
            <h1 className="text-3xl font-bold mb-6">QR Сканер</h1>

            <div className="w-96 h-96 bg-white shadow-lg rounded-lg flex items-center justify-center relative">
                {!scanResult ? (
                    <div id="reader" className="w-full h-full absolute"></div>
                ) : (
                    <p className="text-green-600 font-semibold text-lg p-4">
                        Отсканировано: {scanResult}
                    </p>
                )}
            </div>

            <div className="mt-6 flex gap-4">
                {scanResult && (
                    <button
                        onClick={restartScan}
                        className="px-6 py-3 bg-blue-500 text-white font-semibold rounded-lg shadow-md hover:bg-blue-600 transition"
                    >
                        🔄 Пересканировать
                    </button>
                )}
            </div>
        </div>
    );
};
