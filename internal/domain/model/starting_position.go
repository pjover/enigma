package model

import (
	"fmt"
	"strconv"
)

type StartingPosition uint

func NewStartingPosition(value string) (StartingPosition, error) {
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("'%s' is an invalid enigma rotor starting position number", value)
	}
	if number == 0 || number > 26 {
		return StartingPosition(0), fmt.Errorf("'%s' is an invalid enigma rotor starting position value", value)
	}
	return StartingPosition(number), nil
}

func (r StartingPosition) String() string {
	return strconv.Itoa(int(r))
}
