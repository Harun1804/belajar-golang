package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":   "https://www.google.com",
		"Facebook": "https://www.facebook.com",
		"Amazon Web Services": "https://aws.amazon.com",
	}
	websites["Twitter"] = "https://www.twitter.com"
	delete(websites, "Facebook")

	for i, val := range websites {
		fmt.Printf("%s: %s\n", i, val)
	}
}