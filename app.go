package main

import "fmt"

type TaxRate struct {
	base float64
	rate float64
}

func main() {
	taxLadder := []TaxRate{
		{
			base: 150_000,
			rate: 0.05,
		},
		{
			base: 300_000,
			rate: 0.10,
		},
		{
			base: 500_000,
			rate: 0.15,
		},
		{
			base: 1_000_000,
			rate: 0.20,
		},
		{
			base: 2_000_000,
			rate: 0.25,
		},
		{
			base: 5_000_000,
			rate: 0.30,
		},
	}
	calculateTaxSalary(&taxLadder)
}

func calculateTaxSalary(tax *[]TaxRate) {
	for _, tax := range *tax {
		var taxSalary = tax.base / 12
		var taxLostPerSalary = taxSalary * tax.rate
		var taxLeft = taxSalary - taxLostPerSalary
		fmt.Printf("Base = %.2f Rate %.2f = Salary %.2f, Lost %.2f,  Left %.2f\n",
			tax.base, tax.rate, taxSalary, taxLostPerSalary, taxLeft,
		)
	}
}
