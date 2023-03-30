package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	util "goutil/util"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Select the util that you want to use")
	fmt.Println("Enter the number of the util that you want to use")
	fmt.Println("1 - LowerCase method: convert a string in the lowercase")
	fmt.Println("0 - Exit: zero terminates the program - you can also terminate the program with 'exit' word")

	fmt.Print("Enter your option: ")
	option, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You have selected:", option)

	switch strings.Trim(option, "\r\n") {
	case "1":
		util.LowerCase()
	case "0":
		fmt.Println("Exit")
	case "exit":
		fmt.Println("Exit")
	default:
		fmt.Println("Option not available")
		fmt.Println("Exit")
	}

}
