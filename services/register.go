package services

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/CryptoNerdES/cn.example.cryptobot.dca/models"
)

var (
	RegisterFileName string = "register.json"
)

func CleanMovements() error {
	return os.Remove(RegisterFileName)
}

func getRegisterFile() (*os.File, error) {
	f, err := os.OpenFile(RegisterFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return f, nil
}

func GetMovements() ([]models.Movement, error) {
	movements := []models.Movement{}
	registerFile, err := getRegisterFile()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(registerFile.Name())
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}

	err = json.Unmarshal(data, &movements)
	if err != nil {
		return nil, err
	}

	return movements, nil
}

func AddNewRow(movement models.Movement) error {
	registerFile, err := getRegisterFile()
	if err != nil {
		return err
	}
	movements, err := GetMovements()
	if err != nil {
		return err
	}

	movements = append(movements, movement)
	data, err := json.MarshalIndent(movements, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(registerFile.Name(), data, 0644)
	if err != nil {
		return err
	}
	defer registerFile.Close()

	return nil
}
