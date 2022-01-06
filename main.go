package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Returns a range of integers
// Input: makeRange(1, 9)
// Output: [1, 2, 3, 4, 5, 6, 7, 8, 9]
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// Removes an element from a slice
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func validateInput(s, e string) (int, error) {
	fmt.Println(s)
	var str string
	fmt.Scanln(&str)

	// convert str (string variable) to x (int variable)
	x, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Conversion error with non integer parameter:", str)
		os.Exit(1)
	}
	if x <= 0 {
		return 0, errors.New(e)
	}
	return x, nil
}

// First Solution: findWinnerSlices()
// Returns the winner of Josephus Problem
// with n the number of participants
// and m the step chosen
// via slices
func findWinnerSlices(n, m int) int {
	players := makeRange(1, n)
	index := m - 1
	for len(players) > 1 {
		players = remove(players, index)
		index = (index + m - 1) % len(players)
	}
	return players[0]
}

// Second Solution: findWinnerRecursion()
// Returns the winner of Josephus Problem
// with n the number of participants
// and m the step chosen
// via recursive method
func findWinnerRecursion(n, m int) int {
	if n == 1 {
		return 1
	} else {
		return (findWinnerRecursion(n-1, m)+m-1)%n + 1
	}
}

func main() {

	s1 := "Choose a number of participants (positive integer): "
	e1 := "choose a positive number of participants."

	s2 := "Choose a step (positive integer): "
	e2 := "choose a positive number for step."

	n, err1 := validateInput(s1, e1)
	m, err2 := validateInput(s2, e2)

	if err1 != nil {
		fmt.Printf("Value error: %v\n", err1)
		return
	} else if err2 != nil {
		fmt.Printf("Value error: %v\n", err2)
		return
	}

	fmt.Println("================  First Solution (Slices) ================")
	result := findWinnerSlices(n, m)
	fmt.Printf("Survivor of Josephus Problem with %d participants and step of %d: Participant n°%d\n\n", n, m, result)

	fmt.Println("==============  Second Solution (Recursion) ==============")
	result = findWinnerRecursion(n, m)
	fmt.Printf("Survivor of Josephus Problem with %d participants and step of %d: Participant n°%d\n", n, m, result)
}
