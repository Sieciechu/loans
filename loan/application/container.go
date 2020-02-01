package application

type Container interface {
	Get(name string) interface{}
	Register(name string, f func() interface{})
	Has(name string) bool
}
