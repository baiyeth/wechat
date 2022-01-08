package miniProgram_test

import (
	"fmt"
	"testing"
)

func TestCode2Session(t *testing.T) {
	t.Parallel()
	au := miniapp.GetAuth()
	res, err := au.Code2Session("")
	fmt.Print(res, err)
}
