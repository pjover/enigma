package model

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testEnigma = Enigma{
	rotors: Rotors{
		first: Rotor{
			number:      I,
			ringSetting: RingSetting(22),
			position:    RotorPosition(1),
		},
		second: Rotor{
			number:      III,
			ringSetting: RingSetting(13),
			position:    RotorPosition(24),
		},
		third: Rotor{
			number:      VI,
			ringSetting: RingSetting(5),
			position:    RotorPosition(12),
		},
	},
	plugboardCables: PlugboardCables{
		PlugboardCable{
			from: 'A',
			to:   'Z',
		},
		PlugboardCable{
			from: 'Y',
			to:   'N',
		},
		PlugboardCable{
			from: 'M',
			to:   'F',
		},
		PlugboardCable{
			from: 'O',
			to:   'Q',
		},
		PlugboardCable{
			from: 'W',
			to:   'H',
		},
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
			want:  "[I,22:1] [III,13:24] [VI,5:12] (A-Z,Y-N,M-F,O-Q,W-H)",
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
			name:    "empty value",
			value:   "",
			want:    Enigma{},
			wantErr: errors.New("error parsing enigma values, must define 4 groups separated by '#'"),
		},
		{
			name:    "missed one value",
			value:   "1,3,6#22,13,5#1,24,12",
			want:    Enigma{},
			wantErr: errors.New("error parsing enigma values, must define 4 groups separated by '#'"),
		},
		{
			name:    "happy case",
			value:   "1,3,6#22,13,5#1,24,12#AZ,YN,MF,OQ,WH",
			want:    testEnigma,
			wantErr: nil,
		},
		{
			name:  "no cables",
			value: "1,3,6#22,13,5#1,24,12#",
			want: Enigma{
				rotors: Rotors{
					first: Rotor{
						number:      I,
						ringSetting: RingSetting(22),
						position:    RotorPosition(1),
					},
					second: Rotor{
						number:      III,
						ringSetting: RingSetting(13),
						position:    RotorPosition(24),
					},
					third: Rotor{
						number:      VI,
						ringSetting: RingSetting(5),
						position:    RotorPosition(12),
					},
				},
				plugboardCables: PlugboardCables{},
			},
			wantErr: nil,
		},
		{
			name:    "error in value",
			value:   "1,3,10#1,24,12#22,13,5#AZ,YN,MF,OQ,WH",
			want:    Enigma{},
			wantErr: errors.New("10 is an invalid enigma rotor number"),
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
