package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	averageWinRate := 0.45
	averageRR := 2.5
	riskPerTrade := 0.01
	startingDeposit := 10000.0
	tradesTaken := 16
	exchangeFeePercent := 0.075 / 100

	deposit := startingDeposit
	totalWins := 0
	totalLosses := 0

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Printf("Starting Deposit: $%.2f\n", startingDeposit)

	for i := 1; i <= tradesTaken; i++ {
		riskAmount := deposit * riskPerTrade
		isWin := rng.Float64() < averageWinRate

		var pnl float64
		if isWin {
			totalWins++
			pnl = riskAmount * averageRR
		} else {
			totalLosses++
			pnl = -riskAmount
		}

		fee := deposit * exchangeFeePercent
		pnl -= fee

		deposit += pnl

		status := "Win"
		if !isWin {
			status = "Loss"
		}
		fmt.Printf("#%d. %s | P&L: %.2f | Deposit after trade: %.2f\n", i, status, pnl, deposit)
	}

	profitLoss := deposit - startingDeposit

	fmt.Printf("\nTotal Trades: %d | Wins: %d\n", tradesTaken, totalWins)
	fmt.Printf("Starting Deposit: $%.2f\n", startingDeposit)
	fmt.Printf("Final Deposit: $%.2f\n", deposit)
	fmt.Printf("Profit/Loss: $%.2f\n", profitLoss)
}
