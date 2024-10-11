package hello_test

import (
	"testing"

	hello "github.com/syntaxofffedora/myhello"
)

func TestPrintPrintsHelloMessageToTerminal(t *testing.T) {
	t.Parallel()
	hello.Print()
}
