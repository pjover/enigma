package model

import (
	"fmt"
	"strconv"
)

type RotorPosition uint

func NewRotorPosition(value string) (RotorPosition, error) {
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("'%s' is an invalid enigma rotor position number", value)
	}
	if number == 0 || number > 26 {
		return RotorPosition(0), fmt.Errorf("'%s' is an invalid enigma rotor position value", value)
	}
	return RotorPosition(number), nil
}

func (r RotorPosition) String() string {
	return strconv.Itoa(int(r))
}
