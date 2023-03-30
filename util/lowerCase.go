package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func LowerCase() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the string that you want to lower case: ")
		getString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		lowerCaseString := strings.ToLower(getString)
		fmt.Println(lowerCaseString)

		if strings.Trim(getString, "\r\n") == "0" || strings.Trim(getString, "\r\n") == "exit" {
			fmt.Println("Exit")
			break
		}
	}
}
