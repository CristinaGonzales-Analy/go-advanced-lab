package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}

	if n == 0 {
		return 1, nil
	}

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	return result, nil
}

func IsPrime(n int) (bool, error) {
	if n < 2 {
		return false, errors.New("prime check requires number >= 2")
	}

	if n == 2 {
		return true, nil
	}

	if n%2 == 0 {
		return false, nil
	}

	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, errors.New("negative exponents not supported")
	}

	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}

	return result, nil
}

func MakeCounter(start int) func() int {
	count := start

	return func() int {
		count++
		return count
	}
}

func MakeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func MakeAccumulator(initial int) (func(int), func(int), func() int) {
	value := initial

	add := func(x int) {
		value += x
	}

	subtract := func(x int) {
		value -= x
	}

	get := func() int {
		return value
	}

	return add, subtract, get
}

func Apply(nums []int, operation func(int) int) []int {
	result := make([]int, len(nums))

	for i, v := range nums {
		result[i] = operation(v)
	}

	return result
}

func Filter(nums []int, predicate func(int) bool) []int {
	var result []int

	for _, v := range nums {
		if predicate(v) {
			result = append(result, v)
		}
	}

	return result
}

func Reduce(nums []int, initial int, operation func(int, int) int) int {
	acc := initial

	for _, v := range nums {
		acc = operation(acc, v)
	}

	return acc
}

func Compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

func ExploreProcess() {
	fmt.Println("=== Process Information ===")

	// A Process ID is a unique number assigned by the operating system
	// to identify a running program.
	pid := os.Getpid()

	// Parent process ID is the program that launched this one.
	ppid := os.Getppid()

	fmt.Println("Current Process ID:", pid)
	fmt.Println("Parent Process ID:", ppid)

	data := []int{1, 2, 3, 4, 5}

	// The slice header contains a pointer to the underlying array, length, capacity. Also, functions are passed by copying
	//Element address refers to the actual data items, size of type, points to original array

	fmt.Printf("Slice header address: %p\n", &data)
	fmt.Printf("First element address: %p\n", &data[0])

	// Process isolation is important because each process has protected memory,
	// which prevents one running program from directly reading or modifying
	// another program's memory addresses
	fmt.Println("Other processes cannot access these addresses due to process isolation.")
}

func main() {
	fmt.Println("Working...")
}
