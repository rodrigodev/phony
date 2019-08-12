package main

import (
	"bufio"
	"fmt"
	"github.com/armon/go-radix"
	. "github.com/phony/internal/phone"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var phoneNumbersFile string

	if len(os.Args) > 1 {
		phoneNumbersFile = os.Args[1]
	} else {
		fmt.Println("no phone number file provided")
		os.Exit(1)
	}

	file, err := os.Open("./area_codes.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	content, err := ioutil.ReadFile(phoneNumbersFile)
	if err != nil {
		fmt.Print(err)
	}
	potentialPhoneNumbers := strings.Split(string(content), "\n")

	scanner := bufio.NewScanner(file)

	r := radix.New()

	for scanner.Scan() {
		r.Insert(scanner.Text(), 0)
	}

	for _, number := range potentialPhoneNumbers {
		n, valid := Sanitize(number)
		if !valid {
			continue
		}

		// alternative method to find matches
		m, v, b := r.LongestPrefix(n)
		if b {
			r.Insert(m, v.(int)+1)
		}
	}

	printer := func(s string, v interface{}) bool {
		if v.(int) > 0 {
			fmt.Printf("%s:%d\n", s, v)
		}
		return false
	}
	r.Walk(printer)
}
