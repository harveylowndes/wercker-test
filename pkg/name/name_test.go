package name

import (
	"testing"
)

func TestBasic(t *testing.T) {
	test := &Test{}
	CreateName(test)
	t.Fail()
}
