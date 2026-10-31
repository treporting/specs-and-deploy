// Package demo breaks the build steps on purpose, for their demonstration (SR_51 to SR_58).
package demo

import "os"

// Remove drops its error: a finding of errcheck.
func Remove() { os.Remove("x") }

// Two answers 2.
func Two() int { return 2 }
