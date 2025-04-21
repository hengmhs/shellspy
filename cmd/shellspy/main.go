package main

import (
	"fmt"
	"shellspy"
)

func main() {
	fmt.Println("Recording session to 'shellspy.txt'")
	shellspy.StartMainLoop()
}
