package storage

import (
	"fmt"
	"github.com/multiheaded/go-WebserviceTemplate/pkg/datamodel"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func NewGormRepository[T any](database *gorm.DB) datamodel.CRUDRepository[T] {
	return Repository[T]{
		DB: database,
	}
}

func (ctrl Repository[T]) Create(obj T) (T, error) {
	result := ctrl.DB.Create(&obj)
	if result.Error != nil {
		return obj, result.Error
	} else {
		return obj, nil
	}
}

func (ctrl Repository[T]) ReadAll(objs *[]T) error {
	result := ctrl.DB.Find(objs)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (ctrl Repository[T]) ReadOne(id uint64, obj *T) error {
	result := ctrl.DB.First(obj, id)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (ctrl Repository[T]) Update(id uint64, obj T) (T, error) {
	existing := new(T)

	err := ctrl.ReadOne(id, existing)

	if err != nil {
		return obj, err
	}

	result := ctrl.DB.Model(&existing).Updates(&obj)

	if result.Error != nil {
		return obj, result.Error
	}

	err = ctrl.ReadOne(id, existing)

	fmt.Println(err)

	if err != nil {
		return obj, err
	}

	return *existing, nil
}

func (ctrl Repository[T]) Delete(id uint64) error {
	var a T
	result := ctrl.DB.Delete(&a, id)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
