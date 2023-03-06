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
			name:     "Negative",
			position: -1,
			want:     0,
			wantErr:  errors.New("-1 is an invalid enigma rotor position value"),
		},
		{
			name:     "Zero",
			position: 0,
			want:     0,
			wantErr:  nil,
		},
		{
			name:     "One",
			position: 1,
			want:     1,
			wantErr:  nil,
		},
		{
			name:     "Twenty five",
			position: 25,
			want:     25,
			wantErr:  nil,
		},
		{
			name:     "Twenty six",
			position: 26,
			want:     0,
			wantErr:  errors.New("26 is an invalid enigma rotor position value"),
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
