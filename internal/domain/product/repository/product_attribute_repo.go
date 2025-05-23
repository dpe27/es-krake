package repository

import (
	"context"
	"pech/es-krake/internal/domain/product/entity"
	"pech/es-krake/internal/domain/shared/scope"
	"pech/es-krake/internal/domain/shared/transaction"
)

const OptionAttributeValueTableName = "option_attribute_values"

type OptionAttributeValueRepository interface {
	TakeByConditions(ctx context.Context, conditions map[string]interface{}, scopes ...scope.Base) (entity.OptionAttributeValue, error)

	FindByConditions(ctx context.Context, conditions map[string]interface{}, scopes ...scope.Base) ([]entity.OptionAttributeValue, error)

	CreateBatchWithTx(ctx context.Context, tx transaction.Base, attributeValues []map[string]interface{}, batchSize int) error

	Update(ctx context.Context, attributeValue entity.OptionAttributeValue, attributesToUpdate map[string]interface{}) (entity.OptionAttributeValue, error)

	DeleteByConditionsWithTx(ctx context.Context, tx transaction.Base, conditions map[string]interface{}, scopes ...scope.Base) error
}
