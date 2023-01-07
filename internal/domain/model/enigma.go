package model

import (
	"fmt"
	"strings"
)

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

func (r Rotors) Format() string {
	return fmt.Sprintf("rotors: %s,%s,%s",
		r.first.String(),
		r.second.String(),
		r.third.String(),
	)
}

type RotorPositions struct {
	first  Rotor
	second Rotor
	third  Rotor
}

func (r RotorPositions) Format() string {
	return fmt.Sprintf("pos: %s, %s, %s",
		r.first.String(),
		r.second.String(),
		r.third.String(),
	)
}

type RotorRings struct {
	first  Rotor
	second Rotor
	third  Rotor
}

func (r RotorRings) Format() string {
	return fmt.Sprintf("rings: %s, %s, %s",
		r.first.String(),
		r.second.String(),
		r.third.String(),
	)
}

type PlugboardCables []PlugboardCable

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
		e.rotors.Format(),
		e.rotorPositions.Format(),
		e.rotorRings.Format(),
		e.plugboardCables,
	)
}
