package core

// Secondary port - [ repository interface ]
type OrderRepository interface {
	Save(order Order) error
}
