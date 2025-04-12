package main

import (
	"bufio"
	"fmt"
	"os"
	"shellspy"
)

var input string

func main() {
	fmt.Println("Recording session to 'shellspy.txt'")
	scanner := bufio.NewScanner(os.Stdin)
	shellspy.ReadInputLoop(*scanner)
}
