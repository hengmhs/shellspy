package shellspy

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	transcript := "shellspy.txt"
	CreateTextFile(transcript)
	for true {
		scanner.Scan()
		text := scanner.Text()
		input := text
		if input == "exit" {
			WriteToTextFile(transcript, "exit", "")
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
		WriteToTextFile(transcript, input, out.String())
		fmt.Printf("%v", out.String())
	}
}

func CreateTextFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Failed to create file: %v", err)
	}
	defer file.Close()
}

func WriteToTextFile(fileName, command, output string) {
	// Open the file and append only
	// 0644 is file permissions
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	// Closes the file when the function exits
	defer file.Close()
	// Write the command with >
	_, err = fmt.Fprintf(file, "> %v \n", command)
	// Write the output
	_, err = fmt.Fprintf(file, "%v \n", output)
}
