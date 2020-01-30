package main

import (
	"fmt"

	"github.com/Sieciechu/isodate"
	"github.com/Sieciechu/loans"
)

func main() {
	fmt.Println(isodate.NewDate("2018-02-12"))
	fmt.Println(loans.Fun())
}
