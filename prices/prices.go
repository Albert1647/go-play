package prices

type TaxIncludePriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludePriceJob {
	return &TaxIncludePriceJob{
		InputPrices: []float64{0, 0.07, 0.1, 0.15},
		TaxRate:     taxRate,
	}
}
