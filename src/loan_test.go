package loans

import (
	"testing"

	"github.com/Sieciechu/isodate"
)

var date = isodate.NewDate

func TestWhenPersonLoansItemItHasItOnItsLoanList(t *testing.T) {
	// given
	john := newPerson(1, "John")

	// when
	john.loanTo(item{"dolars", 20.0}, person{name: "Adam"}, date("2020-01-16"))

	// then
	if loansCount := john.getLoans().Len(); 1 != loansCount {
		t.Errorf("Expected 1 loan, but got %d", loansCount)
		return
	}

	loan := john.getLoans().Front().Value.(loan)
	expectedItem := item{"dolars", 20.0}
	if expectedItem != loan.item {
		t.Errorf("Expected item %v, got %v", expectedItem, loan.item)
	}
	if john != loan.loaner {
		t.Errorf("Expected loaner %v, got %v", john, loan.loaner)
	}

	if expectedBorrower := (person{name: "Adam"}); expectedBorrower != loan.borrower {
		t.Errorf("Expected borrower %v, got %v", expectedBorrower, loan.borrower)
	}
}

func TestWhenPersonBorrowsItemItHasItOnItsLoanList(t *testing.T) {
	// given
	john := newPerson(1, "John")

	// when
	john.borrow(item{"dolars", 20.0}, person{name: "Adam"}, date("2020-01-16"))

	// then
	if loansCount := john.getLoans().Len(); 1 != loansCount {
		t.Errorf("Expected 1 loan, but got %d", loansCount)
		return
	}

	loan := john.getLoans().Front().Value.(loan)
	expectedItem := item{"dolars", 20.0}
	if expectedItem != loan.item {
		t.Errorf("Expected item %v, got %v", expectedItem, loan.item)
	}
	if john != loan.borrower {
		t.Errorf("Expected borrower %v, got %v", john, loan.borrower)
	}

	if expectedLoaner := (person{name: "Adam"}); expectedLoaner != loan.loaner {
		t.Errorf("Expected loaner %v, got %v", expectedLoaner, loan.loaner)
	}
}
