package infrastructure

import (
	"fmt"
	"github.com/Sieciechu/loans/loan/domain"
)

type InMemoryPersonRepository struct {
	personList map[string]domain.Person
}

// NewInMemoryPersonRepository - factory method
func NewInMemoryPersonRepository() InMemoryPersonRepository {
	return InMemoryPersonRepository{make(map[string]domain.Person)}
}

func (r *InMemoryPersonRepository) GetPersonByName(name string) (domain.Person, error) {
	var err error = nil

	person, ok := r.personList[name]
	if !ok {
		err = fmt.Errorf("There is no such person as '%s'", name)
	}
	return person, err
}
func (r *InMemoryPersonRepository) AddPerson(p domain.Person) error {
	_, ok := r.personList[p.GetName()]
	if ok {
		return fmt.Errorf("Cannot add person '%s', it already exists", p.GetName())
	}

	r.personList[p.GetName()] = p
	return nil
}
func (r *InMemoryPersonRepository) StorePerson(p domain.Person) error {
	r.personList[p.GetName()] = p
	return nil
}

func (r *InMemoryPersonRepository) Count() int {
	return len(r.personList)
}
