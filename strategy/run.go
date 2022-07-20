package strategy

import (
	"log"
	"math"
	"strconv"

	"github.com/CryptoNerdES/cn.example.cryptobot.dca/models"
)

func Run(lastMovements []models.Movement) {
	totalAmounts := getTotalAmount(lastMovements)
	log.Println(totalAmounts)
	averagePrice := getAveragePrice(lastMovements)
	log.Println(averagePrice)
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
