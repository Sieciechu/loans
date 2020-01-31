package infrastructure

import (
	"testing"
	"github.com/Sieciechu/loans/loan/domain"
)

func TestGet(t *testing.T) {
	var r domain.PersonRepository = &InMemoryPersonRepository{}
	r.GetPersonByName("kjkjk")
	t.Error("fail")
}