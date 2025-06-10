package main

import (
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0.05, 0.07, 0.06, 0.04, 0.03, 0.02}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}
}