package model

import (
	"fmt"
	"strconv"
)

type RotorPosition uint

func (r RotorPosition) String() string {
	return strconv.Itoa(int(r))
}

func NewRotorPosition(value uint) (RotorPosition, error) {

	if value == 0 || value > 26 {
		return RotorPosition(0), fmt.Errorf("'%d' is an invalid enigma rotor position value", value)
	}
	return RotorPosition(value), nil
}

func (r RotorPosition) Format() string {
	return fmt.Sprintf("rotor position %d", r)
}
