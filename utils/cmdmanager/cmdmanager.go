package cmdmanager

import "fmt"

type CMDManager struct{}

func (cmd CMDManager) ReadFile() ([]string, error) {
	fmt.Println("Please emter your prices. Confirm every price with Enter")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)
		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMDManager) WriteJson(data interface{}) error {
	fmt.Println("Writing JSON data:", data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
