import { Html5QrcodeScanner, Html5QrcodeScanType } from 'html5-qrcode';
import { useEffect, useRef, useState } from 'react';

export const Scanner: React.FC = () => {
    const [scanResult, setScanResult] = useState<string | null>(null);
    const scannerRef = useRef<Html5QrcodeScanner | null>(null); // Используем useRef для хранения сканера

    useEffect(() => {
        if (!scannerRef.current) { // Проверяем, не был ли сканер уже создан
            scannerRef.current = new Html5QrcodeScanner(
                'reader',
                {
                    fps: 10,
                    qrbox: { width: 240, height: 240 },
                    rememberLastUsedCamera: true,
                    supportedScanTypes: [Html5QrcodeScanType.SCAN_TYPE_CAMERA],
                },
                false
            );

            scannerRef.current.render(
                (decodedText) => {
                    setScanResult(decodedText);
                    scannerRef.current?.clear();

                    if (isValidURL(decodedText)) {
                        window.location.href = decodedText;
                    }
                },
                (errorMessage) => console.warn(errorMessage)
            );
        }

        return () => {
            scannerRef.current?.clear();
            scannerRef.current = null; // Очищаем референс, чтобы избежать повторного создания
        };
    }, []);

    const restartScan = () => {
        setScanResult(null);
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

            <div className="w-[17rem] h-[17rem] bg-white shadow-lg rounded-lg flex items-center justify-center relative">
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
