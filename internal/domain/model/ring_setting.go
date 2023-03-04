package model

import (
	"fmt"
	"strconv"
)

type RingSetting uint

func NewRingSetting(ringSetting uint) (RingSetting, error) {
	if ringSetting == 0 || ringSetting > 26 {
		return RingSetting(0), fmt.Errorf("'%d' is an invalid enigma rotor ring setting value", ringSetting)
	}
	return RingSetting(ringSetting), nil
}

func (r RingSetting) String() string {
	return strconv.Itoa(int(r))
}
