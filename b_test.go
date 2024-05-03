package b_test

import (
	"testing"

	b "github.com/nobishino/gomodules-b"
)

func TestB(t *testing.T) {
	var _ b.B = struct{}{}
}
