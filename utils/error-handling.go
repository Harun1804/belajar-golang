package utils

import "fmt"

func HandleError(err error, message string) bool {
	if err != nil {
		fmt.Println(message)
		fmt.Println(err)
		return true
	}
	return false
}