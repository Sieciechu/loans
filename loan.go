package loans

import (
	"container/list"
	"time"
)

// Fun - return 1
func Fun() int {
	return 1
}

type person struct {
	id    int
	name  string
	loans *list.List
}

func newPerson(id int, name string) person {
	return person{id, name, list.New()}
}

func (p *person) getName() string {
	return p.name
}

func (p *person) loanTo(i item, borrower person, t time.Time) {
	p.loans.PushFront(loan{
		loaner:   *p,
		borrower: borrower,
		item:     i,
		date:     t,
	})
}

func (p *person) getLoans() *list.List {
	return p.loans
}

type item struct {
	name     string
	quantity float64
}

type loan struct {
	loaner   person
	borrower person
	item     item
	date     time.Time
}
