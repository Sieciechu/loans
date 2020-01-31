package infrastructure

import (
	"container/list"
	"github.com/Sieciechu/loans/loan/domain"
)


type InMemoryPersonRepository struct {
	personList *list.List
}

func (r *InMemoryPersonRepository) GetPersonByName(name string) (domain.Person, error) {
	return domain.Person{}, nil
}
func (r *InMemoryPersonRepository) AddPerson(p domain.Person) error {
	return nil
}
func (r *InMemoryPersonRepository) StorePerson(p domain.Person) error {
	return nil
}