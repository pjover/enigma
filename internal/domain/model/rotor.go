package model

import "fmt"

type Rotor struct {
	number      RotorNumber
	ringSetting RingSetting
	position    RotorPosition
}

func NewRotor(number int, ringSetting int, startingPosition int) (Rotor, error) {
	rn, err := NewRotorNumber(number)
	if err != nil {
		return Rotor{}, err
	}
	rs, err := NewRingSetting(ringSetting)
	if err != nil {
		return Rotor{}, err
	}
	rp, err := NewRotorPosition(startingPosition)
	if err != nil {
		return Rotor{}, err
	}
	return Rotor{
		number:      rn,
		ringSetting: rs,
		position:    rp,
	}, nil
}

func (r *Rotor) Number() RotorNumber {
	return r.number
}

func (r *Rotor) RingSetting() RingSetting {
	return r.ringSetting
}

func (r *Rotor) Position() RotorPosition {
	return r.position
}

func (r *Rotor) encipher(i int, position RotorPosition, ring RingSetting, wiring []int) int {
	shift := int(position) - int(ring)
	return (wiring[(i+shift+26)%26] - shift + 26) % 26
}

func (r *Rotor) Forward(i int) int {
	return r.encipher(i, r.position, r.ringSetting, r.number.ForwardWiring())
}

func (r *Rotor) Backward(i int) int {
	return r.encipher(i, r.position, r.ringSetting, r.number.InverseWiring())
}

func (r *Rotor) IsAtNotch() bool {
	switch r.number {
	case I:
		return r.position == 16
	case II:
		return r.position == 4
	case III:
		return r.position == 21
	case IV:
		return r.position == 9
	case V:
		return r.position == 25
	default:
		return r.position == 12 || r.position == 25
	}
}

func (r *Rotor) TurnOver() {
	r.position = (r.position + 1) % 26
}

func (r *Rotor) String() string {
	return fmt.Sprintf("[%s,%s,%d]",
		r.number.String(),
		r.ringSetting.String(),
		r.position,
	)
}
