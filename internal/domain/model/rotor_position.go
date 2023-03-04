package model

import (
	"fmt"
	"strconv"
)

type RotorPosition uint

func NewRotorPosition(position int) (RotorPosition, error) {
	if position == 0 || position > 26 {
		return RotorPosition(0), fmt.Errorf("%d is an invalid enigma rotor position value", position)
	}
	return RotorPosition(position), nil
}

func (r RotorPosition) String() string {
	return strconv.Itoa(int(r))
}
