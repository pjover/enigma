package model

import (
	"errors"
	"fmt"
	"strings"
)

type RotorPositions struct {
	first  RotorPosition
	second RotorPosition
	third  RotorPosition
}

func NewRotorPositions(value string) (RotorPositions, error) {
	values, err := getValues(value)
	if err != nil {
		return RotorPositions{}, err
	}
	first, err := NewRotorPosition(values[0])
	if err != nil {
		return RotorPositions{}, err
	}
	second, err := NewRotorPosition(values[1])
	if err != nil {
		return RotorPositions{}, err
	}
	third, err := NewRotorPosition(values[2])
	if err != nil {
		return RotorPositions{}, err
	}
	return RotorPositions{
		first:  first,
		second: second,
		third:  third,
	}, nil
}

func (r RotorPositions) String() string {
	return fmt.Sprintf("pos: %s,%s,%s",
		r.first.String(),
		r.second.String(),
		r.third.String(),
	)
}

type Enigma struct {
	rotors          Rotors
	rotorPositions  RotorPositions
	rotorRings      RotorRings
	plugboardCables PlugboardCables
}

func (e Enigma) Rotors() Rotors {
	return e.rotors
}

func (e Enigma) RotorPositions() RotorPositions {
	return e.rotorPositions
}

func (e Enigma) RotorRings() RotorRings {
	return e.rotorRings
}

func (e Enigma) PlugboardCables() PlugboardCables {
	return e.plugboardCables
}

func NewEnigmaMachine(value string) (Enigma, error) {
	values := strings.Split(value, "#")
	if len(values) != 4 {
		return Enigma{}, errors.New("error parsing enigma values, must define 4 groups separated by '#'")
	}
	rotors, err := NewRotors(values[0])
	if err != nil {
		return Enigma{}, err
	}
	rotorPositions, err := NewRotorPositions(values[1])
	if err != nil {
		return Enigma{}, err
	}
	rotorRings, err := NewRotorRings(values[2])
	if err != nil {
		return Enigma{}, err
	}
	plugboardCables, err := NewPlugboardCables(values[3])
	if err != nil {
		return Enigma{}, err
	}

	return Enigma{
		rotors:          rotors,
		rotorPositions:  rotorPositions,
		rotorRings:      rotorRings,
		plugboardCables: plugboardCables,
	}, nil
}

func (e Enigma) String() string {
	return fmt.Sprintf("Enigma %s, %s, %s, %v",
		e.rotors.String(),
		e.rotorPositions.String(),
		e.rotorRings.String(),
		e.plugboardCables.String(),
	)
}
