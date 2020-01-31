package infrastructure

import (
	"github.com/Sieciechu/loans/loan/domain"
	"testing"
)

func TestGetPersonByNameReturnsPerson(t *testing.T) {
	// given
	repo := NewInMemoryPersonRepository()
	repo.personList["Adam"] = domain.NewPerson(1, "Adam")
	repo.personList["John"] = domain.NewPerson(2, "John")
	repo.personList["Jack"] = domain.NewPerson(3, "Jack")
	var personRepo domain.PersonRepository = &repo

	// when
	person, err := personRepo.GetPersonByName("John")

	// then
	expectedId := 2
	expectedName := "John"
	if person.GetId() != expectedId && person.GetName() != expectedName {
		t.Errorf("Expected person {id: %d, name: %s}, got {id: %d, name: %s}",
			expectedId, expectedName, person.GetId(), person.GetName())
	}
	if nil != err {
		t.Errorf("Expected no errors, got %#v", err)
	}
}

func TestGetPersonByNameReturnsErrorWhenThereIsNoSuchName(t *testing.T) {
	repo := NewInMemoryPersonRepository()
	repo.personList["Adam"] = domain.NewPerson(1, "Adam")
	repo.personList["John"] = domain.NewPerson(2, "John")
	repo.personList["Jack"] = domain.NewPerson(3, "Jack")
	var personRepo domain.PersonRepository = &repo

	// when
	_, err := personRepo.GetPersonByName("Some Not Existing Name")

	// then
	if nil == err {
		t.Errorf("Expected error, got nothing")
	}

}

func TestWhenAddPersonThenItIsInRepository(t *testing.T) {
	// given
	repo := NewInMemoryPersonRepository()
	repo.personList["Adam"] = domain.NewPerson(1, "Adam")
	var personRepo domain.PersonRepository = &repo

	// when
	newPersonName := "John"
	personRepo.AddPerson(domain.NewPerson(2, newPersonName))

	// then
	if 2 != personRepo.Count() {
		t.Errorf("Expected 2 people, got %d", personRepo.Count())
	}
	if _, err := personRepo.GetPersonByName(newPersonName); nil != err {
		t.Errorf("Expected to have %s in repo, but got error '%s'", newPersonName, err)
	}
}

func TestWhenAddPersonWhichIsInRepoThenItIsImpossibleToAddIt(t *testing.T) {
	// given
	repo := NewInMemoryPersonRepository()
	repo.personList["Adam"] = domain.NewPerson(1, "Adam")
	var personRepo domain.PersonRepository = &repo

	// when
	personRepo.AddPerson(domain.NewPerson(2, "Adam"))

	// then
	if 1 != personRepo.Count() {
		t.Errorf("Expected 1 people, got %d", personRepo.Count())
	}

	person, _ := personRepo.GetPersonByName("Adam")
	if person.GetId() != 1 {
		t.Errorf("Expected not to change data of person")
	}
}

func TestWhenStoreAndPersonExistsThenPersonIsUpdated(t *testing.T) {
	// given
	repo := NewInMemoryPersonRepository()
	repo.personList["Adam"] = domain.NewPerson(1, "Adam")
	var personRepo domain.PersonRepository = &repo

	// when
	newId := 2
	personRepo.StorePerson(domain.NewPerson(newId, "Adam"))

	// then
	if 1 != personRepo.Count() {
		t.Errorf("Expected 1 people, got %d", personRepo.Count())
	}

	person, _ := personRepo.GetPersonByName("Adam")
	if person.GetId() != newId {
		t.Errorf("Expected new id: %d, got %d", newId, person.GetId())
	}
}

func TestWhenStoreNewPersonThenPersonIsStored(t *testing.T) {
	// given
	repo := NewInMemoryPersonRepository()
	repo.personList["Adam"] = domain.NewPerson(1, "Adam")
	var personRepo domain.PersonRepository = &repo

	// when
	john := domain.NewPerson(2, "John")
	personRepo.StorePerson(john)

	// then
	if 2 != personRepo.Count() {
		t.Errorf("Expected 2 people, got %d", personRepo.Count())
	}
	person, _ := personRepo.GetPersonByName("John")
	if person != john {
		t.Errorf("Expected person %v, got %v", john, person)
	}
}
