package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		//fmt.Println(text)

		if text == "exit" {
			fmt.Println("Exiting gracefully...")
			exit_shell()
		}

		lsCmd := exec.Command("bash", "-c", text)
		lsOut, err := lsCmd.Output()
		if err != nil {
			panic(err)
		}

		fmt.Println("Command executing:", text)
		fmt.Println(string(lsOut))
	}
}
