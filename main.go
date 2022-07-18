package main

import (
	"log"

	"github.com/CryptoNerdES/cn.example.cryptobot.dca/services"
)

func main() {
	client := services.GetBinanceClient()

	// check Binance API is up
	if err := services.Health(client); err != nil {
		log.Printf("Binance API: %s", err.Error())
	}
	log.Println("Binance API is working right!")

	// check last steps
	movements, err := services.GetMovements()
	if err != nil {
		log.Printf("GetMovements: %s", err.Error())
	}
	log.Println("GetMovements:", movements)
}
