package model

import (
	"fmt"
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

type RotorRings struct {
	first  RotorRing
	second RotorRing
	third  RotorRing
}

func NewRotorRings(value string) (RotorRings, error) {
	values, err := getValues(value)
	if err != nil {
		return RotorRings{}, err
	}
	first, err := NewRotorRing(values[0])
	if err != nil {
		return RotorRings{}, err
	}
	second, err := NewRotorRing(values[1])
	if err != nil {
		return RotorRings{}, err
	}
	third, err := NewRotorRing(values[2])
	if err != nil {
		return RotorRings{}, err
	}
	return RotorRings{
		first:  first,
		second: second,
		third:  third,
	}, nil
}

func (r RotorRings) String() string {
	return fmt.Sprintf("rings: %s,%s,%s",
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

func NewEnigmaMachine(rotors Rotors, rotorPositions RotorPositions, rotorRings RotorRings, plugboardCables PlugboardCables) Enigma {
	return Enigma{
		rotors:          rotors,
		rotorPositions:  rotorPositions,
		rotorRings:      rotorRings,
		plugboardCables: plugboardCables,
	}
}

func (e Enigma) Format() string {
	return fmt.Sprintf("Enigma %s, %s, %s, %v",
		e.rotors.String(),
		e.rotorPositions.String(),
		e.rotorRings.String(),
		e.plugboardCables,
	)
}
