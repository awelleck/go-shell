package main

import (
	"bufio"
	"fmt"
	"github.com/google/shlex"
	"os"
	"os/exec"
)

func shell() {
	fmt.Printf("in shell func\n")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		token_text, _ := shlex.Split(text)
		fmt.Println(token_text)
		return token_text
	}
}

func executor() {
	//using os/exec to execute commands
	cmd := exec.Command()
	err := cmd.Run()
    }
}

func main() {
	fmt.Printf("in main func\n")
	shell()
}
