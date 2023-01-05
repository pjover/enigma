package model

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPlugboardCable(t *testing.T) {
	tests := []struct {
		name   string
		value  string
		actual PlugboardCable
		err    error
	}{
		{
			name:   "AB",
			value:  "AB",
			actual: PlugboardCable{from: 'A', to: 'B'},
			err:    nil,
		},
		{
			name:   "YZ",
			value:  "YZ",
			actual: PlugboardCable{from: 'Y', to: 'Z'},
			err:    nil,
		},
		{
			name:   "ZY",
			value:  "ZY",
			actual: PlugboardCable{from: 'Z', to: 'Y'},
			err:    nil,
		},
		{
			name:   "Empty input",
			value:  "",
			actual: PlugboardCable{from: 0, to: 0},
			err:    errors.New("'' is an invalid enigma plugboard cable value"),
		},
		{
			name:   "Too short",
			value:  "A",
			actual: PlugboardCable{from: 0, to: 0},
			err:    errors.New("'A' is an invalid enigma plugboard cable value"),
		},
		{
			name:   "Too long",
			value:  "ABC",
			actual: PlugboardCable{from: 0, to: 0},
			err:    errors.New("'ABC' is an invalid enigma plugboard cable value"),
		},
		{
			name:   "Lower boundary",
			value:  "@Y",
			actual: PlugboardCable{from: 0, to: 0},
			err:    errors.New("'@' is an invalid enigma plugboard value"),
		},
		{
			name:   "Upper boundary",
			value:  "A[",
			actual: PlugboardCable{from: 0, to: 0},
			err:    errors.New("'[' is an invalid enigma plugboard value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPlugboardCable(tt.value)
			assert.Equal(t, tt.actual, got)
			assert.Equal(t, tt.err, err)
		})
	}
}
