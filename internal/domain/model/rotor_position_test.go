package model

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRotorPosition(t *testing.T) {
	tests := []struct {
		name   string
		value  uint
		actual RotorPosition
		err    error
	}{
		{
			name:   "Zero",
			value:  0,
			actual: 0,
			err:    errors.New("'0' is an invalid enigma rotor position value"),
		},
		{
			name:   "One",
			value:  1,
			actual: 1,
			err:    nil,
		},
		{
			name:   "Twenty six",
			value:  26,
			actual: 26,
			err:    nil,
		},
		{
			name:   "Twenty seven",
			value:  27,
			actual: 0,
			err:    errors.New("'27' is an invalid enigma rotor position value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRotorPosition(tt.value)
			assert.Equal(t, tt.actual, got)
			assert.Equal(t, tt.err, err)
		})
	}
}
