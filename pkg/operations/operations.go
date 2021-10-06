package operations

import (
	"fmt"
	"strconv"
)

func Add(nums ...string) (int64, bool) {
    var sum int64
    sum = 0

    for _, num := range nums {
        num, err := strconv.ParseInt(num, 10, 0) // attempt to convert to int 

        if err != nil {
            fmt.Println("Couldn't parse an argument!")
            return sum, false // false for failure of function
        }
        
        sum += num
    }

    return sum, true // true for success of add function
}

func Subtract(nums ...string) (int64, bool) {
    num1, err := strconv.ParseInt(nums[0], 10, 0)

    if err != nil {
        fmt.Println("Couldn't parse an argument!")
        return 0, false // false for failure of function
    }

    num2, err := strconv.ParseInt(nums[1], 10, 0)

    if err != nil {
        fmt.Println("Couldn't parse an argument!")
        return 0, false // false for failure of function
    }

    return (num1 - num2), true

}

func Multiply(nums ...string) (int64, bool) {
    var prod int64
    prod = 1

    for _, num := range nums {
        num, err := strconv.ParseInt(num, 10, 0) // attempt to convert to int 

        if err != nil {
            fmt.Println("Couldn't parse an argument!")
            return prod, false // false for failure of function
        }
        
        prod *= num
    }

    return prod, true // true for success of multiply function
}

func Divide() {
	return
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
