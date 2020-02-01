package domain

import (
	"container/list"
	"time"
)

// Person - represents person.
// loans field contains both: loans and borrowings. In Polish we have one word for it (pożyczki), in English I couldn't find
type Person struct {
	id    int
	name  string
	loans *list.List
}

// GetId - returns person id
func (p *Person) GetId() int {
	return p.id
}

// GetName - returns person name
func (p *Person) GetName() string {
	return p.name
}

// NewPerson - factory function
func NewPerson(id int, name string) Person {
	return Person{id, name, list.New()}
}

func (p *Person) loanTo(i item, borrower Person, t time.Time) {
	p.loans.PushFront(loan{
		loaner:   *p,
		borrower: borrower,
		item:     i,
		date:     t,
	})
}

func (p *Person) borrow(i item, from Person, t time.Time) {
	p.loans.PushFront(loan{
		loaner:   from,
		borrower: *p,
		item:     i,
		date:     t,
	})
}

func (p *Person) getLoans() *list.List {
	return p.loans
}
