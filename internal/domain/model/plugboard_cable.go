package model

import (
	"fmt"
	"strings"
)

type PlugboardCable struct {
	from byte
	to   byte
}

func (p PlugboardCable) From() byte {
	return p.from
}
func (p PlugboardCable) To() byte {
	return p.to
}

func NewPlugboardCable(value string) (PlugboardCable, error) {
	if len(value) != 2 {
		return PlugboardCable{from: 0, to: 0}, fmt.Errorf("'%s' is an invalid enigma plugboard cable value", value)
	}

	value = strings.ToUpper(value)
	from, err := validatePlugboardCableValue(value[0])
	if err != nil {
		return PlugboardCable{from: 0, to: 0}, err
	}

	to, err := validatePlugboardCableValue(value[1])
	if err != nil {
		return PlugboardCable{from: 0, to: 0}, err
	}

	return PlugboardCable{from: from, to: to}, nil
}

func validatePlugboardCableValue(value byte) (byte, error) {
	if value < 'A' || value > 'Z' {
		return 0, fmt.Errorf("'%c' is an invalid enigma plugboard value", value)
	}
	return value, nil
}

func (p PlugboardCable) String() string {
	return fmt.Sprintf("%c-%c", p.from, p.to)
}

func (p PlugboardCable) Format() string {
	return fmt.Sprintf("plugboard cable %s-", p.String())
}
