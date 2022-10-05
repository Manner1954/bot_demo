package product

import (
	"errors"
)

type ServiceInterface interface {
	List() []Entity
	Get(idx int) (*Entity, error)
	New(title string) (string, error)
	Update(idx int, title string) (string, error)
	Remove(idx int) (string, error)
}

type ProductService struct{}

func NewService() *ProductService {
	return &ProductService{}
}

func (s *ProductService) List() []Entity {
	return allEntities
}

func (s *ProductService) Get(idx int) (*Entity, error) {
	if idx <= 0 {
		return nil, errors.New("Service.Get: idx must be more than 0")
	}

	if idx > (len(allEntities)) {
		return nil, errors.New("Service.Get: idx more than amount of entities")
	}

	idx--

	foundEntity := &allEntities[idx]
	return foundEntity, nil
}

func (s *ProductService) New(title string) (string, error) {
	if title == "" {
		return "", errors.New("Service.New: title is empty")
	}

	allEntities = append(allEntities, Entity{Title: title})

	return "Entity was created", nil
}

func (s *ProductService) Update(idx int, title string) (string, error) {
	if idx <= 0 {
		return "", errors.New("Service.Update: idx must be more than 0")
	}

	if idx > (len(allEntities)) {
		return "", errors.New("Service.Update: idx more than amount of entities")
	}

	if title == "" {
		return "", errors.New("Service.New: title is empty")
	}

	idx--

	allEntitiesBuf := allEntities[:idx]
	allEntitiesBuf = append(allEntitiesBuf, Entity{Title: title})
	allEntities = append(allEntitiesBuf, allEntities[idx+1:]...)

	return "Entity was update", nil
}

func (s *ProductService) Remove(idx int) (string, error) {
	if idx <= 0 {
		return "", errors.New("Service.Get: idx must be more than 0")
	}

	if idx > (len(allEntities)) {
		return "", errors.New("Service.Get: idx more than amount of entities")
	}

	idx--

	allEntities = append(allEntities[:idx], allEntities[idx+1:]...)
	return "Entity has been removed", nil
}
