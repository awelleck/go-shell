package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
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
	//Execute command using exec
	cmdIn := exec.Command("bash", "-c", text)
	cmdOut, err := cmdIn.Output()

	if err != nil {
		return "", fmt.Errorf("%s was not a valid command.\n", text)
	}

	return string(cmdOut), nil
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
