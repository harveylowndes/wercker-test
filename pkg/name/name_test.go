package name

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestBasic(t *testing.T) {
	test := &Test{}
	file, _ := os.OpenFile((os.Getenv("WERCKER_REPORT_DIR") + "/" + "test.log"), os.O_WRONLY|os.O_CREATE, 0666)
	s := CreateName(test)
	io.Copy(file, strings.NewReader(s))
	t.Fail()
}
