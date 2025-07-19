package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) loadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file!")
		fmt.Println(err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))
	for i, line := range lines {
		price, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Converting price to float failed!")
			fmt.Println(err)
			file.Close()
			return
		}
		prices[i] = price
	}

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