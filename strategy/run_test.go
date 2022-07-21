package strategy

import (
	"testing"

	"github.com/CryptoNerdES/cn.example.cryptobot.dca/models"
	"github.com/CryptoNerdES/cn.example.cryptobot.dca/services"
	"github.com/stretchr/testify/assert"
)

func TestReturnEntryPriceAverageOK(t *testing.T) {
	defer services.CleanMovements()
	services.RegisterFileName = "test.json"
	firstMovement := models.Movement{Token: "BTC", Amount: "0.00467", EntryPrice: "23897.23"}
	secondMovement := models.Movement{Token: "BTC", Amount: "0.00467", EntryPrice: "25632.45"}
	err, err2 := services.AddNewRow(firstMovement), services.AddNewRow(secondMovement)
	if err != nil || err2 != nil {
		t.Fatal(err)
	}

	movements, err := services.GetMovements()
	if err != nil {
		t.Fatal(err)
	}

	entryPriceAverage := getAveragePrice(movements)

	assert.Equal(t, entryPriceAverage, 24764.84)
}

func TestTotalAmountIsOK(t *testing.T) {
	defer services.CleanMovements()
	services.RegisterFileName = "test.json"
	firstMovement := models.Movement{Token: "BTC", Amount: "0.00467", EntryPrice: "23897.23"}
	secondMovement := models.Movement{Token: "BTC", Amount: "0.00467", EntryPrice: "25632.45"}
	err, err2 := services.AddNewRow(firstMovement), services.AddNewRow(secondMovement)
	if err != nil || err2 != nil {
		t.Fatal(err)
	}

	movements, err := services.GetMovements()
	if err != nil {
		t.Fatal(err)
	}

	totalAmount := getTotalAmount(movements)

	assert.Equal(t, totalAmount, 0.00934)
}


func TestCalculatePercentageIncrease(t *testing.T) {
	averagePrice := 24764.84
	currentPrice := 25632.45
	percentageIncrease := percentageIncreaseCalculator(averagePrice, currentPrice)
	assert.Equal(t, percentageIncrease, 3.5)
}