package main

import (
	"fmt"
	"errors"
)

func main(){
	var myString string = "test"
	printIt(myString)

	numerator := 11
	denominator := 0
	result, remainder, err := intDivision(numerator, denominator)
	if err!=nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("The result is %v and the remainder %v",result, remainder)

	

}

func printIt(aString string) {
	fmt.Println(aString)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	result := numerator/denominator
	remainder := numerator%denominator
	return result, remainder, err
}