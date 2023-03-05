package model

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Enigma struct {
	leftRotor       Rotor
	middleRotor     Rotor
	rightRotor      Rotor
	plugboardCables PlugboardCables
}

func NewEnigmaMachine(text string) (Enigma, error) {
	values, err := splitText(text)
	if err != nil {
		return Enigma{}, err
	}

	var rotors []Rotor
	for i := 0; i < 3; i++ {
		rotor, err := parseRotor(values[i])
		if err != nil {
			return Enigma{}, err
		}
		rotors = append(rotors, rotor)
	}

	var plugboardCables PlugboardCables
	if len(values) == 4 {
		plugboardCables, err = NewPlugboardCables(values[3])
		if err != nil {
			return Enigma{}, err
		}
	}

	return Enigma{
		leftRotor:       rotors[0],
		middleRotor:     rotors[1],
		rightRotor:      rotors[2],
		plugboardCables: plugboardCables,
	}, nil
}

var regExpWithCables = regexp.MustCompile(`(?s)\[([IV,\d]*)] \[([IV,\d]*)] \[([IV,\d]*)] ?\(?([A-Z\-,]*)\)?`)
var regExpWithoutCables = regexp.MustCompile(`(?s)\[([IV,\d]*)] \[([IV,\d]*)] \[([IV,\d]*)]`)

var parseError = errors.New("error parsing enigma values, " +
	"must define 3 rotors, like [I,22,1] [III,13,24] [VI,5,12] (AZ,YN,MF,OQ,WH)")

func splitText(text string) ([]string, error) {
	matchesWithCables := regExpWithCables.FindStringSubmatch(text)
	if matchesWithCables != nil {
		groups := matchesWithCables[1:]
		return groups, nil
	}

	matchesWithoutCables := regExpWithoutCables.FindStringSubmatch(text)
	if matchesWithoutCables != nil {
		groups := matchesWithoutCables[1:]
		return groups, nil
	}

	return nil, parseError
}

func parseRotor(text string) (Rotor, error) {
	values := strings.Split(text, ",")
	if len(values) != 3 {
		return Rotor{}, parseError
	}

	number, err := romanNumberToInt(values[0])
	if err != nil {
		return Rotor{}, err
	}

	ringSetting, err := strconv.Atoi(values[1])
	if err != nil {
		return Rotor{}, err
	}

	startingPosition, err := strconv.Atoi(values[2])
	if err != nil {
		return Rotor{}, err
	}

	return NewRotor(number, ringSetting, startingPosition)
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

func (e Enigma) String() string {
	return fmt.Sprintf("%s %s %s %v",
		e.leftRotor.String(),
		e.middleRotor.String(),
		e.rightRotor.String(),
		e.plugboardCables.String(),
	)
}

func (e Enigma) rotate() {
	if e.middleRotor.IsAtNotch() {
		e.middleRotor.TurnOver()
		e.leftRotor.TurnOver()
	} else if e.rightRotor.IsAtNotch() {
		e.middleRotor.TurnOver()
	}
	e.rightRotor.TurnOver()
}
