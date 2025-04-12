package shellspy

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func CommandFromString(input string) (*exec.Cmd, error) {
	inputs := strings.Split(input, " ")
	cmd := exec.Command(inputs[0], inputs[1:]...)
	return cmd, nil
	// how do I do error handling for exec.Command?
	// e.g. if the input is not a terminal command
	// https://pkg.go.dev/os/exec#Command
}

func ReadInputLoop(scanner bufio.Scanner) {
	for true {
		scanner.Scan()
		text := scanner.Text()
		input := text
		if input == "exit" {
			break
		}
		cmd, err := CommandFromString(input)
		if err != nil {
			log.Fatal(err)
		}
		// todo: sort out error handling
		var out strings.Builder
		cmd.Stdout = &out
		cmd.Run()
		fmt.Printf("%v", out.String())
	}
}
