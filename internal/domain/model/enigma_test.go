package model

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testEnigma = Enigma{
	leftRotor: Rotor{
		number:      I,
		ringSetting: RingSetting(22),
		position:    RotorPosition(1),
	},
	middleRotor: Rotor{
		number:      III,
		ringSetting: RingSetting(13),
		position:    RotorPosition(24),
	},
	rightRotor: Rotor{
		number:      VI,
		ringSetting: RingSetting(5),
		position:    RotorPosition(12),
	},
	plugboardCables: Plugboard{
		cables: []PlugboardCable{
			{
				from: 0,
				to:   25,
			},
			{
				from: 24,
				to:   13,
			},
			{
				from: 12,
				to:   5,
			},
			{
				from: 14,
				to:   16,
			},
			{
				from: 22,
				to:   7,
			},
		},
		wiring: []int{25, 1, 2, 3, 4, 12, 6, 22, 8, 9, 10, 11, 5, 24, 16, 15, 14, 17, 18, 19, 20, 21, 7, 23, 13, 0},
	},
}

func TestEnigma_String(t *testing.T) {
	tests := []struct {
		name  string
		value Enigma
		want  string
	}{
		{
			name:  "String",
			value: testEnigma,
			want:  "[I,22,1] [III,13,24] [VI,5,12] (AZ,YN,MF,OQ,WH)",
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

func TestNewEnigmaMachine(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    Enigma
		wantErr error
	}{
		{
			name:  "empty value",
			value: "",
			want:  Enigma{},
			wantErr: errors.New("error parsing enigma values, " +
				"must define 3 rotors, like [I,22,1] [III,13,24] [VI,5,12] (AZ,YN,MF,OQ,WH)"),
		},
		{
			name:  "random value",
			value: "loren ipsum",
			want:  Enigma{},
			wantErr: errors.New("error parsing enigma values, " +
				"must define 3 rotors, like [I,22,1] [III,13,24] [VI,5,12] (AZ,YN,MF,OQ,WH)"),
		},
		{
			name:  "wrong format",
			value: "[I,22:1] [III,13,24] (AZ,YN,MF,OQ,WH)",
			want:  Enigma{},
			wantErr: errors.New("error parsing enigma values, " +
				"must define 3 rotors, like [I,22,1] [III,13,24] [VI,5,12] (AZ,YN,MF,OQ,WH)"),
		},
		{
			name:  "missed one rotor",
			value: "[I,22,1] [III,13,24] (AZ,YN,MF,OQ,WH)",
			want:  Enigma{},
			wantErr: errors.New("error parsing enigma values, " +
				"must define 3 rotors, like [I,22,1] [III,13,24] [VI,5,12] (AZ,YN,MF,OQ,WH)"),
		},
		{
			name:    "happy case",
			value:   "[I,22,1] [III,13,24] [VI,5,12] (AZ,YN,MF,OQ,WH)",
			want:    testEnigma,
			wantErr: nil,
		},
		{
			name:  "no cables",
			value: "[I,22,1] [III,13,24] [VI,5,12]",
			want: Enigma{
				leftRotor: Rotor{
					number:      I,
					ringSetting: RingSetting(22),
					position:    RotorPosition(1),
				},
				middleRotor: Rotor{
					number:      III,
					ringSetting: RingSetting(13),
					position:    RotorPosition(24),
				},
				rightRotor: Rotor{
					number:      VI,
					ringSetting: RingSetting(5),
					position:    RotorPosition(12),
				},
				plugboardCables: Plugboard{
					cables: []PlugboardCable{},
					wiring: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
				},
			},
			wantErr: nil,
		},
		{
			name:    "error in roman number",
			value:   "[I,22,1] [III,13,24] [IIV,5,2]",
			want:    Enigma{},
			wantErr: errors.New("invalid roman number"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEnigmaMachine(tt.value)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
