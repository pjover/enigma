package model

import (
	"errors"
	"fmt"
	"strings"
)

var emptyPlugboard = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}

type PlugboardCable struct {
	from int
	to   int
}

func (p PlugboardCable) From() int {
	return p.from
}
func (p PlugboardCable) To() int {
	return p.to
}

func NewPlugboardCable(value string) (PlugboardCable, error) {
	if len(value) != 2 {
		return PlugboardCable{}, fmt.Errorf("'%s' is an invalid enigma plugboard cable value", value)
	}

	value = strings.ToUpper(value)
	from, err := validatePlugboardCableValue(int(value[0]))
	if err != nil {
		return PlugboardCable{}, err
	}

	to, err := validatePlugboardCableValue(int(value[1]))
	if err != nil {
		return PlugboardCable{}, err
	}

	if from == to {
		return PlugboardCable{}, errors.New("cannot repeat values in a enigma plugboard cable value")
	}

	return PlugboardCable{from: from - 65, to: to - 65}, nil
}

func validatePlugboardCableValue(value int) (int, error) {
	if value < 'A' || value > 'Z' {
		return 0, fmt.Errorf("'%c' is an invalid enigma plugboard value", value)
	}
	return value, nil
}

func (p PlugboardCable) String() string {
	return fmt.Sprintf("%c%c", p.from+65, p.to+65)
}

func (p PlugboardCable) Format() string {
	return fmt.Sprintf("plugboard cable %s-", p.String())
}

type Plugboard struct {
	cables []PlugboardCable
	wiring []int
}

func NewPlugboard(value string) (Plugboard, error) {
	cables := make([]PlugboardCable, 0)
	if value == "" {
		return Plugboard{
			cables,
			emptyPlugboard,
		}, nil
	}
	values := strings.Split(value, ",")
	controlMap := make(map[int]bool)
	for _, pair := range values {
		cable, err := loadCable(controlMap, pair)
		if err != nil {
			return Plugboard{}, err
		}
		cables = append(cables, cable)
	}

	return Plugboard{
		cables,
		buildWiring(cables),
	}, nil
}

func loadCable(controlMap map[int]bool, pair string) (PlugboardCable, error) {
	cable, err := NewPlugboardCable(pair)
	if err != nil {
		return PlugboardCable{}, err
	}
	if err := repeatControl(controlMap, cable.from); err != nil {
		return PlugboardCable{}, err
	}
	if err := repeatControl(controlMap, cable.to); err != nil {
		return PlugboardCable{}, err
	}
	return cable, nil
}

func repeatControl(controlMap map[int]bool, value int) error {
	_, exists := controlMap[value]
	if exists {
		return fmt.Errorf("cannot repeat the same value '%c' in different cables", value+65)
	}
	controlMap[value] = true
	return nil
}

func buildWiring(cables []PlugboardCable) []int {
	wiring := make([]int, len(emptyPlugboard))
	copy(wiring, emptyPlugboard)

	if len(cables) == 0 {
		return wiring
	}

	for _, cable := range cables {
		wiring[cable.from] = cable.to
		wiring[cable.to] = cable.from
	}
	return wiring
}

func (p Plugboard) String() string {
	s := make([]string, len(p.cables))
	for i, v := range p.cables {
		s[i] = v.String()
	}
	return "(" + strings.Join(s, ",") + ")"
}

func (p Plugboard) GetFreePlugs() []int {
	freePlugs := make(map[int]bool)
	for i := 0; i < 26; i++ {
		freePlugs[i] = true
	}

	for i := 0; i < len(p.cables); i++ {
		cable := p.cables[i]
		freePlugs[cable.from] = false
		freePlugs[cable.to] = false
	}

	result := make([]int, 0)
	for i := 0; i < 26; i++ {
		if freePlugs[i] {
			result = append(result, i)
		}
	}
	return result
}

func (p Plugboard) Forward(i int) int {
	return p.wiring[i]
}
