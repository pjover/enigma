package model

import (
	"fmt"
	"strconv"
)

type RotorRing uint

func NewRotorRing(value uint) (RotorRing, error) {

	if value == 0 || value > 26 {
		return RotorRing(0), fmt.Errorf("'%d' is an invalid enigma rotor ring value", value)
	}
	return RotorRing(value), nil
}

func (r RotorRing) String() string {
	return strconv.Itoa(int(r))
}

func (r RotorRing) Format() string {
	return fmt.Sprintf("rotor ring %s", r.String())
}
