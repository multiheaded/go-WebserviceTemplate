package datamodel

type CRUDRepository[T any] interface {
	Create(obj T) (T, error)
	ReadAll(objs *[]T) error
	ReadOne(id uint64, obj *T) error
	Update(id uint64, obj T) (T, error)
	Delete(id uint64) error
}
