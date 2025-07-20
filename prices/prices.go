package prices

import (
	"fmt"

	"example.com/price-calculator/utils/converter"
	"example.com/price-calculator/utils/iomanager"
)

type TaxIncludedPriceJob struct {
	IO              iomanager.IOManager				`json:"-"`
	TaxRate         float64                  	`json:"tax_rate"`
	InputPrices     []float64                	`json:"input_prices"`
	TaxIncludedPrices map[string]string      	`json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) loadData() error {
	lines, err := job.IO.ReadFile()
	if err != nil {
		return err
	}

	prices, err := converter.StringsToFloats(lines)
	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	err := job.loadData()
	if err != nil {
		errorChan <- err
		return
	}

	total := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		total[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = total
	job.IO.WriteJson(job)
	doneChan <- true
}

func NewTaxIncludedPriceJob(io iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IO:           io,
		TaxRate:     taxRate,
		InputPrices: []float64{7.0, 1.0, 5.0, 3.0, 6.0, 4.0},
	}
}