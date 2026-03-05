//Repository functions to comunicate with interface "extends" (product *entitites.Product)

type user_repository struct{
	db *gorm.DB

}

func ProductRepository(db *gorm.DB) ProductRepositoryInterface{
	return &user_repository {db:db}
}

func (repository *user_repository) CreateProduct(product *entitites.Product)error{
	repository.db.Create(product)
}

func (repository *user_repository) UpdateProduct(product *entitites.Product)error{
	repository.db.Update(product)
}

func (repository *user_repository) DeleteProduct(product *entitites.Product)error{
	repository.db.Delete(product)
}