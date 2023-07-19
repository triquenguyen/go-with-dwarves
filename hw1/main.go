package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args[1:]

	if len(input) < 3 {
		fmt.Println("Please enter your first name, last name and country code")
		return
	}

	middleName := ""
	firstName := input[0]
	lastName := input[1]
	countryCode := input[len(input)-1]

	if len(input) > 4 {
		middleName = strings.Join(input[2:len(input)-1], " ")
	} else if len(input) == 4 {
		middleName = input[2]
	}

	switch countryCode {
	case "VN", "CN", "JP", "KR", "TH", "KH":
		if middleName == "" {
			fmt.Println(lastName + " " + firstName)
		} else {
			fmt.Println(lastName + " " + middleName + " " + firstName)
		}

	case "US", "UK", "AU", "CA":
		if middleName == "" {
			fmt.Println(firstName + " " + lastName)
		} else {
			fmt.Println(firstName + " " + middleName + " " + lastName)
		}
	}
}
