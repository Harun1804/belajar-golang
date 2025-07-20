package main

import (
	"fmt"

	"example.com/price-calculator/prices"
	// "example.com/price-calculator/utils/cmdmanager"
	"example.com/price-calculator/utils/filemanager"
)

func main() {
	taxRates := []float64{0.05, 0.07, 0.06, 0.04, 0.03, 0.02}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error)
		fileManager := filemanager.New("prices.txt", "result", fmt.Sprintf("result/result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fileManager, taxRate)
		go priceJob.Process(doneChans[i], errorChans[i])
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			fmt.Printf("Error processing job %d: %v\n", index, err)
		case <-doneChans[index]:
			fmt.Printf("Job %d completed successfully\n", index)
		}
	}
}