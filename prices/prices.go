package prices

import "fmt"

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) Process() {
	total := make(map[string]float64)

	for _, price := range job.InputPrices {
		total[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	fmt.Println(total)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{7.0, 1.0, 5.0, 3.0, 6.0, 4.0},
	}
}