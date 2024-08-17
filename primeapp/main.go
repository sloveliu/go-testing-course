package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// print a welcome
	intro()
	// create a channel
	doneChan := make(chan bool)
	// start goroutine
	go readUserInput(os.Stdin, doneChan)
	// block until get a value
	<-doneChan
	// close the channel
	close(doneChan)
	// say goodbye
	fmt.Println("Bye!")
}

// 這樣寫 os.Stdin 是 hardcode 無法測試
// func readUserInput(doneChan chan bool) {
// 	scanner := bufio.NewScanner(os.Stdin)

// 	for {
// 		res, done := checkNumbers(scanner)
// 		if done {
// 			doneChan <- true
// 			return
// 		}

// 		fmt.Println(res)
// 		prompt()
// 	}
// }

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkNumbers(scanner)
		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read user input
	scanner.Scan()
	// check to see if the user wants to quit
	// EqualFold 忽略大小寫，比較字串是否相等，回傳 bool
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// try to convert what the user typed into an int
	numberToCheck, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number!", false
	}

	// check to see if the number is prime
	_, msg := isPrime(numberToCheck)
	return msg, false
}

func intro() {
	fmt.Println(`
		Welcome to the Prime Number Tool
		Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.
		`)
	prompt()
}

func prompt() {
	fmt.Print("-> ")
}

func isPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition!", n)
	}
	if n < 0 {
		return false, "Negative number are not prime, by definition!"
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number because it is divisible by %d!", n, i)
		}
	}
	return true, fmt.Sprintf("%d is a prime number!", n)
}
