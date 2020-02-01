package infrastructure

type serviceMap map[string](func() interface{})

func NewInMemoryContainer() InMemoryContainer {
	return InMemoryContainer{make(serviceMap)}
}

type InMemoryContainer struct {
	services serviceMap
}

func (c *InMemoryContainer) Register(key string, f func() interface{}) {
	c.services[key] = f
}

func (c *InMemoryContainer) Get(name string) interface{} {
	return c.services[name]()
}

func (c *InMemoryContainer) Has(name string) bool {
	_, ok := c.services[name]
	return ok
}
