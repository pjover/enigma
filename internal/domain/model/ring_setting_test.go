package model

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRingSetting(t *testing.T) {
	tests := []struct {
		name        string
		ringSetting int
		want        RingSetting
		wantErr     error
	}{
		{
			name:        "Negative",
			ringSetting: -1,
			want:        0,
			wantErr:     errors.New("-1 is an invalid enigma rotor ring setting value"),
		},
		{
			name:        "Zero",
			ringSetting: 0,
			want:        0,
			wantErr:     nil,
		},
		{
			name:        "One",
			ringSetting: 1,
			want:        1,
			wantErr:     nil,
		},
		{
			name:        "Twenty five",
			ringSetting: 25,
			want:        25,
			wantErr:     nil,
		},
		{
			name:        "Twenty six",
			ringSetting: 26,
			want:        0,
			wantErr:     errors.New("26 is an invalid enigma rotor ring setting value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRingSetting(tt.ringSetting)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
