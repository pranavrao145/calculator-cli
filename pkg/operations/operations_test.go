package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
    // test 1: 2 args
    x, y := Add("3", "3")
    assert.Equal(t, x, float64(6), "3 + 3 = 6")
    assert.Equal(t, y, true, "Must return with success")

    // test 2: more args
    x, y = Add("3", "3", "3", "3")
    assert.Equal(t, x, float64(12), "3 + 3 + 3 + 3 = 12")
    assert.Equal(t, y, true, "Must return with success")

    // test 3: failure
    x, y = Add("3", "3", "s", "3")
    assert.Equal(t, y, false, "Should not return with success")
}

func TestSubtraction(t *testing.T) {
    // test 1: positive
    x, y := Subtract("6", "3")
    assert.Equal(t, x, float64(3), "6 - 3 = 3")
    assert.Equal(t, y, true, "Must return with success")

    // test 2: positive
    x, y = Subtract("-6", "3")
    assert.Equal(t, x, float64(-9), "-6 - 3 = -9")
    assert.Equal(t, y, true, "Must return with success")

    // test 3: failure
    x, y = Subtract("s", "3")
    assert.Equal(t, y, false, "Should not return with success")
}

func TestMultiplication(t *testing.T) {
    // test 1: 2 args
    x, y := Multiply("3", "3")
    assert.Equal(t, x, float64(9), "3 * 3 = 9")
    assert.Equal(t, y, true, "Must return with success")

    // test 2: more args
    x, y = Multiply("3", "3", "3", "3")
    assert.Equal(t, x, float64(81), "3 * 3 * 3 * 3 = 81")
    assert.Equal(t, y, true, "Must return with success")

    // test 3: failure
    x, y = Multiply("3", "3", "s", "3")
    assert.Equal(t, y, false, "Should not return with success")
}

func TestDivision(t *testing.T) {
    // test 1: integer 
    x, y := Divide("6", "3")
    assert.Equal(t, x, float64(2), "6 / 3 = 2")
    assert.Equal(t, y, true, "Must return with success")

    // test 2: float and int
    x, y = Divide("-6.3", "3")
    assert.Equal(t, x, float64(-2.1), "-6.3 / 3 = -2.1")
    assert.Equal(t, y, true, "Must return with success")

    // test 3: float and float
    x, y = Divide("-6.3", "3.33")
    assert.Equal(t, x, float64(-1.8918918918918919), "-6.3 / 3.33 = -1.8918918918919")
    assert.Equal(t, y, true, "Must return with success")

    // test 4: divide by zero
    x, y = Divide("3", "0")
    assert.Equal(t, y, false, "Must not return with success")

    // test 5: failure
    x, y = Divide("s", "3")
    assert.Equal(t, y, false, "Should not return with success")
}

func TestLog(t *testing.T) {
    // test 1: valid
    x, y := Log("10")
    assert.Equal(t, x, float64(1), "log(10) = 1")
    assert.Equal(t, y, true, "Must return with success")

    // test 2: invalid
    x, y = Subtract("s")
    assert.Equal(t, y, false, "Must not return with success")
}

func TestMod(t *testing.T) {
    // test 1: ints
    x, y := Mod("6", "4")
    assert.Equal(t, x, int64(2), "6 % 4 = 2")
    assert.Equal(t, y, true, "Must return with success")

    // test 1: floats (fail)
    x, y = Mod("6.3", "4")
    assert.Equal(t, y, false, "Must return with success")

    // test 3: failure
    x, y = Mod("s", "3")
    assert.Equal(t, y, false, "Should not return with success")
}
