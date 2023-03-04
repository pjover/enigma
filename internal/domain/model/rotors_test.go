package model

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRotors(t *testing.T) {
	type args struct {
		numbers           [3]uint
		ringSettings      [3]uint
		startingPositions [3]uint
	}
	tests := []struct {
		name    string
		args    args
		want    Rotors
		wantErr error
	}{
		{
			name: "OK",
			args: args{
				numbers:           [3]uint{1, 8, 2},
				ringSettings:      [3]uint{3, 13, 23},
				startingPositions: [3]uint{2, 12, 22},
			},
			want: Rotors{
				first: Rotor{
					number:      I,
					ringSetting: RingSetting(3),
					position:    RotorPosition(2),
				},
				second: Rotor{
					number:      VIII,
					ringSetting: RingSetting(13),
					position:    RotorPosition(12),
				},
				third: Rotor{
					number:      II,
					ringSetting: RingSetting(23),
					position:    RotorPosition(22),
				},
			},
			wantErr: nil,
		},
		{
			name: "Error wrong value",
			args: args{
				numbers:           [3]uint{1, 8, 2},
				ringSettings:      [3]uint{3, 13, 23},
				startingPositions: [3]uint{2, 12, 36},
			},
			want:    Rotors{},
			wantErr: errors.New("36 is an invalid enigma rotor position value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRotors(tt.args.numbers, tt.args.ringSettings, tt.args.startingPositions)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestRotors_String(t *testing.T) {
	type fields struct {
		first  Rotor
		second Rotor
		third  Rotor
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "OK",
			fields: fields{
				first: Rotor{
					number:      I,
					ringSetting: RingSetting(3),
					position:    RotorPosition(2),
				},
				second: Rotor{
					number:      VIII,
					ringSetting: RingSetting(13),
					position:    RotorPosition(12),
				},
				third: Rotor{
					number:      II,
					ringSetting: RingSetting(23),
					position:    RotorPosition(22),
				},
			},
			want: "[I,3:2] [VIII,13:12] [II,23:22]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rotors{
				first:  tt.fields.first,
				second: tt.fields.second,
				third:  tt.fields.third,
			}
			assert.Equalf(t, tt.want, r.String(), "String()")
		})
	}
}
