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

func ReadInputLoop(scanner *bufio.Scanner, transcript string, transcriptFile *os.File) {
	for {
		scanner.Scan()
		text := scanner.Text()
		input := text
		if input == "exit" {
			WriteToTextFile(transcript, "exit", "")
			break
			// TODO: maybe return something here to check that the loop has broken?
		}
		fmt.Printf("file: %v\n", transcriptFile)
		cmd := CommandFromString(input)
		fmt.Printf("Running command: %v\n", cmd.Args)

		// var out strings.Builder
		// cmd.Stdout = &out
		cmd.Stdout = transcriptFile
		err := cmd.Run()

		if err != nil {
			fmt.Printf("Error running command: %v", err)
			// TODO: write a test that shows this error handling is working
		}

		// TODO: add formatting for transcript

		// Tried to use cmd.Stdout = transcriptFile but that didn't do anything..
		// How do I get the text formatting if I pipe it out directly?
		// Keep getting exit status 1 or 2 errors

		// WriteToTextFile(transcript, input, out.String())
		// fmt.Printf("> %v\n", input)
		// fmt.Printf("%v", out.String())
	}
}

func CreateTextFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Failed to create file: %v", err)
	}
	return file
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
	const transcriptFileName = "shellspy.txt"
	transcriptFile := CreateTextFile(transcriptFileName)
	defer transcriptFile.Close() // Close the file only after ReadInputLoop has written to it
	scanner := bufio.NewScanner(os.Stdin)
	ReadInputLoop(scanner, transcriptFileName, transcriptFile)
}
