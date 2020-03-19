package infrastructure

import (
	"testing"

	"github.com/Sieciechu/loans/loan/application"
	"github.com/Sieciechu/loans/loan/domain"
	"github.com/Sieciechu/gotest.tools/assert"
)

func TestGet(t *testing.T) {
	// given
	c := NewInMemoryContainer()
	c.services["returnAlwaysSameNameService"] = func() interface{} {
		return "John"
	}

	var container application.Container = &c

	// when
	service := container.Get("returnAlwaysSameNameService")

	// then
	assert.Equal(t, "John", service)
}

func TestRegister(t *testing.T) {
	// given
	c := NewInMemoryContainer()
	var container application.Container = &c

	assert.Assert(t, !container.Has("returnAlwaysSameNameService"))

	// when
	container.Register("returnAlwaysSameNameService", func() interface{} {
		return "John"
	})

	// then
	assert.Assert(t, container.Has("returnAlwaysSameNameService"))
	assert.Equal(t, "John", container.Get("returnAlwaysSameNameService"))
	assert.Equal(t, "John", container.Get("returnAlwaysSameNameService"))
}

func TestGetSingletonService(t *testing.T) {
	// given
	c := NewInMemoryContainer()

	// function which returns func() interface{}
	// it's a closure, so var repo will remain and will be singleton
	singletonPersonFactoryFunction := func() func() interface{} {
		var repo domain.PersonRepository = nil

		factoryFunction := func() interface{} {
			if nil == repo {
				x := NewInMemoryPersonRepository()
				x.AddPerson(domain.NewPerson(1, "John"))
				repo = &x
			}
			return repo
		}
		return factoryFunction
	}
	c.services["personrepository"] = singletonPersonFactoryFunction()

	var container application.Container = &c

	// when
	service, ok := container.Get("personrepository").(domain.PersonRepository)

	// then
	if !ok {
		t.Error("Expected domain.PersonRepository. Could not cast to it")
		t.Errorf("service is %#v", service)
	}

	service.AddPerson(domain.NewPerson(2, "Adam"))

	service2 := container.Get("personrepository").(domain.PersonRepository)
	assert.Equal(t, service, service2)
	assert.Equal(t, service.Count(), service2.Count())
}
