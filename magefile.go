// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

const (
	appBin = "noublock.exe"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = Run

// Clean clean up after yourself
func Clean() {
	fmt.Println("Clean...")
	os.Remove(appBin)
}

// Build build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(Clean)
	fmt.Println("Build...")
	return sh.RunV("go", "build", "-o", appBin, "-v")
}

// Run execute app
func Run() error {
	mg.Deps(Build)
	fmt.Println("Run...")
	return sh.RunV("./" + appBin)
}
