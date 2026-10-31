// Package demo breaks the build steps on purpose, for their demonstration (SR_51 to SR_58).
package demo

// Two answers a string where it promises an int: it does not build.
func Two() int { return "2" }
