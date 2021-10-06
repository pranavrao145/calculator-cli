package operations

import (
	"fmt"
	"strconv"
)

func Add(nums ...string) (float64, bool) {
	var sum float64
	sum = 0

	for _, num := range nums {
		num, err := strconv.ParseFloat(num, 10) // attempt to convert to int

		if err != nil {
			fmt.Println("Couldn't parse an argument!")
			return sum, false // false for failure of function
		}

		sum += num
	}

	return sum, true // true for success of add function
}

func Subtract(nums ...string) (float64, bool) {
	num1, err := strconv.ParseFloat(nums[0], 10)

	if err != nil {
		fmt.Println("Couldn't parse an argument!")
		return 0, false // false for failure of function
	}

	num2, err := strconv.ParseFloat(nums[1], 10)

	if err != nil {
		fmt.Println("Couldn't parse an argument!")
		return 0, false // false for failure of function
	}

	return (num1 - num2), true

}

func Multiply(nums ...string) (float64, bool) {
	var prod float64
	prod = 1

	for _, num := range nums {
		num, err := strconv.ParseFloat(num, 10) // attempt to convert to int

		if err != nil {
			fmt.Println("Couldn't parse an argument!")
			return prod, false // false for failure of function
		}

		prod *= num
	}

	return prod, true // true for success of multiply function
}

func Divide(nums ...string) (float64, bool) {
	num1, err := strconv.ParseFloat(nums[0], 0)

	if err != nil {
		fmt.Println("Couldn't parse an argument!")
		return 0, false // false for failure of function
	}

	num2, err := strconv.ParseFloat(nums[1], 0)

	if err != nil {
		fmt.Println("Couldn't parse an argument!")
		return 0, false // false for failure of function
    }

    // check for division by 0
    if num2 == float64(0) {
		fmt.Println("Can't divide by zero!")
		return 0, false // false for failure of function
    }

	return (num1 / num2), true
}

func Exp() {
	return
}

func Log() {
	return
}

func Mod() {
	return
}
