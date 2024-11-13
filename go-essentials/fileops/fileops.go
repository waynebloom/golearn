package fileops

import (
  "errors"
  "fmt"
  "os"
  "strconv"
)

func WriteFloat(fileName string, balance float64) {
  balanceString := fmt.Sprint(balance)
  os.WriteFile(fileName, []byte(balanceString), 0644)
}

func ReadFloat(fileName string) (float64, error) {
  bytes, err := os.ReadFile(fileName)

  if err != nil {
    return 1000, errors.New(fileName + " was not found")
  }

  balanceString := string(bytes)
  balance, err := strconv.ParseFloat(balanceString, 64)

  if err != nil {
    return 1000, errors.New(fileName + " contained invalid data")
  }

  return balance, nil
}

