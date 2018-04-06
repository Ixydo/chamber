package main

import "github.com/ixydo/chamber/cmd"

var (
	// This is updated by linker flags during build
	Version = "dev"
)

func main() {
	cmd.Execute(Version)
}
