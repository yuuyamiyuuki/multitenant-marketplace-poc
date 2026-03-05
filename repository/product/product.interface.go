//Creating interface for Product

import "../entities"

type ProductRepositoryInterface interface{
	create(
		product *entities.Product
	) error
	
	update(
		product *entities.Product
	) error

	delete(
		id datatypes.UUID
	) error

	find_by_id(
		id datatyupes
	) (error,*entities.Product) 
}