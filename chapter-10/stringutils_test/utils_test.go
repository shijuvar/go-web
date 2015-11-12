package stringutils_test

import (
	"fmt"
	"testing"
	"time"

	. "github.com/shijuvar/go-web/chapter-10/stringutils"
)

// Test case for the SwapCase function to execute in parallel
func TestSwapCaseInParallel(t *testing.T) {
	t.Parallel()
	// Delaying 1 second for the sake of demonstration
	time.Sleep(1 * time.Second)
	input, expected := "Hello, World", "hELLO, wORLD"
	result := SwapCase(input)

	if result != expected {

		t.Errorf("SwapCase(%q) == %q, expected %q", input, result, expected)
	}

}

// Test case for the Reverse function to execute in parallel
func TestReverseInParallel(t *testing.T) {
	t.Parallel()
	// Delaying 2 seconds for the sake of demonstration
	time.Sleep(2 * time.Second)
	input, expected := "Hello, World", "dlroW ,olleH"
	result := Reverse(input)

	if result != expected {

		t.Errorf("Reverse(%q) == %q, expected %q", input, result, expected)
	}

}

// Test case for the SwapCase function
func TestSwapCase(t *testing.T) {
	input, expected := "Hello, World", "hELLO, wORLD"
	result := SwapCase(input)

	if result != expected {

		t.Errorf("SwapCase(%q) == %q, expected %q", input, result, expected)
	}

}

// Test case for the Reverse function
func TestReverse(t *testing.T) {
	input, expected := "Hello, World", "dlroW ,olleH"
	result := Reverse(input)

	if result != expected {

		t.Errorf("Reverse(%q) == %q, expected %q", input, result, expected)
	}

}

//Benchmark for SwapCase function
func BenchmarkSwapCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SwapCase("Hello, World")
	}
}

//Benchmark for Reverse function
func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("Hello, World")
	}
}

//Example code for Reverse function
func ExampleReverse() {
	fmt.Println(Reverse("Hello, World"))
	// Output: dlroW ,olleH
}

//Example code for Reverse function
func ExampleSwapCase() {
	fmt.Println(SwapCase("Hello, World"))
	// Output: hELLO, wORLD
}
func TestLongRun(t *testing.T) {
	// Checks whether the short flag is provided
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}
	// Long running implementation goes here
	time.Sleep(5 * time.Second)
}
