package infrastructure

import (
	"github.com/Sieciechu/loans/loan/domain"
	"testing"
)

func TestGetPersonByNameReturnsPerson(t *testing.T) {
	// given
	// when
	// then
	var r domain.PersonRepository = &InMemoryPersonRepository{}
	r.GetPersonByName("kjkjk")
	t.Error("fail")
}

func TestGetPersonByNameReturnsErrorWhenThereIsNoSuchName(t *testing.T) {
	// given
	// when
	// then
	t.Fail()
}

func TestWhenAddPersonThenItIsInRepository(t *testing.T) {
	// given
	// when
	// then
	t.Fail()
}

func TestWhenAddPersonWhichIsInRepoThenItIsImpossibleToAddIt(t *testing.T) {
	// given
	// when
	// then
	t.Fail()
}

func TestWhenStoreThenPersonIsStored(t *testing.T) {
	// given
	// when
	// then
	t.Fail()
}

func TestWhenStoreDidNotSucceedThenReasonForFailureIsKnown(t *testing.T) {
	// given
	// when
	// then
	t.Fail()
}
