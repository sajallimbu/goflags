package main

import (
	"fmt"
	"runtime"
)

var version string

func main() {
	fmt.Printf("OS: %s\nArchitechture: %s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("\nVersion: %s", version)
}
