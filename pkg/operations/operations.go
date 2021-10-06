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

func Subtract() {
	return
}

func Multiply() {
	return
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
