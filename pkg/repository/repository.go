package repository

type CrudRepository[T any] interface {
	Create(entity *T) error
	FindByID(id string) (*T, error)
	FindAll() ([]T, error)
	Update(entity *T) error
	Delete(id string) error
}
