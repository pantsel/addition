package main

import (
	"fmt"
	"math/big"
	"os"
)

func getNumbers() []string {
	return os.Args[1:]
}

func isInputValid(numbers []string) bool {
	for _, v := range numbers {
		if i, ok := new(big.Int).SetString(v, 10); !ok || i.Cmp(big.NewInt(0)) < 0 {
			return false
		}
	}

	return true
}

func addNumbers(numbers []string) *big.Int {
	result := new(big.Int)

	for _, v := range numbers {
		n, _ := new(big.Int).SetString(v, 10)
		result.Add(result, n)
	}

	return result
}

func main() {

	numbers := getNumbers()

	if !isInputValid(numbers) {
		panic("Only unsigned integers can be used as arguments")
	}

	result := addNumbers(numbers)

	fmt.Println(result)

}
