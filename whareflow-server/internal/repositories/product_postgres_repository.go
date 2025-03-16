package repositories

import (
	"errors"
	"github.com/Miroslovelife/whareflow/internal/domain"
	custom_errors "github.com/Miroslovelife/whareflow/internal/errors"
	"github.com/Miroslovelife/whareflow/pkg/database"
	"gorm.io/gorm"
	"log/slog"
)

type ProductRepository interface {
	InsertProductData(in *domain.Product, userId string, warehouseId int) (*domain.Product, error)
	UpdateProductData(in *domain.Product, userId string, warehouseId int) error
	DeleteProductData(in *domain.Product, userId string, warehouseId int) error
	FindAllProductFromZoneData(userId string, zoneId int) (*[]domain.Product, error)
	FindAllProductFromWarehouseData(userId string, warehouseId int) (*[]domain.Product, error)
	FindAllProductData(userId string) (*[]domain.Product, error)
	FindProductData(userId string, productId string) (*domain.Product, error)
}

type ProductPostgresRepository struct {
	db     database.Database
	logger slog.Logger
}

func NewProductPostgresRepository(db database.Database, logger slog.Logger) *ProductPostgresRepository {
	return &ProductPostgresRepository{
		db:     db,
		logger: logger,
	}
}

func (pr *ProductPostgresRepository) InsertProductData(in *domain.Product, userId string, warehouseId int) (*domain.Product, error) {
	var warehouse domain.WareHouse
	if err := pr.db.GetDb().Where("id = ? AND uuid_user = ?", warehouseId, userId).First(&warehouse).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, custom_errors.ErrWareHouseNotFound
		}
		return nil, err
	}

	var zone domain.Zone
	if err := pr.db.GetDb().Where("id = ? AND ware_house_id = ?", in.ZoneId, warehouseId).First(&zone).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, custom_errors.ErrZoneNotFound
		}
		return nil, err
	}

	err := pr.db.GetDb().Create(&in)
	if err.Error != nil {
		return nil, err.Error
	}

	return in, nil
}

func (pr *ProductPostgresRepository) UpdateProductData(in *domain.Product, userId string, warehouseId int) error {
	warehouse := domain.WareHouse{}
	product := domain.Product{}

	if err := pr.db.GetDb().Where("id = ? AND uuid_user = ?", warehouseId, userId).First(&warehouse).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return custom_errors.ErrWareHouseNotFound
		}
		return err
	}

	var zone domain.Zone
	if err := pr.db.GetDb().Where("id = ? AND ware_house_id = ?", in.ZoneId, warehouseId).First(&zone).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return custom_errors.ErrZoneNotFound
		}
		return err
	}

	resultProduct := pr.db.GetDb().Model(&product).Where("uuid = ?", string(in.Uuid[:])).Select("title", "count", "description", "zone_id", "qr").Updates(in)
	if resultProduct.Error != nil {
		return resultProduct.Error
	}
	if resultProduct.RowsAffected == 0 {
		return custom_errors.ErrProductNotFound
	}

	return nil
}

func (pr *ProductPostgresRepository) DeleteProductData(in *domain.Product, userId string, warehouseId int) error {
	var warehouse domain.WareHouse
	if err := pr.db.GetDb().Where("id = ? AND uuid_user = ?", warehouseId, userId).First(&warehouse).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return custom_errors.ErrWareHouseNotFound
		}
		return err
	}

	var zone domain.Zone
	if err := pr.db.GetDb().Where("id = ? AND ware_house_id = ?", in.ZoneId, warehouseId).First(&zone).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return custom_errors.ErrZoneNotFound
		}
		return err
	}

	if err := pr.db.GetDb().Where("uuid = ?", in.Uuid).Delete(in); err != nil {
		return err.Error
	}

	return nil
}

func (pr *ProductPostgresRepository) FindAllProductFromZoneData(userId string, zoneId int) (*[]domain.Product, error) {
	var products []domain.Product

	result := pr.db.GetDb().Model(&domain.Product{}).
		Joins("JOIN zones ON products.zone_id = zones.id").
		Joins("JOIN ware_houses ON zones.ware_house_id = ware_houses.id").
		Where("ware_houses.uuid_user = ? AND zones.id = ?", userId, zoneId).
		Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}

func (pr *ProductPostgresRepository) FindAllProductFromWarehouseData(userId string, warehouseId int) (*[]domain.Product, error) {
	var products []domain.Product

	result := pr.db.GetDb().Model(&domain.Product{}).
		Joins("JOIN zones ON products.zone_id = zones.id").
		Joins("JOIN ware_houses ON zones.ware_house_id = ware_houses.id").
		Where("ware_houses.uuid_user = ? AND ware_houses.id = ?", userId, warehouseId).
		Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}

func (pr *ProductPostgresRepository) FindAllProductData(userId string) (*[]domain.Product, error) {
	var products []domain.Product

	if err := pr.db.GetDb().Model(&domain.Product{}).Joins("JOIN zones ON products.zone_id = zones.id").
		Joins("JOIN ware_houses ON zones.ware_house_id = ware_houses.id").
		Where("ware_houses.uuid_user = ?", userId).
		Find(&products); err != nil {
		return nil, err.Error
	}

	return &products, nil
}

func (pr *ProductPostgresRepository) FindProductData(userId string, productId string) (*domain.Product, error) {
	var product domain.Product

	if err := pr.db.GetDb().Model(&domain.Product{}).Joins("JOIN zones ON products.zone_id = zones.id").
		Joins("JOIN ware_houses ON zones.ware_house_id = ware_houses.id").
		Where("ware_houses.uuid_user = ? AND products.uuid = ?", userId, productId).
		Find(&product); err != nil {
		return nil, err.Error
	}

	return &product, nil
}
