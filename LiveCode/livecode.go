package main

import "fmt"

// Main function for continuous input
func main() {

	var n int

	// Get user input
	fmt.Print("Enter an integer: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Invalid input. Please enter an integer.")
		return
	}

	//new pattern
	for i := n; i >= 1; i-- {
		
		line := fmt.Sprintf("%*s", n-i+1, "")
		for j := 1; j <= i; j++ {
			line += "*"
			if j < i {
				line += " "
			}
		}

		fmt.Println(line)
	}

}
