package model

import (
	"fmt"
)

type Rotors struct {
	first  Rotor
	second Rotor
	third  Rotor
}

func NewRotors(rotorNumbers [3]uint, ringSettings [3]uint, startingPositions [3]uint) (Rotors, error) {
	first, err := NewRotor(rotorNumbers[0], ringSettings[0], startingPositions[0])
	if err != nil {
		return Rotors{}, err
	}
	second, err := NewRotor(rotorNumbers[1], ringSettings[1], startingPositions[1])
	if err != nil {
		return Rotors{}, err
	}
	third, err := NewRotor(rotorNumbers[2], ringSettings[2], startingPositions[2])
	if err != nil {
		return Rotors{}, err
	}

	return Rotors{
		first:  first,
		second: second,
		third:  third,
	}, nil
}

func (r Rotors) String() string {
	return fmt.Sprintf("%s %s %s",
		r.first.String(),
		r.second.String(),
		r.third.String(),
	)
}
