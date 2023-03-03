package model

import (
	"fmt"
	"strings"
)

type Rotors struct {
	first  RotorNumber
	second RotorNumber
	third  RotorNumber
}

func NewRotors(value string) (Rotors, error) {
	values, err := getValues(value)
	if err != nil {
		return Rotors{}, err
	}
	first, err := NewRotorType(values[0])
	if err != nil {
		return Rotors{}, err
	}
	second, err := NewRotorType(values[1])
	if err != nil {
		return Rotors{}, err
	}
	third, err := NewRotorType(values[2])
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

func (r Rotors) String() string {
	return fmt.Sprintf("rotors: %s,%s,%s",
		r.first.String(),
		r.second.String(),
		r.third.String(),
	)
}
