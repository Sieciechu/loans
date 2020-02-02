package infrastructure

import (
	"testing"

	"github.com/Sieciechu/loans/loan/domain"
	"gotest.tools/assert"
)

func TestGetPersonByNameReturnsPerson(t *testing.T) {
	// given
	repo := NewInMemoryPersonRepository()
	repo.personList["Adam"] = domain.NewPerson(1, "Adam")
	repo.personList["Jack"] = domain.NewPerson(3, "Jack")

	originalPerson := domain.NewPerson(2, "John")
	repo.personList["John"] = originalPerson

	var personRepo domain.PersonRepository = &repo

	// when
	person, err := personRepo.GetPersonByName("John")

	// then
	assert.NilError(t, err)
	assert.Equal(t, originalPerson, person)
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
	assert.Equal(t, 2, personRepo.Count(), "Expected 2 people")
	if _, err := personRepo.GetPersonByName(newPersonName); nil != err {
		t.Errorf("Expected to have new person in repo, but got error '%s'", err)
	}
}

func TestWhenAddPersonWhichIsInRepoThenItIsImpossibleToAddIt(t *testing.T) {
	// given
	repo := NewInMemoryPersonRepository()
	repo.personList["Adam"] = domain.NewPerson(1, "Adam")
	var personRepo domain.PersonRepository = &repo

	// when
	err := personRepo.AddPerson(domain.NewPerson(2, "Adam"))

	// then
	if nil == err {
		t.Errorf("Expected error, got nothing")
	}
	assert.Equal(t, 1, personRepo.Count())

	person, _ := personRepo.GetPersonByName("Adam")
	assert.Equal(t, 1, person.GetId(), "Expected not to change data of person")
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
	assert.Equal(t, 1, personRepo.Count(), "Expected 1 person")

	person, _ := personRepo.GetPersonByName("Adam")
	assert.Equal(t, newId, person.GetId(), "Expected new id %d", newId)
}

func TestWhenStoreNewPersonThenPersonIsStored(t *testing.T) {
	// given
	repo := NewInMemoryPersonRepository()
	repo.personList["Adam"] = domain.NewPerson(1, "Adam")
	var personRepo domain.PersonRepository = &repo

	// when
	expectedPerson := domain.NewPerson(2, "John")
	personRepo.StorePerson(expectedPerson)

	// then
	assert.Equal(t, 2, personRepo.Count(), "Expected 2 people")
	person, _ := personRepo.GetPersonByName("John")
	assert.Equal(t, expectedPerson, person)
}
