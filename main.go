package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func execute_unixlike() {
	out, err := exec.Command("ifconfig", "-a").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	fmt.Println(output)
}

func execute_windows() {
	out, err := exec.Command("ipconfig", "/all").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	fmt.Println(output)
}

func main() {
	if runtime.GOOS == "windows" {
		execute_windows()
	} else {
		execute_unixlike()
	}
}
