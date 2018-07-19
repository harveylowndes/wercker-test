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
	fmt.Printf("Wercker Report Artifacts Directory: %s\n", os.Getenv("WERCKER_REPORT_ARTIFACTS_DIR"))
	fmt.Printf("Wercker Step ID: %s\n", os.Getenv("WERCKER_STEP_ID"))
	file, _ := os.OpenFile((os.Getenv("WERCKER_REPORT_ARTIFACTS_DIR") + "/" + "test.log"), os.O_WRONLY|os.O_CREATE, 0666)
	s := CreateName(test)
	io.Copy(file, strings.NewReader(s))
}
