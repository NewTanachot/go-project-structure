package usecases

import (
	"clean-architecture/entities"
)

type OrderRepository interface {
	Save(order entities.Order) error
}
