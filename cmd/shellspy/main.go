package main

import (
	"fmt"
	"shellspy"
)

var input string

func main() {
	fmt.Println("Recording session to 'shellspy.txt'")
	shellspy.StartMainLoop()
}
