package model

import (
	"errors"
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
		return PlugboardCable{}, fmt.Errorf("'%s' is an invalid enigma plugboard cable value", value)
	}

	value = strings.ToUpper(value)
	from, err := validatePlugboardCableValue(value[0])
	if err != nil {
		return PlugboardCable{}, err
	}

	to, err := validatePlugboardCableValue(value[1])
	if err != nil {
		return PlugboardCable{}, err
	}

	if from == to {
		return PlugboardCable{}, errors.New("cannot repeat values in a enigma plugboard cable value")
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

type PlugboardCables []PlugboardCable

func NewPlugboardCables(value string) (PlugboardCables, error) {
	if value == "" {
		return PlugboardCables{}, nil
	}
	values := strings.Split(value, ",")
	var cables []PlugboardCable
	controlMap := make(map[byte]bool)
	for _, pair := range values {
		cable, err := loadCable(controlMap, pair)
		if err != nil {
			return PlugboardCables{}, err
		}
		cables = append(cables, cable)
	}
	return cables, nil
}

func loadCable(controlMap map[byte]bool, pair string) (PlugboardCable, error) {
	cable, err := NewPlugboardCable(pair)
	if err != nil {
		return PlugboardCable{}, err
	}
	if err := repeatControl(controlMap, cable.from); err != nil {
		return PlugboardCable{}, err
	}
	if err := repeatControl(controlMap, cable.to); err != nil {
		return PlugboardCable{}, err
	}
	return cable, nil
}

func repeatControl(controlMap map[byte]bool, value byte) error {
	_, exists := controlMap[value]
	if exists {
		return fmt.Errorf("cannot repeat the same value '%c' in different cables", value)
	}
	controlMap[value] = true
	return nil
}

func (p PlugboardCables) String() string {
	s := make([]string, len(p))
	for i, v := range p {
		s[i] = v.String()
	}
	return "(" + strings.Join(s, ",") + ")"
}
