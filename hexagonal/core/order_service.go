package core

import "errors"

// Primary port - [ Interface ]
type OrderService interface {
	CreateOrder(order Order) error
}

// struc dependency - [ class dependency ]
type orderServiceImpl struct {
	repo OrderRepository
}

// struct init - [ class constructor for set default dependency value in class ]
func NewOrderService(repo OrderRepository) OrderService {
	return &orderServiceImpl{repo: repo}
}

// function implementation (automatic become member of interface)
func (s *orderServiceImpl) CreateOrder(order Order) error {
	if order.Total <= 0 {
		return errors.New("total must be positive")
	}

	if err := s.repo.Save(order); err != nil {
		return err
	}
	return nil
}
