package model

type Rotor struct {
	number        RotorNumber
	ringSetting   RingSetting
	rotorPosition uint
}

func NewRotor(number RotorNumber, startingPosition StartingPosition, ringSetting RingSetting) Rotor {
	return Rotor{
		number:        number,
		ringSetting:   ringSetting,
		rotorPosition: uint(startingPosition),
	}
}

func (r *Rotor) Number() RotorNumber {
	return r.number
}

func (r *Rotor) RingSetting() RingSetting {
	return r.ringSetting
}

func (r *Rotor) RotorPosition() uint {
	return r.rotorPosition
}

func (r *Rotor) encipher(i uint, rotorPosition uint, ring RingSetting, wiring []uint) uint {
	shift := rotorPosition - uint(ring)
	return (wiring[(i+shift+26)%26] - shift + 26) % 26
}

func (r *Rotor) Forward(i uint) uint {
	return r.encipher(i, r.rotorPosition, r.ringSetting, r.number.ForwardWiring())
}

func (r *Rotor) Backward(i uint) uint {
	return r.encipher(i, r.rotorPosition, r.ringSetting, r.number.InverseWiring())
}

func (r *Rotor) IsAtNotch() bool {
	return r.number.NotchPosition() == r.rotorPosition
}

func (r *Rotor) TurnOver() {
	r.rotorPosition = (r.rotorPosition + 1) % 26
}
