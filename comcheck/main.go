package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	myArgs := arguments[1:]
	secArr := []string{"01", "galaxy", "galaxy 01"}
	isFound := false
	for _, myEl := range myArgs {
		if isFound {
			break
		}
		for _, mySec := range secArr {
			if myEl == mySec {
				fmt.Println("Alert!!!")
				isFound = true
				break
			}
		}

	}
}
