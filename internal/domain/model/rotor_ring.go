package model

import (
	"fmt"
	"strconv"
)

type RotorRing uint

func NewRotorRing(value string) (RotorRing, error) {
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("'%s' is an invalid enigma rotor ring number", value)
	}
	if number == 0 || number > 26 {
		return RotorRing(0), fmt.Errorf("'%s' is an invalid enigma rotor ring value", value)
	}
	return RotorRing(number), nil
}

func (r RotorRing) String() string {
	return strconv.Itoa(int(r))
}
