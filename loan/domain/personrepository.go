package domain

// PersonRepository - interface for getting and storing Person
type PersonRepository interface {
	// GetPersonByName - returns person or error if not found
	GetPersonByName(name string) (Person, error)

	// AddPerson - adds person. If it exists, error is returned
	AddPerson(p Person) error

	// StorePerson - stores new person. If person exists it is updated
	StorePerson(p Person) error

	// Count - returns count
	Count() int
}
