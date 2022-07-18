package services

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/CryptoNerdES/cn.example.cryptobot.dca/models"
	"github.com/adshao/go-binance/v2"
)

var (
	ErrBinanceAPIIsDown error = errors.New("binance API is not up")
)

const (
	BinanceHealthPath string = "/sapi/v1/system/status"
)

// Check if conexión with Binance is okey
func Health(client *binance.Client) error {
	var healthResponse models.HealthResponse

	response, err := http.Get(getBinanceURL(BinanceHealthPath))
	if err != nil {
		return err
	}
	if err := json.NewDecoder(response.Body).Decode(&healthResponse); err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK || healthResponse.Status != 0 {
		return ErrBinanceAPIIsDown
	}

	return nil
}
