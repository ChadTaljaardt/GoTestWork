package tempdir1

import (
	"fmt"

	"github.com/ChadTaljaardt/GoTestWork/tempdir2"
)

func CleanUpErrors() {
	// Some Error Occured Somewhere
	fmt.Println("Deleting Instance")
	// Running the delete instance function
	tempdir2.DeleteInstance()
}
