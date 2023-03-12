package randgenerator_test

import (
	"testing"

	"github.com/p1ass/id/backend/pkg/randgenerator"
)

func TestMustGenerateToString_GeneratesDifferentValue(t *testing.T) {
	t.Parallel()

	got1 := randgenerator.MustGenerateToString(16)
	got2 := randgenerator.MustGenerateToString(16)

	if got1 == got2 {
		t.Errorf("two random string should be different, but identical: %s", got1)
	}
}
