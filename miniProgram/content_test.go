package miniProgram_test

import (
	"fmt"
	"testing"
)

var (
	content = miniapp.GetContent()
)

func TestCheckText(t *testing.T) {
	t.Parallel()
	text := ""
	err := content.CheckText(text)
	fmt.Print(err)
}

func TestCheckImage(t *testing.T) {
	t.Parallel()
	media := ""
	err := content.CheckImage(media)
	fmt.Print(err)
}
