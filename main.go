package main

import (
	// "fmt"

	"example.com/price-calculator/prices"
	"example.com/price-calculator/utils/cmdmanager"
	// "example.com/price-calculator/utils/filemanager"
)

func main() {
	taxRates := []float64{0.05, 0.07, 0.06, 0.04, 0.03, 0.02}

	for _, taxRate := range taxRates {
		// fileManager := filemanager.New("prices.txt", "result", fmt.Sprintf("result/result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		priceJob.Process()
	}
}