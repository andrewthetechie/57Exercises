package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func userInput(prompt string, stripNewlines bool) (string, error) {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if stripNewlines {
		input = strings.TrimSuffix(input, "\n")
		input = strings.TrimSuffix(input, "\r")
	}

	if input == "" {
		return "", errors.New("User input cannot be blank")
	}


	return input, nil
}

func main() {
	name, err := userInput("What is your name?", true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Hello %s, it is nice to meet you\n", name)
	fmt.Printf("%s your name has %d letters\n", name, len(name))

	d, err := userInput("What is your birthday? Enter it like mm/dd/yyyy", true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dob, err := time.Parse("01/02/2006", d)
	if err != nil {
		fmt.Println("Error parsing your date of birth. Be sure to enter it like mm/dd/yyyy")
		os.Exit(1)
	}
	yearsOld := int(time.Now().Sub(dob).Hours() / 8760)
	fmt.Printf("You are %d years old today.\n", yearsOld)
	r, err := userInput("At what age would you like to retire?", true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	retire, err := strconv.Atoi(r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if retire < yearsOld {
		fmt.Println("You can retire right now!")
		os.Exit(0)
	}

	yearsUntilRetirement := retire - yearsOld
	yearOfRetirement := time.Now().AddDate(yearsUntilRetirement, 0, 0).Year()
	fmt.Printf("You want to retire in %d years. That'll be the year %d\n",
		yearsUntilRetirement,
		yearOfRetirement)
}
