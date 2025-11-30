// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-30
// Fileoverview: This program asks the user for their 10 best friends and displays them in order and reverse order

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// variables
	var bestFriends [10]string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please enter your 10 best friends in order from oldest friend to newest friend:")

	// get 10 friends from user
	for i := 0; i < 10; i++ {
		fmt.Printf("Enter friend %d: ", i+1)
		scanner.Scan()
		bestFriends[i] = strings.TrimSpace(scanner.Text())
	}

	// display friends in order
	fmt.Println("\nHere is your list of best friends in order of length of friendship starting from oldest friend to newest friend:")
	for i, friend := range bestFriends {
		if i == len(bestFriends)-1 {
			fmt.Printf("%s.\n", friend)
		} else {
			fmt.Printf("%s, ", friend)
		}
	}

	// display friends in reverse order
	fmt.Println("\nHere is your list of best friends in reverse order of length of friendship starting from newest friend to oldest friend:")
	for i := len(bestFriends) - 1; i >= 0; i-- {
		if i == 0 {
			fmt.Printf("%s.\n", bestFriends[i])
		} else {
			fmt.Printf("%s, ", bestFriends[i])
		}
	}

	fmt.Println("\nDone.")
}
