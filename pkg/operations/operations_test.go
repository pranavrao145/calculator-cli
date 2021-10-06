package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
    // test 1: 2 args
    x, y := Add("3", "3")
    assert.Equal(t, x, int64(6), "3 + 3 = 6")
    assert.Equal(t, y, true, "Must return with success")

    // test 2: more args
    x, y = Add("3", "3", "3", "3")
    assert.Equal(t, x, int64(12), "3 + 3 + 3 + 3 = 12")
    assert.Equal(t, y, true, "Must return with success")

    // test 3: failure
    x, y = Add("3", "3", "s", "3")
    assert.Equal(t, y, false, "Should not return with success")
}