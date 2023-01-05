package rotor

import (
	"fmt"
	"strings"
)

type Rotor uint

const (
	I Rotor = iota
	II
	III
	IV
	V
	VI
	VII
	VIII
)

var stringValues = []string{
	"I",
	"II",
	"III",
	"IV",
	"V",
	"VI",
	"VII",
	"VIII",
}

func (r Rotor) String() string {
	return stringValues[r]
}

func NewRotor(value string) (Rotor, error) {

	value = strings.ToUpper(value)
	for i, val := range stringValues {
		if val == value {
			return Rotor(i), nil
		}
	}
	return Rotor(0), fmt.Errorf("'%s' is an invalid enigma rotor value", value)
}

func (r Rotor) Format() string {
	return fmt.Sprintf("rotor %s", stringValues[r])
}
