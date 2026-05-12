package main

import (
	"bufio"
	"fmt"
	"os"
)

// bufio: buffered i/o, instead of reading/writing one character at a time it reads/writes in chunks that improves performance
// os: gives access to  operatinng system features like: files,environmental variable,cla,and standard  i/o 
// input,_ is comma ok || comma err syntax
// if we don't want to deal with error we just put(_) and vise/versa
func main(){


	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your rating:")

	input,_ := reader.ReadString('\n')
	
	fmt.Printf("Thanks  for your rating %s",input)
}