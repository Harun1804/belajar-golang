package prices

import (
	"fmt"

	"example.com/price-calculator/utils/converter"
	"example.com/price-calculator/utils/errorhandling"
	"example.com/price-calculator/utils/filemanager"
)

type TaxIncludedPriceJob struct {
	IO              filemanager.FileManager
	TaxRate         float64
	InputPrices     []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) loadData() {
	lines, err := job.IO.ReadFile()
	errorhandling.HandleError(err, "Failed to load prices")

	prices, err := converter.StringsToFloats(lines)
	errorhandling.HandleError(err, "Failed to convert prices to float64")

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.loadData()

	total := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		total[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = total
	err := job.IO.WriteJson(job)
	errorhandling.HandleError(err, "Failed to write tax included prices to JSON")
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IO:           fm,
		TaxRate:     taxRate,
		InputPrices: []float64{7.0, 1.0, 5.0, 3.0, 6.0, 4.0},
	}
}