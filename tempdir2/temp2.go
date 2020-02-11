package tempdir2

import (
	"fmt"

	"github.com/ChadTaljaardt/GoTestWork/tempdir1"
)

func CreateInstance() {
	fmt.Println("Creating Instance")
	fmt.Println("Instance Created")

	// Some Random Error Occured somewhere so run the cleanup method
	tempdir1.CleanUpErrors()
}

func DeleteInstance() {
	fmt.Println("Deleting Instance")
}
