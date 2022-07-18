package services

import (
	"fmt"
	"log"
	"os"

	"github.com/adshao/go-binance/v2"
)

var (
	BinanceEndpoint string = "https://api.binance.com"
)

func getBinanceURL(path string) string {
	return fmt.Sprintf("%s%s", BinanceEndpoint, path)
}

func GetBinanceClient() *binance.Client {
	client := binance.NewClient(os.Getenv("BINANCE_API_KEY"), os.Getenv("BINANCE_API_SECRET"))
	if client == nil {
		log.Fatal("Binance client is null, something went wrong")
	}

	return client
}