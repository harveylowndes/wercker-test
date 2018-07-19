package integration

import (
	"testing"

	"github.com/harveylowndes/wercker-test/pkg/name"
)

func TestBasic(t *testing.T) {
	test := &name.Test{}
	name.CreateName(test)
	t.Fail()
}
