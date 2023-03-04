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

	rotorNumbers, err := parseRomanNumbers(values[0], "rotor number")
	if err != nil {
		return Enigma{}, err
	}

	ringSettings, err := parseArabicNumbers(values[1], "ring setting")
	if err != nil {
		return Enigma{}, err
	}

	startingPositions, err := parseArabicNumbers(values[2], "starting position")
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

func parseRomanNumbers(value string, name string) ([3]uint, error) {
	return parseNumbers(value, name, romanNumberToInt)
}

func parseArabicNumbers(value string, name string) ([3]uint, error) {
	return parseNumbers(value, name, strconv.Atoi)
}

func romanNumberToInt(value string) (int, error) {
	switch value {
	case "I":
		return 1, nil
	case "II":
		return 2, nil
	case "III":
		return 3, nil
	case "IV":
		return 4, nil
	case "V":
		return 5, nil
	case "VI":
		return 6, nil
	case "VII":
		return 7, nil
	case "VIII":
		return 8, nil
	default:
		return 0, errors.New("invalid roman number")
	}
}

func parseNumbers(value string, name string, converter func(string) (int, error)) ([3]uint, error) {
	values, err := splitNumbers(value, name)
	if err != nil {
		return [3]uint{}, err
	}

	var numbers [3]uint
	for i, val := range values {
		number, err := converter(val)
		if err != nil {
			return [3]uint{}, fmt.Errorf("%s is an invalid enigma %s", val, name)
		}
		numbers[i] = uint(number)
	}
	return numbers, nil
}

func splitNumbers(value string, name string) ([]string, error) {
	values := strings.Split(value, ",")
	length := len(values)
	if length != 3 {
		return []string{}, fmt.Errorf("please define 3 %s values instead of %d at '%s'", name, length, value)
	}
	return values, nil
}

func (e Enigma) String() string {
	return fmt.Sprintf("%s %v",
		e.rotors.String(),
		e.plugboardCables.String(),
	)
}
