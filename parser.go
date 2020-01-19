package main

import (
    "strings"
    "errors"
    "strconv"
)

func validateAmount(val int64) bool {
    return val > 0
}

type UpdateData struct {
    val int64
}

// TODO: Parse data to data-structure fields
func ParseUpdateAmountMsg(s string) (UpdateData, error) {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
        return UpdateData{0},
               errors.New("Wrong string format, expected 2 tokens separated with space")
	}
	amount, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
        return UpdateData{0},
               errors.New("Can not parse amount, second token must be integer")
	}
    if !validateAmount(amount) {
        return UpdateData{0},
               errors.New("Amount must be positive number")
    }
    return UpdateData{amount}, nil
}
