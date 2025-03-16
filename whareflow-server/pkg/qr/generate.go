package qr

import (
	"encoding/base64"
	"fmt"
	"github.com/skip2/go-qrcode"
	"log/slog"
	"os"
	"path/filepath"
)

type GeneratorQR interface {
	Generate(data interface{}, outputDir string, fileName string) (string, error)
	DecodeToBase64(qrPath string) (string, error)
}

type Generator struct {
	logger slog.Logger
}

func NewGenerator(logger slog.Logger) *Generator {
	return &Generator{
		logger: logger,
	}
}

func (g *Generator) Generate(data interface{}, outputDir string, fileName string) (string, error) {
	// Преобразуем данные в строку
	dataStr := fmt.Sprintf("%+v", data)

	// Определяем имя файла
	outputPath := filepath.Join(outputDir, fileName)

	// Генерируем QR-код
	err := qrcode.WriteFile(dataStr, qrcode.Medium, 256, outputPath)
	if err != nil {
		g.logger.Error(fmt.Sprintf("error: %s", err))
		return "", fmt.Errorf("failed to generate QR code: %w", err)
	}

	return outputPath, nil
}

func (g *Generator) DecodeToBase64(qrPath string) (string, error) {
	qrData, err := os.ReadFile(qrPath)
	if err != nil {
		return "", err
	}

	QrImageDecde := "data:image/png;base64," + base64.StdEncoding.EncodeToString(qrData)

	return QrImageDecde, nil
}
