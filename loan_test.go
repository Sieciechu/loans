package loans

import (
	"fmt"
	"testing"

	"github.com/Sieciechu/isodate"
)

var date = isodate.NewDate

func TestFun(t *testing.T) {
	if 1 != Fun() {
		t.Errorf("Fun not working")
	}
}

func TestFun2(t *testing.T) {
	inDate := "2020-01-02"
	aDate := date(inDate)
	outDate := fmt.Sprintf("%.4d-%.2d-%.2d", aDate.Year(), int(aDate.Month()), aDate.Day())

	if inDate != outDate {
		t.Errorf("Expected %s, but got %s", inDate, outDate)
	}
}

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
