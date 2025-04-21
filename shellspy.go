package shellspy

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func CommandFromString(input string) *exec.Cmd {
	inputs := strings.Split(input, " ")
	cmd := exec.Command(inputs[0], inputs[1:]...)
	// exec.Command does not return an error immediately
	// it only constructs a Cmd object
	// error only when .Run(), .Output() or .CombinedOutput() is called
	// where is this in the docs?? https://pkg.go.dev/os/exec#Command
	return cmd
}

func ReadInputLoop(scanner *bufio.Scanner, transcript string) {
	for {
		scanner.Scan()
		text := scanner.Text()
		input := text
		if input == "exit" {
			WriteToTextFile(transcript, "exit", "")
			break
		}
		cmd := CommandFromString(input)
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

func StartMainLoop() {
	const transcript = "shellspy.txt"
	CreateTextFile(transcript)
	scanner := bufio.NewScanner(os.Stdin)
	ReadInputLoop(scanner, transcript)
}
