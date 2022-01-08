package miniProgram_test

import (
	"fmt"
	"testing"
)

func TestGetWeRunData(t *testing.T) {
	t.Parallel()
	sessionKey := ""
	encryptedData := ""
	iv := ""
	wr := miniapp.GetWerun()
	data, err := wr.GetWeRunData(sessionKey, encryptedData, iv)
	fmt.Print(data, err)
}
