package model

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testEnigma = Enigma{
	rotors: Rotors{
		first:  I,
		second: III,
		third:  VI,
	},
	startingPositions: StartingPositions{
		first:  1,
		second: 24,
		third:  12,
	},
	ringSettings: RingSettings{
		first:  22,
		second: 13,
		third:  5,
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
			want:  "Enigma rotors: I,III,VI, pos: 1,24,12, rings: 22,13,5, cables: A-Z,Y-N,M-F,O-Q,W-H",
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
			value:   "I,III,VI#1,24,12#22,13,5",
			want:    Enigma{},
			wantErr: errors.New("error parsing enigma values, must define 4 groups separated by '#'"),
		},
		{
			name:    "happy case",
			value:   "I,III,VI#1,24,12#22,13,5#AZ,YN,MF,OQ,WH",
			want:    testEnigma,
			wantErr: nil,
		},
		{
			name:  "no cables",
			value: "I,III,VI#1,24,12#22,13,5#",
			want: Enigma{
				rotors: Rotors{
					first:  I,
					second: III,
					third:  VI,
				},
				startingPositions: StartingPositions{
					first:  1,
					second: 24,
					third:  12,
				},
				ringSettings: RingSettings{
					first:  22,
					second: 13,
					third:  5,
				},
				plugboardCables: PlugboardCables{},
			},
			wantErr: nil,
		},
		{
			name:    "error in value",
			value:   "I,III,X#1,24,12#22,13,5#AZ,YN,MF,OQ,WH",
			want:    Enigma{},
			wantErr: errors.New("'X' is an invalid enigma rotor number"),
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
