package model

import (
	"errors"
	"fmt"
	"strings"
)

type StartingPositions struct {
	first  StartingPosition
	second StartingPosition
	third  StartingPosition
}

func NewStartingPositions(value string) (StartingPositions, error) {
	values, err := getValues(value)
	if err != nil {
		return StartingPositions{}, err
	}
	first, err := NewStartingPosition(values[0])
	if err != nil {
		return StartingPositions{}, err
	}
	second, err := NewStartingPosition(values[1])
	if err != nil {
		return StartingPositions{}, err
	}
	third, err := NewStartingPosition(values[2])
	if err != nil {
		return StartingPositions{}, err
	}
	return StartingPositions{
		first:  first,
		second: second,
		third:  third,
	}, nil
}

func (r StartingPositions) String() string {
	return fmt.Sprintf("pos: %s,%s,%s",
		r.first.String(),
		r.second.String(),
		r.third.String(),
	)
}

type Enigma struct {
	rotors            Rotors
	startingPositions StartingPositions
	ringSettings      RingSettings
	plugboardCables   PlugboardCables
}

func (e Enigma) Rotors() Rotors {
	return e.rotors
}

func (e Enigma) StartingPositions() StartingPositions {
	return e.startingPositions
}

func (e Enigma) RingSettings() RingSettings {
	return e.ringSettings
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
	rotorPositions, err := NewStartingPositions(values[1])
	if err != nil {
		return Enigma{}, err
	}
	rotorRings, err := NewRingSettings(values[2])
	if err != nil {
		return Enigma{}, err
	}
	plugboardCables, err := NewPlugboardCables(values[3])
	if err != nil {
		return Enigma{}, err
	}

	return Enigma{
		rotors:            rotors,
		startingPositions: rotorPositions,
		ringSettings:      rotorRings,
		plugboardCables:   plugboardCables,
	}, nil
}

func (e Enigma) String() string {
	return fmt.Sprintf("Enigma %s, %s, %s, %v",
		e.rotors.String(),
		e.startingPositions.String(),
		e.ringSettings.String(),
		e.plugboardCables.String(),
	)
}
