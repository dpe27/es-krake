package repository

import (
	"context"
	"pech/es-krake/internal/domain/product/entity"
	domainRepo "pech/es-krake/internal/domain/product/repository"
	"pech/es-krake/internal/infrastructure/db"
	"pech/es-krake/pkg/log"
)

type attributeTypeRepository struct {
	logger *log.Logger
	pg     *db.PostgreSQL
}

func NewAttributeTypeRepository(pg *db.PostgreSQL) domainRepo.AttributeTypeRepository {
	return &attributeTypeRepository{
		logger: log.With("repo", "attribute_type_repo"),
		pg:     pg,
	}
}

// TakeByID implements repository.AttributeTypeRepository.
func (r *attributeTypeRepository) TakeByID(ctx context.Context, ID int) (entity.AttributeType, error) {
	attrType := entity.AttributeType{}
	db := r.pg.DB.WithContext(ctx)
	err := db.Take(&attrType, ID).Error
	return attrType, err
}

// GetAsDictionary implements repository.AttributeTypeRepository.
func (r *attributeTypeRepository) GetAsDictionary(ctx context.Context) (map[int]string, error) {
	var attrTypes []entity.AttributeType

	err := r.pg.DB.
		WithContext(ctx).
		Table(domainRepo.AttributeTypeTableName).
		Select("id", "name").
		Find(&attrTypes).Error
	if err != nil {
		return nil, err
	}

	dict := make(map[int]string)
	for _, attrType := range attrTypes {
		dict[attrType.ID] = attrType.Name
	}
	return dict, nil
}
