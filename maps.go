package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":   "https://www.google.com",
		"Facebook": "https://www.facebook.com",
		"Amazon Web Services": "https://aws.amazon.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	websites["Twitter"] = "https://www.twitter.com"
	fmt.Println(websites["Twitter"])

	delete(websites, "Facebook")
	fmt.Println(websites)
}