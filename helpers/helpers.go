package helpers

import "fmt"

func DisplayError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}