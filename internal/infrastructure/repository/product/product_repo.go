package repository

import (
	"context"
	"fmt"
	"pech/es-krake/internal/domain/product/entity"
	domainRepo "pech/es-krake/internal/domain/product/repository"
	domainScope "pech/es-krake/internal/domain/shared/scope"
	"pech/es-krake/internal/domain/shared/transaction"
	"pech/es-krake/internal/infrastructure/db"
	gormScope "pech/es-krake/internal/infrastructure/db/gorm/scope"
	"pech/es-krake/pkg/log"
	"pech/es-krake/pkg/utils"

	"gorm.io/gorm"
)

type productRepository struct {
	logger *log.Logger
	pg     *db.PostgreSQL
}

func NewProductRepository(pg *db.PostgreSQL) domainRepo.ProductRepository {
	return &productRepository{
		logger: log.With("repo", "product_repo"),
		pg:     pg,
	}
}

// Create implements repository.ProductRepository.
func (p *productRepository) Create(
	ctx context.Context,
	attributes map[string]interface{},
) (entity.Product, error) {
	product := entity.Product{}
	err := utils.MapToStruct(attributes, &product)
	if err != nil {
		p.logger.Error(ctx, utils.ErrorMapToStruct, "error", err.Error())
		return entity.Product{}, err
	}

	err = p.pg.DB.WithContext(ctx).Create(&product).Error
	return product, err
}

// CreateWithTx implements repository.ProductRepository.
func (p *productRepository) CreateWithTx(
	ctx context.Context,
	tx transaction.Base,
	attributes map[string]interface{},
) (entity.Product, error) {
	gormTx, ok := tx.GetTx().(*gorm.DB)
	if !ok {
		return entity.Product{}, fmt.Errorf(utils.ErrorGetTx)
	}

	product := entity.Product{}
	err := utils.MapToStruct(attributes, &product)
	if err != nil {
		p.logger.Error(ctx, utils.ErrorMapToStruct, "error", err.Error())
		return entity.Product{}, err
	}

	err = gormTx.Create(&product).Error
	return product, err
}

// FindByConditions implements repository.ProductRepository.
func (p *productRepository) FindByConditions(
	ctx context.Context,
	conditions map[string]interface{},
	scopes ...domainScope.Base,
) ([]entity.Product, error) {
	gormScopes, err := gormScope.ToGormScopes(scopes...)
	if err != nil {
		p.logger.Error(ctx, err.Error())
		return nil, err
	}
	products := []entity.Product{}
	err = p.pg.DB.
		WithContext(ctx).
		Scopes(gormScopes...).
		Where(conditions).
		Find(&products).Error
	return products, err
}

// TakeByConditions implements repository.ProductRepository.
func (p *productRepository) TakeByConditions(
	ctx context.Context,
	conditions map[string]interface{},
	scopes ...domainScope.Base,
) (entity.Product, error) {
	gormScopes, err := gormScope.ToGormScopes(scopes...)
	if err != nil {
		p.logger.Error(ctx, err.Error())
		return entity.Product{}, err
	}

	var product entity.Product
	err = p.pg.DB.
		WithContext(ctx).
		Scopes(gormScopes...).
		Where(conditions).
		Take(&product).Error
	return product, err
}

// Update implements repository.ProductRepository.
func (p *productRepository) Update(
	ctx context.Context,
	product entity.Product,
	attributesToUpdate map[string]interface{},
) (entity.Product, error) {
	err := utils.MapToStruct(attributesToUpdate, &product)
	if err != nil {
		p.logger.Error(ctx, utils.ErrorMapToStruct, "error", err.Error())
		return entity.Product{}, err
	}

	err = p.pg.DB.
		WithContext(ctx).
		Model(product).
		Updates(attributesToUpdate).Error
	return product, err
}

// UpdateWithTx implements repository.ProductRepository.
func (p *productRepository) UpdateWithTx(
	ctx context.Context,
	tx transaction.Base,
	product entity.Product,
	attributesToUpdate map[string]interface{},
) (entity.Product, error) {
	gormTx, ok := tx.GetTx().(*gorm.DB)
	if !ok {
		return entity.Product{}, fmt.Errorf(utils.ErrorGetTx)
	}

	err := utils.MapToStruct(attributesToUpdate, &product)
	if err != nil {
		p.logger.Error(ctx, utils.ErrorMapToStruct, "error", err.Error())
		return entity.Product{}, err
	}

	err = gormTx.Model(product).Updates(attributesToUpdate).Error
	return product, err
}
