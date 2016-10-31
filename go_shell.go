package main
import ("fmt"
		"bufio"
		"os"
		"github.com/google/shlex")

func shell() {
	fmt.Printf("in shell func\n")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		token_text, _ := shlex.Split(text)
		fmt.Println(token_text)
	}
}

func executor() {
	//possibly use os.Open()
}

func main() {
	fmt.Printf("in main func\n")
	shell()	
}