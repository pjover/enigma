package model

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRotorPosition1(t *testing.T) {
	tests := []struct {
		name     string
		position int
		want     RotorPosition
		wantErr  error
	}{
		{
			name:     "Zero",
			position: 0,
			want:     0,
			wantErr:  errors.New("0 is an invalid enigma rotor position value"),
		},
		{
			name:     "One",
			position: 1,
			want:     1,
			wantErr:  nil,
		},
		{
			name:     "Twenty six",
			position: 26,
			want:     26,
			wantErr:  nil,
		},
		{
			name:     "Twenty seven",
			position: 27,
			want:     0,
			wantErr:  errors.New("27 is an invalid enigma rotor position value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRotorPosition(tt.position)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
