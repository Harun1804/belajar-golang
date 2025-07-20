package prices

import (
	"fmt"

	"example.com/price-calculator/utils/converter"
	"example.com/price-calculator/utils/errorhandling"
	"example.com/price-calculator/utils/iomanager"
)

type TaxIncludedPriceJob struct {
	IO              iomanager.IOManager				`json:"-"`
	TaxRate         float64                  	`json:"tax_rate"`
	InputPrices     []float64                	`json:"input_prices"`
	TaxIncludedPrices map[string]string      	`json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) loadData() {
	lines, err := job.IO.ReadFile()
	errorhandling.HandleError(err, "Failed to load prices")

	prices, err := converter.StringsToFloats(lines)
	errorhandling.HandleError(err, "Failed to convert prices to float64")

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool) {
	job.loadData()

	total := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		total[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = total
	err := job.IO.WriteJson(job)
	errorhandling.HandleError(err, "Failed to write tax included prices to JSON")
	doneChan <- true
}

func NewTaxIncludedPriceJob(io iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IO:           io,
		TaxRate:     taxRate,
		InputPrices: []float64{7.0, 1.0, 5.0, 3.0, 6.0, 4.0},
	}
}