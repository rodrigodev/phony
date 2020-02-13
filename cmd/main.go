package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/armon/go-radix"

	. "github.com/phony/internal/phone"
)

func main() {
	var phoneNumbersFile string

	if len(os.Args) > 1 {
		phoneNumbersFile = os.Args[1]
	} else {
		fmt.Println("no phone number file provided")
		os.Exit(1)
	}

	r := radix.New()

	// Reads all the file
	areaCodes, err := os.Open("./area_codes.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer areaCodes.Close()

	codeScanner := bufio.NewScanner(areaCodes)

	v := 0
	for codeScanner.Scan() {
		r.Insert(codeScanner.Text(), v)
		v += 1
	}

	potentialPhoneNumbers, err := os.Open(phoneNumbersFile)
	if err != nil {
		fmt.Print(err)
	}
	defer potentialPhoneNumbers.Close()

	numbersScanner := bufio.NewScanner(potentialPhoneNumbers)

	result := make(map[string]int)
	for numbersScanner.Scan() {
		n, valid := Sanitize(numbersScanner.Text())
		if !valid {
			continue
		}

		m, _, b := r.LongestPrefix(n)
		if b {
			result[m] = result[m] + 1
		}
	}

	for p, v := range result {
		if v > 0 {
			fmt.Printf("%s:%d\n", p, v)
		}
	}
}
