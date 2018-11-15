package main

import (
	"fmt"
	"os"
	"dateconvert/eurodate"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Expecting date as argument")
		os.Exit(1)
	}

	cleanDate, err := eurodate.Format(os.Args[1])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cleanDate)
	}
}

