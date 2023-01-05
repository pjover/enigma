package model

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRotor(t *testing.T) {
	tests := []struct {
		name   string
		value  string
		actual Rotor
		err    error
	}{
		{
			name:   "I",
			value:  "i",
			actual: I,
			err:    nil,
		},
		{
			name:   "II",
			value:  "II",
			actual: II,
			err:    nil,
		},
		{
			name:   "III",
			value:  "iii",
			actual: III,
			err:    nil,
		},
		{
			name:   "IV",
			value:  "IV",
			actual: IV,
			err:    nil,
		},
		{
			name:   "V",
			value:  "v",
			actual: V,
			err:    nil,
		},
		{
			name:   "VI",
			value:  "VI",
			actual: VI,
			err:    nil,
		},
		{
			name:   "VII",
			value:  "vii",
			actual: VII,
			err:    nil,
		},
		{
			name:   "VIII",
			value:  "VIII",
			actual: VIII,
			err:    nil,
		},
		{
			name:   "Invalid rotor value",
			value:  "iX",
			actual: I,
			err:    errors.New("'IX' is an invalid enigma rotor value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRotor(tt.value)
			assert.Equal(t, tt.actual, got)
			assert.Equal(t, tt.err, err)
		})
	}
}
