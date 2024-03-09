package prices

import (
	"fmt"

	"natthan.com/go-play/iomanager"
	"natthan.com/go-play/utils"
)

type TaxIncludePriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]float64  `json:"tax_included_prices"`
}

func (job *TaxIncludePriceJob) LoadData() error {
	strings, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println("FILE FAIL", err)
		return err
	}

	prices, err := utils.StringsToFloats(strings)

	if err != nil {
		fmt.Println("PARSE STRING FAIL", err)
		return err
	}
	job.InputPrices = prices

	return nil
}

func (job *TaxIncludePriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}
	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	job.TaxIncludedPrices = result

	return job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludePriceJob {
	return &TaxIncludePriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
