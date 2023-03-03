package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotor_TurnOver(t *testing.T) {
	type fields struct {
		number        RotorNumber
		ringSetting   RingSetting
		rotorPosition uint
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		{
			name: "Case 0",
			fields: fields{
				I,
				RingSetting(uint(1)),
				0,
			},
			want: uint(1),
		},
		{
			name: "Case 1",
			fields: fields{
				I,
				RingSetting(uint(1)),
				1,
			},
			want: uint(2),
		},
		{
			name: "Case 24",
			fields: fields{
				I,
				RingSetting(uint(2)),
				24,
			},
			want: uint(25),
		},
		{
			name: "Case 25",
			fields: fields{
				I,
				RingSetting(uint(3)),
				25,
			},
			want: uint(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rotor{
				number:        tt.fields.number,
				ringSetting:   tt.fields.ringSetting,
				rotorPosition: tt.fields.rotorPosition,
			}
			r.TurnOver()
			assert.Equal(t, tt.want, r.rotorPosition)
		})
	}
}

func TestRotor_IsAtNotch(t *testing.T) {
	type fields struct {
		number        RotorNumber
		ringSetting   RingSetting
		rotorPosition uint
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Is a notch",
			fields: fields{
				I,
				RingSetting(uint(2)),
				16,
			},
			want: true,
		},
		{
			name: "Is not a notch",
			fields: fields{
				I,
				RingSetting(uint(2)),
				17,
			},
			want: false,
		},
		{
			name: "Is a notch",
			fields: fields{
				VI,
				RingSetting(uint(2)),
				0,
			},
			want: true,
		},
		{
			name: "Is not a notch",
			fields: fields{
				VI,
				RingSetting(uint(2)),
				17,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rotor{
				number:        tt.fields.number,
				ringSetting:   tt.fields.ringSetting,
				rotorPosition: tt.fields.rotorPosition,
			}
			assert.Equalf(t, tt.want, r.IsAtNotch(), "IsAtNotch()")
		})
	}
}
