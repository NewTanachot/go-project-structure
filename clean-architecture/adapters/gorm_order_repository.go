package adapters

import (
	"clean-architecture/entities"
	"clean-architecture/usecases"

	"gorm.io/gorm"
)

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) usecases.OrderRepository {
	return &GormOrderRepository{db: db}
}

func (r GormOrderRepository) Save(order entities.Order) error {
	return r.db.Create(&order).Error
}
