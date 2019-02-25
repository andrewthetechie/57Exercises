package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func userInput(prompt string) (string, error) {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	if input == "" {
		return "", errors.New("user input cannot be blank")
	}

	return input, nil
}

func main() {
	principle, err := userInput("How much principal are you investing?")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pf, err := strconv.ParseFloat(principle, 32)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rate, err := userInput("What is the rate of interest?")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rf, err := strconv.ParseFloat(rate, 32)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	years, err := userInput("How many years are you going to invest?")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	yf, err := strconv.ParseFloat(years, 32)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	compounds, err := userInput("How many times a year will the interest be compounded?")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cf, err := strconv.ParseFloat(compounds, 32)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Simple Interest:")

	fmt.Printf("After %s years at $%s your investment %s will be worth $%.2f\n",
		years, rate, principle, pf+ (pf* rf / 100 * yf))
	fmt.Println("Compounded Interest:")
	fmt.Printf("After %s years at %s compounded %s times a year your investment $%s will be worth $%.2f\n",
		years, rate, compounds, principle, pf * math.Pow(1 + (rf / 100) / cf, yf * cf))
}