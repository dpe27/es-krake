package entity

import "pech/es-krake/internal/domain/shared/model"

type ProductOption struct {
	ID          int                     `gorm:"column:id;primaryKey;type:bigint;not null;autoIncrement" json:"id"`
	ProductID   int                     `gorm:"column:product_id;type:bigint;not null:"                 json:"product_id"`
	Name        *string                 `gorm:"column:name;type:varchar(50);not null"                   json:"name"`
	Description *string                 `gorm:"column:description;type:text"                            json:"description"`
	Attributes  []*OptionAttributeValue `gorm:"-:all"`
	model.BaseModel
}
