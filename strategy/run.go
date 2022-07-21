package strategy

import (
	"log"
	"math"
	"strconv"

	"github.com/CryptoNerdES/cn.example.cryptobot.dca/models"
	"github.com/CryptoNerdES/cn.example.cryptobot.dca/services"
)

func Run(lastMovements []models.Movement) error {
	totalAmounts := getTotalAmount(lastMovements)
	log.Println(totalAmounts)
	averagePrice := getAveragePrice(lastMovements)
	log.Println(averagePrice)
	currentPrice, err := services.GetCurrentPriceOfAToken("BTC")
	if err != nil {
		return err
	}
	percentageIncrease := percentageIncreaseCalculator(averagePrice, currentPrice)
	log.Println(percentageIncrease)

	return nil
}

func percentageIncreaseCalculator(averagePrice float64, currentPrice float64) float64 {
	percentageIncrease := ((currentPrice - averagePrice) / averagePrice) *100
	return math.Round(percentageIncrease*100) / 100
}

func getTotalAmount(movements []models.Movement) float64 {
	var totalAmount float64
	for _, movement := range movements {
		s, err := strconv.ParseFloat(movement.Amount, 64)
		if err != nil {
			log.Fatal(err)
		}
		totalAmount += s
	}
	return totalAmount
}

func getAveragePrice(movements []models.Movement) float64 {
	var totalPrice float64
	for _, movement := range movements {
		price, err := strconv.ParseFloat(movement.EntryPrice, 64)
		if err != nil {
			log.Fatal(err)
		}
		totalPrice += price
	}
	average := totalPrice / float64(len(movements))
	return math.Round(average*100) / 100
}
