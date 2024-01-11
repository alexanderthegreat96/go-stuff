package main

import (
	"fmt"
)

// to check the go env
// run in terminal
// go env GOARCH GOOS

// build commands

// go build . //builds for the current architecture
// bash commands
// GOOS=darwin go build // mac
// GOOS=amd64 go build // windows
// GOOS=linux GOARCH=amd64 go build

// windows terminal / shell
// go build -o DIR_PATH_HERE/filename.exe program
// in the powershell / cmd promp we use $env: to declare GO ENV variables
// $env:GOOS="linux"; $env:GOARCH="amd64"; go build -o binaries/filename_amd64_linux

func main() {
	fmt.Println("Hello world. This is the first test-module")

	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
