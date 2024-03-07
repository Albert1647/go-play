package prices

import (
	"fmt"

	"natthan.com/go-play/utils"
)

type TaxIncludePriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludePriceJob) getPriceFromFile() {
	strings, err := utils.GetStringsFromFile("price.txt")
	if err != nil {
		fmt.Println("FILE FAIL", err)
		return
	}
	prices, err := utils.StringsToFloats(strings)
	if err != nil {
		fmt.Println("PARSE STRING FAIL", err)
		return
	}
	job.InputPrices = prices
}

func (job *TaxIncludePriceJob) Process() {
	job.getPriceFromFile()

	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	fmt.Println("Cal:", result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludePriceJob {
	return &TaxIncludePriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
