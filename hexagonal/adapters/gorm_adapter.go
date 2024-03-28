package adapters

import (
	"hexagonal/core"

	"gorm.io/gorm"
)

// Secondary adapter - [ repository dependency ]
type GormOrderRepository struct {
	db *gorm.DB
}

// struct init - [ repository constructor for set default dependency value ]
func NewGormOrderRepository(db *gorm.DB) core.OrderRepository {
	return &GormOrderRepository{db: db}
}

func (r *GormOrderRepository) Save(order core.Order) error {
	if result := r.db.Create(&order); result.Error != nil {
		// Handle database errors
		return result.Error
	}
	return nil
}
