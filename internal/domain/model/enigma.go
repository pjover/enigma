package model

type Rotors struct {
	first  Rotor
	second Rotor
	third  Rotor
}

type RotorPositions struct {
	first  Rotor
	second Rotor
	third  Rotor
}

type RotorRings struct {
	first  Rotor
	second Rotor
	third  Rotor
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
