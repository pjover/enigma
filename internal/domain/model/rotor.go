package model

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

func NewRotor(value string) (Rotor, error) {

	value = strings.ToUpper(value)
	for i, val := range stringValues {
		if val == value {
			return Rotor(i), nil
		}
	}
	return Rotor(0), fmt.Errorf("'%s' is an invalid enigma rotor value", value)
}

func (r Rotor) String() string {
	return stringValues[r]
}

type Rotors struct {
	first  Rotor
	second Rotor
	third  Rotor
}

func NewRotors(value string) (Rotors, error) {
	values, err := getValues(value)
	if err != nil {
		return Rotors{}, err
	}
	first, err := NewRotor(values[0])
	if err != nil {
		return Rotors{}, err
	}
	second, err := NewRotor(values[1])
	if err != nil {
		return Rotors{}, err
	}
	third, err := NewRotor(values[2])
	if err != nil {
		return Rotors{}, err
	}
	return Rotors{
		first:  first,
		second: second,
		third:  third,
	}, nil
}

func getValues(value string) ([]string, error) {
	values := strings.Split(value, ",")
	length := len(values)
	if length != 3 {
		return []string{}, fmt.Errorf("please define 3 values instead of %d at '%s'", length, value)
	}
	return values, nil
}

func (r Rotors) String() string {
	return fmt.Sprintf("rotors: %s,%s,%s",
		r.first.String(),
		r.second.String(),
		r.third.String(),
	)
}
