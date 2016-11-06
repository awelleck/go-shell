package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		text := scanner.Text()

		if text == "exit" {
			fmt.Println("Exiting gracefully...")
			os.Exit(0)
		}

		cmdIn := exec.Command("bash", "-c", text)
		cmdOut, err := cmdIn.Output()

		if err != nil {
			fmt.Printf("%s was not a valid command.", text)
		}

		fmt.Println(string(cmdOut))
		fmt.Print("> ")
	}
}
