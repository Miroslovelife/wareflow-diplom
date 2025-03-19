package usecase

import (
	"fmt"
	"github.com/Miroslovelife/whareflow/internal/config"
	delivery "github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/model"
	"github.com/Miroslovelife/whareflow/internal/domain"
	"github.com/Miroslovelife/whareflow/internal/repositories"
	"github.com/Miroslovelife/whareflow/pkg/qr"
)

type ProductUsecase interface {
	CreateProduct(in *delivery.ProductModelRequest, userId string, warehouseId int, zoneId uint64) error
	FindProduct(userId, productId string) (*delivery.ProductModelResponse, error)
	FindAllProductFromZone(userId string, zoneId int) (*[]delivery.ProductModelResponse, error)
	FindAllProductFromWarehouse(userId string, warehouseId int) (*[]delivery.ProductModelResponse, error)
	UpdateProduct(in *delivery.ProductModelRequest, warehouseId int, productId, userId string) error
	//DeleteProduct(in *delivery.ProductModelRequest, userId string, warehouseId int) error
}

type IProductUsecase struct {
	productRepository repositories.ProductRepository
	qrGenerator       qr.GeneratorQR
	cfg               config.Config
}

func NewIProductUsecase(productRepository repositories.ProductRepository, qrGenerator qr.GeneratorQR, cfg config.Config) *IProductUsecase {
	return &IProductUsecase{
		productRepository: productRepository,
		qrGenerator:       qrGenerator,
		cfg:               cfg,
	}
}

func (pu *IProductUsecase) CreateProduct(in *delivery.ProductModelRequest, userId string, warehouseId int, zoneId uint64) error {
	product := &domain.Product{
		Title:       in.Title,
		Count:       in.Count,
		QrPath:      "",
		Description: in.Description,
		ZoneId:      zoneId,
	}

	createdProduct, err := pu.productRepository.InsertProductData(product, userId, warehouseId)
	if err != nil {
		return err
	}

	qrData := fmt.Sprintf("%s%s", pu.cfg.QR.UrlFrontend, createdProduct.Uuid)

	pathToFle, err := pu.qrGenerator.Generate(qrData, pu.cfg.QR.PathToFile, fmt.Sprintf("%s.png", string(createdProduct.Uuid)))
	if err != nil {
		return err
	}

	createdProduct.QrPath = fmt.Sprintf("./%s", pathToFle)
	fmt.Println(createdProduct)

	errUpdate := pu.productRepository.UpdateProductData(createdProduct, userId, warehouseId)
	if errUpdate != nil {
		return err
	}

	return nil
}

func (pu *IProductUsecase) FindProduct(userId, productId string) (*delivery.ProductModelResponse, error) {
	product, err := pu.productRepository.FindProductData(userId, productId)
	if err != nil {
		return nil, err
	}

	fmt.Println(product)


	productResponse := delivery.ProductModelResponse{
		Uuid:        string(product.Uuid),
		Title:       product.Title,
		Count:       product.Count,
		QrImage:     product.QrPath,
		Description: product.Description,
		ZoneId:      product.ZoneId,
	}

	return &productResponse, nil

}

func (pu *IProductUsecase) FindAllProductFromZone(userId string, zoneId int) (*[]delivery.ProductModelResponse, error) {
	products, err := pu.productRepository.FindAllProductFromZoneData(userId, zoneId)
	if err != nil {
		return nil, err
	}

	var productsRepo []delivery.ProductModelResponse
	fmt.Println(products, "wafaf")
	for _, product := range *products {
		productRepo := delivery.ProductModelResponse{
			Uuid:        string(product.Uuid),
			Title:       product.Title,
			Description: product.Description,
			QrImage:     product.QrPath,
			ZoneId:      product.ZoneId,
		}

		productsRepo = append(productsRepo, productRepo)
	}

	return &productsRepo, nil
}

func (pu *IProductUsecase) FindAllProductFromWarehouse(userId string, warehouseId int) (*[]delivery.ProductModelResponse, error) {
	products, err := pu.productRepository.FindAllProductFromWarehouseData(userId, warehouseId)
	if err != nil {
		return nil, err
	}

	var productsRepo []delivery.ProductModelResponse
	for _, product := range *products {
		productRepo := delivery.ProductModelResponse{
			Uuid:        string(product.Uuid),
			Title:       product.Title,
			Description: product.Description,
			QrImage:     product.QrPath,
			ZoneId:      product.ZoneId,
		}

		productsRepo = append(productsRepo, productRepo)
	}

	return &productsRepo, nil
}

func (pu *IProductUsecase) UpdateProduct(in *delivery.ProductModelRequest, warehouseId int, productId, userId string) error {
	product, err := pu.productRepository.FindProductData(userId, productId)
	if err != nil {
		return err
	}

	product = &domain.Product{
		Uuid:        product.Uuid,
		Title:       in.Title,
		Count:       in.Count,
		QrPath:      product.QrPath,
		Description: in.Description,
		ZoneId:      product.ZoneId,
	}

	errUpdate := pu.productRepository.UpdateProductData(product, userId, warehouseId)
	if errUpdate != nil {
		return err
	}

	return nil

}
