package check

import "testing"

// Fails on purpose: the demonstration of SR-10. Never merged.
func TestTheDemonstrationFails(t *testing.T) {
	t.Fatal("fails on purpose, to show that a failed test blocks the merge")
}
