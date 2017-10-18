// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = Build

// A step that helps to get the needed + latest Nomad + Consul binaries
func Setup() error {
	checkAvailableNomadConsul()
	cmd := exec.Command("curl", "-L", "https://github.com/leowmjw")
	fmt.Println("ENV: ", cmd.Env)
	cmd.Dir = "/tmp"
	fmt.Println("My current working directory is ", cmd.Dir, " with PATH being ", cmd.Path)
	return cmd.Run()
}

func checkAvailableNomadConsul() error {
	fmt.Println("Checking the needed available binaries for Nomad + Consul")
	return nil
}

func findOutLatestNomadVersion() error {
	return nil
}

func findOutLatestConsulVersion() error {
	return nil
}

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(InstallDeps)
	fmt.Println("Building...")
	cmd := exec.Command("go", "build", "-o", "MyApp", ".")
	return cmd.Run()
}

// A custom install step if you need your bin someplace other than go/bin
func Install() error {
	mg.Deps(Build)
	fmt.Println("Installing...")
	return os.Rename("./MyApp", "/usr/bin/MyApp")
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	cmd := exec.Command("go", "get", "github.com/stretchr/piglatin")
	return cmd.Run()
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("MyApp")
}
