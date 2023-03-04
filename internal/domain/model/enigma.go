package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Enigma struct {
	rotors          Rotors
	plugboardCables PlugboardCables
}

func (e Enigma) Rotors() Rotors {
	return e.rotors
}

func (e Enigma) PlugboardCables() PlugboardCables {
	return e.plugboardCables
}

func NewEnigmaMachine(value string) (Enigma, error) {
	values := strings.Split(value, "#")
	if len(values) != 4 {
		return Enigma{}, errors.New("error parsing enigma values, must define 4 groups separated by '#'")
	}

	rotorNumbers, err := splitNumbers(values[0], "rotor numbers")
	if err != nil {
		return Enigma{}, err
	}

	ringSettings, err := splitNumbers(values[1], "ring settings")
	if err != nil {
		return Enigma{}, err
	}

	startingPositions, err := splitNumbers(values[2], "starting positions")
	if err != nil {
		return Enigma{}, err
	}

	plugboardCables, err := NewPlugboardCables(values[3])
	if err != nil {
		return Enigma{}, err
	}

	rotors, err := NewRotors(rotorNumbers, ringSettings, startingPositions)
	if err != nil {
		return Enigma{}, err
	}

	return Enigma{
		rotors:          rotors,
		plugboardCables: plugboardCables,
	}, nil
}

func splitNumbers(value string, name string) ([3]uint, error) {
	values := strings.Split(value, ",")
	length := len(values)
	if length != 3 {
		return [3]uint{}, fmt.Errorf("please define 3 %s values instead of %d at '%s'", name, length, value)
	}
	var numbers [3]uint
	for i, val := range values {
		number, err := strconv.Atoi(val)
		if err != nil {
			return [3]uint{}, fmt.Errorf("'%s' is an invalid number", value)
		}
		numbers[i] = uint(number)
	}
	return numbers, nil
}

func (e Enigma) String() string {
	return fmt.Sprintf("%s %v",
		e.rotors.String(),
		e.plugboardCables.String(),
	)
}
