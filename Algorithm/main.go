package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//==========================================================LOGIC================================================

// findCombinations is the main function to find combinations
func findCombinations(l, t int) [][]int {
	ch := make(chan []int)
	results := [][]int{}

	// Start a goroutine to generate combinations
	go func() {
		defer close(ch)
		generateCombinations([]int{}, 1, l, t, ch)
	}()

	// Collect results from the channel
	for combination := range ch {
		results = append(results, combination)
	}

	return results
}

// generateCombinations generates combinations using recursion
func generateCombinations(current []int, start, length, total int, ch chan []int) {

	// Base case: if the combination is valid
	if len(current) == length {
		sum := 0
		//summarize combinations
		for _, num := range current {
			sum += num
		}
		//check total
		if sum == total {
			// Send the combination to the channel
			ch <- append([]int{}, current...)
		}
		return
	}

	// Generate combinations recursively
	for i := start; i <= 9; i++ {
		generateCombinations(append(current, i), i+1, length, total, ch)
	}
}

//====================================================HANDLER AND VALIDATION==========================================

// ExitHandler checks if the user wants to exit the program
func exitHandler(input string) bool {
	return strings.ToLower(input) == "done"
}

// UserInput Validation checks if the input is a valid positive integer
func validateInput(input string) (int, bool) {
	value, err := strconv.Atoi(input)
	if err != nil || value < 1 {
		return 0, false
	}
	return value, true
}

// Main function for continuous input
func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Combination Finder")
	fmt.Println("===================")
	fmt.Println("Enter `done` anytime to quit the program.")

	for {
		// Prompt for length
		fmt.Print("\nEnter the length of the combination (l): ")
		inputL, _ := reader.ReadString('\n')
		inputL = strings.TrimSpace(inputL)

		// Exit Handler
		if exitHandler(inputL) {
			fmt.Println("Exiting the program. Goodbye!")
			break
		}

		// User input validation for length variable
		l, isValid := validateInput(inputL)
		if !isValid {
			fmt.Println("Invalid input for length. Please enter a positive integer.")
			continue
		}

		// Prompt for total
		fmt.Print("Enter the total sum of the combination (t): ")
		inputT, _ := reader.ReadString('\n')
		inputT = strings.TrimSpace(inputT)

		// Exit Handler
		if exitHandler(inputT) {
			fmt.Println("Exiting the program. Goodbye!")
			break
		}

		// User input validation for total variable
		t, isValid := validateInput(inputT)
		if !isValid {
			fmt.Println("Invalid input for total. Please enter a positive integer.")
			continue
		}

		// Find and display combinations
		combinations := findCombinations(l, t)
		if len(combinations) == 0 {
			fmt.Println("No combinations found.")
		} else {
			fmt.Println("Possible combinations:")
			for _, combination := range combinations {
				fmt.Println(combination)
			}
		}
	}
}
