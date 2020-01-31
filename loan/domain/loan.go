package domain

import (
	"time"
)

// Fun - returns 1
func Fun() int {
	return 1
}

type item struct {
	name     string
	quantity float64
}

type loan struct {
	loaner   Person
	borrower Person
	item     item
	date     time.Time
}
