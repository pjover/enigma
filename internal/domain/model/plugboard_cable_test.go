package model

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPlugboardCable(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    PlugboardCable
		wantErr error
	}{
		{
			name:    "AB",
			value:   "AB",
			want:    PlugboardCable{from: 'A', to: 'B'},
			wantErr: nil,
		},
		{
			name:    "YZ",
			value:   "YZ",
			want:    PlugboardCable{from: 'Y', to: 'Z'},
			wantErr: nil,
		},
		{
			name:    "ZY",
			value:   "ZY",
			want:    PlugboardCable{from: 'Z', to: 'Y'},
			wantErr: nil,
		},
		{
			name:    "Empty input",
			value:   "",
			want:    PlugboardCable{from: 0, to: 0},
			wantErr: errors.New("'' is an invalid enigma plugboard cable value"),
		},
		{
			name:    "Too short",
			value:   "A",
			want:    PlugboardCable{from: 0, to: 0},
			wantErr: errors.New("'A' is an invalid enigma plugboard cable value"),
		},
		{
			name:    "Too long",
			value:   "ABC",
			want:    PlugboardCable{from: 0, to: 0},
			wantErr: errors.New("'ABC' is an invalid enigma plugboard cable value"),
		},
		{
			name:    "Lower boundary",
			value:   "@Y",
			want:    PlugboardCable{from: 0, to: 0},
			wantErr: errors.New("'@' is an invalid enigma plugboard value"),
		},
		{
			name:    "Upper boundary",
			value:   "A[",
			want:    PlugboardCable{from: 0, to: 0},
			wantErr: errors.New("'[' is an invalid enigma plugboard value"),
		},
		{
			name:    "Repeated value",
			value:   "AA",
			want:    PlugboardCable{from: 0, to: 0},
			wantErr: errors.New("cannot repeat values in a enigma plugboard cable value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPlugboardCable(tt.value)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestNewPlugboardCables(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    PlugboardCables
		wantErr error
	}{
		{
			name:  "OK",
			value: "AB,ZY,MN",
			want: PlugboardCables{
				{from: 'A', to: 'B'},
				{from: 'Z', to: 'Y'},
				{from: 'M', to: 'N'},
			},
			wantErr: nil,
		},
		{
			name:    "Empty plugboard",
			value:   "",
			want:    PlugboardCables{},
			wantErr: nil,
		},
		{
			name:    "Wrong cable",
			value:   "AB,ZY,MM",
			want:    PlugboardCables{},
			wantErr: errors.New("cannot repeat values in a enigma plugboard cable value"),
		},
		{
			name:    "Repeat the same value in different cables",
			value:   "AB,ZY,AM",
			want:    PlugboardCables{},
			wantErr: errors.New("cannot repeat the same value 'A' in different cables"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPlugboardCables(tt.value)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
