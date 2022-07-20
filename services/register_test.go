package services

import (
	"os"
	"testing"

	"github.com/CryptoNerdES/cn.example.cryptobot.dca/models"
	"github.com/stretchr/testify/assert"
)

func TestCreatedRegisterFile(t *testing.T) {
	RegisterFileName = "test.json"
	f, err := getRegisterFile()
	if err != nil {
		t.Fatal(err)
	}

	assert.FileExists(t, f.Name())

	defer os.Remove(f.Name())
}

func TestReturnEmptyMovements(t *testing.T) {
	defer CleanMovements()
	RegisterFileName = "test.json"
	movements, err := GetMovements()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 0, len(movements))
}

func TestReturnMovements(t *testing.T) {
	defer CleanMovements()
	RegisterFileName = "test.json"
	newMovement := models.Movement{Token: "BTC", Amount: "0,00467", EntryPrice: "23897.23"}
	err := AddNewRow(newMovement)
	if err != nil {
		t.Fatal(err)
	}

	movements, err := GetMovements()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(movements))
}

func TestAddedNewRowToRegisterFile(t *testing.T) {
	defer CleanMovements()
	// variables
	RegisterFileName = "test.json"
	newMovement := models.Movement{Token: "BTC", Amount: "0,00467", EntryPrice: "23897.23"}

	// execute command
	for i := 0; i < 3; i++ {
		err := AddNewRow(newMovement)
		if err != nil {
			t.Fatal(err)
		}
	}

	testMovements, err := GetMovements()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 3, len(testMovements))
}
