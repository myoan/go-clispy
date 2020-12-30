package main

import (
	"fmt"
	"os"
)

func useFileRead(filename string) string {
	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	buf := make([]byte, 64)
	var result string
	for {
		n, err := fp.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		result += string(buf[0:n])
	}
	return result
}

/*
func useStdinRead() {
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Printf("> ")
	for stdin.Scan() {
		if err := stdin.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		text := stdin.Text()
		if text == "quit" {
			fmt.Println("bye.")
			break
		}
		fmt.Println(text)
		fmt.Printf("> ")
	}
}
*/

func main() {
	if len(os.Args) != 2 {
		return
	}
	program := useFileRead(os.Args[1])
	// fmt.Println(program)
	// fmt.Println("")
	tokens, err := Tokenize(program)
	if err != nil {
		panic("SyntaxError")
	}
	ast, _ := Parse(tokens)
	// ast.Show()
	instSet := Compile(ast)
	// instSet.Show()
	vm := NewVirtualMachine()
	vm.Exec(instSet)
	// vm.Show()
	fmt.Printf("result: %d\n", vm.Result())
}
