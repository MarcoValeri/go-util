package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func UrlGenerator() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter '0' or 'exit' to end the program")
		fmt.Print("Enter the string that you want to transform to url: ")
		getString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if strings.Trim(getString, "\r\n") == "0" || strings.Trim(getString, "\r\n") == "exit" {
			fmt.Println("Exit")
			break
		}

		stringToUrl := strings.ReplaceAll(getString, " ", "-")
		stringToUrl = strings.ToLower(stringToUrl)
		fmt.Println(stringToUrl)
	}
}
