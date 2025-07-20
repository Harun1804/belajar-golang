package prices

import (
	"fmt"
	"example.com/price-calculator/utils"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) loadData() {
	lines, err := utils.ReadFile("prices.txt")
	utils.HandleError(err, "Failed to load prices")

	prices, err := utils.StringsToFloats(lines)
	utils.HandleError(err, "Failed to convert prices to float64")

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.loadData()

	total := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		total[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(total)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{7.0, 1.0, 5.0, 3.0, 6.0, 4.0},
	}
}