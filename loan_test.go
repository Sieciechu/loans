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
