package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

var (
	history    []string
	historyPos int
)

func helper() {
	//Use a channel to handle SIGINT, prompt user
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT)
	for {
		select {
		case <-sigc:
			fmt.Println("SIGINT attempt: if you want to quit use the 'exit' command.")
			fmt.Print("> ")
		}
	}
}

func executor(text string) (string, error) {
	//add to history, so we can go back to it.
	history = append(history, text)

	// seperate strings by spaces
	tokenCommand := strings.Split(text, " ")

	var output string
	switch tokenCommand[0] {
	case "cd":
		err := os.Chdir(tokenCommand[1])
		if err != nil {
			log.Printf("Error changing to directory %s: %v", tokenCommand[1], err)
		}
	case "history":
		fmt.Println(history)
	default:
		// Execute command using exec
		cmdIn := exec.Command("bash", "-c", text)
		cmdOut, err := cmdIn.Output()
		if err != nil {
			return "", fmt.Errorf("%s was not a valid command.\n", text)
		}
		output = string(cmdOut)
	}
	return output, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	//Go routine
	go helper()

	fmt.Print("> ")
	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			fmt.Println("Exiting gracefully...")
			os.Exit(0)
		}

		output, err := executor(text)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf(output)
		}

		fmt.Print("> ")
	}
}
