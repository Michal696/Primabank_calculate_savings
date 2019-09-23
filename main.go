package main

import "fmt"

func addMonthMoney(moneySumInput, addedMoneyEachMonth float64) float64 {
	return moneySumInput + addedMoneyEachMonth
}

func addMonthInterest(moneySumInput, interestRatePerMonth float64) float64 {
	return moneySumInput * interestRatePerMonth
}

func calculateSavings(addedMoneyEachMonth, interestRate float64, numOfMonths int) {
	interestRatePerMonth := interestRate / 12
	moneySumInput := 0.0
	interestMoney := 0.0

	for i := 0; i < numOfMonths; i++ {
		moneySumInput = addMonthMoney(moneySumInput, addedMoneyEachMonth)
		interestMoney = interestMoney + addMonthInterest(moneySumInput, interestRatePerMonth)
	}

	fmt.Println("###############################################")
	fmt.Println("Number of months: ", numOfMonths)
	fmt.Println("Money input each month: ", addedMoneyEachMonth)
	fmt.Println("Money Input Sum: ", moneySumInput)
	fmt.Println("Interest money: ", interestMoney)
	fmt.Println("Come back: ", moneySumInput+interestMoney)
	fmt.Println("###############################################")

}

// student 10-30
// normal 10-50

func main() {
	fmt.Println("Hello World")
	fmt.Println("")

	calculateSavings(30.0, 0.05, 12)
	calculateSavings(30.0, 0.05, 18)
	calculateSavings(30.0, 0.05, 24)

	calculateSavings(50.0, 0.05, 12)
	calculateSavings(50.0, 0.05, 18)
	calculateSavings(50.0, 0.05, 24)

	fmt.Println("")
	fmt.Println("End Of World")
}
