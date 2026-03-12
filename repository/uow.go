package repository

import "gorm.io/gorm"

type UnitOfWork interface {
	Do(fn func(uow UnitOfWork) error) error
	CartRepo() CartRepository
	ProductRepo() ProductRepository
	ClientRepo() ClientRepository
	OrderRepo() OrderRepository
	OrderItemRepo() OrderItemRepository
}

type unitOfWork struct {
	db *gorm.DB
}

func NewUnitOfWork(db *gorm.DB) UnitOfWork {
	return &unitOfWork{db: db}
}

func (u *unitOfWork) Do(fn func(uow UnitOfWork) error) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		txUow := &unitOfWork{db: tx}
		return fn(txUow)
	})
}

func (u *unitOfWork) CartRepo() CartRepository {
	return NewCartRepository(u.db)
}

func (u *unitOfWork) ProductRepo() ProductRepository {
	return NewProductRepository(u.db)
}

func (u *unitOfWork) ClientRepo() ClientRepository {
	return NewClientRepository(u.db)
}

func (u *unitOfWork) OrderRepo() OrderRepository {
	return NewOrderRepository(u.db)
}

func (u *unitOfWork) OrderItemRepo() OrderItemRepository {
	return NewOrderItemRepository(u.db)
}
