package mapstore

import (
	"errors"
	"maya/assig3/domain"
)

type MapStore struct {
	store map[string]domain.Customer
}

func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}

func (ms *MapStore) Create(cust domain.Customer) error {
	if cust.ID == "" {
		return errors.New("error")
	}

	ms.store[cust.ID] = cust
	return nil
}

func (ms *MapStore) Update(ID string, c domain.Customer) error {
	if ID == "" {
		return errors.New("error")
	}
	ms.store[ID] = c
	return nil
}

func (ms *MapStore) GetByID(ID string) (domain.Customer, error) {
	if ID == "" {
		return domain.Customer{}, errors.New("error")
	}

	return ms.store[ID], nil
}

func (ms *MapStore) Delete(ID string) error {
	if ID == "" {
		return errors.New("error")
	}
	delete(ms.store, ID)
	return nil
}

func (ms *MapStore) GetAll() ([]domain.Customer, error) {
	var customers []domain.Customer
	for _, v := range ms.store {
		customers = append(customers, v)
	}

	return customers, nil
}
