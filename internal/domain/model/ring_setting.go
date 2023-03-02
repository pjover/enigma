package model

import (
	"fmt"
	"strconv"
)

type RingSetting uint

func NewRingSetting(value string) (RingSetting, error) {
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("'%s' is an invalid enigma rotor ring setting number", value)
	}
	if number == 0 || number > 26 {
		return RingSetting(0), fmt.Errorf("'%s' is an invalid enigma rotor ring setting value", value)
	}
	return RingSetting(number), nil
}

func (r RingSetting) String() string {
	return strconv.Itoa(int(r))
}

type RingSettings struct {
	first  RingSetting
	second RingSetting
	third  RingSetting
}

func NewRingSettings(value string) (RingSettings, error) {
	values, err := getValues(value)
	if err != nil {
		return RingSettings{}, err
	}
	first, err := NewRingSetting(values[0])
	if err != nil {
		return RingSettings{}, err
	}
	second, err := NewRingSetting(values[1])
	if err != nil {
		return RingSettings{}, err
	}
	third, err := NewRingSetting(values[2])
	if err != nil {
		return RingSettings{}, err
	}
	return RingSettings{
		first:  first,
		second: second,
		third:  third,
	}, nil
}

func (r RingSettings) String() string {
	return fmt.Sprintf("rings: %s,%s,%s",
		r.first.String(),
		r.second.String(),
		r.third.String(),
	)
}
