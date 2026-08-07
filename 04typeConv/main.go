package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main(){

	fmt.Println("Please enter the room rating form 1-5: ")

	reader := bufio.NewReader(os.Stdin)

	input,_ := reader.ReadString('\n')

	fmt.Printf("Thanks for rating %s, ",input )

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Rating increased by +1: ",numRating+1)
	}
}