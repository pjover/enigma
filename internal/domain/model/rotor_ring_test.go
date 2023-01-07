package model

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRotorRing(t *testing.T) {
	tests := []struct {
		name    string
		value   uint
		want    RotorRing
		wantErr error
	}{
		{
			name:    "Zero",
			value:   0,
			want:    0,
			wantErr: errors.New("'0' is an invalid enigma rotor ring value"),
		},
		{
			name:    "One",
			value:   1,
			want:    1,
			wantErr: nil,
		},
		{
			name:    "Twenty six",
			value:   26,
			want:    26,
			wantErr: nil,
		},
		{
			name:    "Twenty seven",
			value:   27,
			want:    0,
			wantErr: errors.New("'27' is an invalid enigma rotor ring value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRotorRing(tt.value)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
