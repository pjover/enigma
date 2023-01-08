package model

import (
	"fmt"
	"strconv"
)

type RotorRing uint

func NewRotorRing(value string) (RotorRing, error) {
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("'%s' is an invalid enigma rotor ring number", value)
	}
	if number == 0 || number > 26 {
		return RotorRing(0), fmt.Errorf("'%s' is an invalid enigma rotor ring value", value)
	}
	return RotorRing(number), nil
}

func (r RotorRing) String() string {
	return strconv.Itoa(int(r))
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
