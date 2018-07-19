package name

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestBasic(t *testing.T) {
	test := &Test{}
	fmt.Printf("Wercker Report Directory: %s\n", os.Getenv("WERCKER_REPORT_DIR"))
	file, _ := os.OpenFile((os.Getenv("WERCKER_REPORT_DIR") + "/" + os.Getenv("WERCKER_STEP_ID") + "/" + "test.log"), os.O_WRONLY|os.O_CREATE, 0666)
	s := CreateName(test)
	io.Copy(file, strings.NewReader(s))
}
