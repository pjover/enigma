package model

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRotorRing(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    RotorRing
		wantErr error
	}{
		{
			name:    "Zero",
			value:   "0",
			want:    0,
			wantErr: errors.New("'0' is an invalid enigma rotor ring value"),
		},
		{
			name:    "One",
			value:   "1",
			want:    1,
			wantErr: nil,
		},
		{
			name:    "Twenty six",
			value:   "26",
			want:    26,
			wantErr: nil,
		},
		{
			name:    "Twenty seven",
			value:   "27",
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

func TestNewRotorRings(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    RotorRings
		wantErr error
	}{
		{
			name:  "OK",
			value: "1,15,26",
			want: RotorRings{
				first:  1,
				second: 15,
				third:  26,
			},
			wantErr: nil,
		},
		{
			name:    "Error 2 values",
			value:   "1,15",
			want:    RotorRings{},
			wantErr: errors.New("please define 3 values instead of 2 at '1,15'"),
		},
		{
			name:    "Error 1 value",
			value:   "",
			want:    RotorRings{},
			wantErr: errors.New("please define 3 values instead of 1 at ''"),
		},
		{
			name:    "Error wrong value",
			value:   "1,15,27",
			want:    RotorRings{},
			wantErr: errors.New("'27' is an invalid enigma rotor ring value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRotorRings(tt.value)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestRotorRings_String(t *testing.T) {
	tests := []struct {
		name  string
		value RotorRings
		want  string
	}{
		{
			name: "String",
			value: RotorRings{
				first:  1,
				second: 2,
				third:  26,
			},
			want: "rings: 1,2,26",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := tt.value
			got := sut.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
