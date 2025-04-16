package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGoRoutine(t *testing.T) {
	go GoRoutine("1")
	go GoRoutine("2")
	fmt.Println("Main Goroutine")

	// Wait for goroutines to finish (not a good practice, just for demonstration)
	time.Sleep(1 * time.Second)

}