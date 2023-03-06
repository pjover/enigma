package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotor_Forward(t *testing.T) {
	sut := &Rotor{
		number:      I,
		ringSetting: RingSetting(3),
		position:    RotorPosition(2),
	}
	tests := []struct {
		name string
		i    int
		want int
	}{
		{"Forward 0", 0, 10},
		{"Forward 1", 1, 5},
		{"Forward 2", 2, 11},
		{"Forward 3", 3, 13},
		{"Forward 4", 4, 6},
		{"Forward 5", 5, 12},
		{"Forward 6", 6, 7},
		{"Forward 7", 7, 4},
		{"Forward 8", 8, 17},
		{"Forward 9", 9, 22},
		{"Forward 10", 10, 0},
		{"Forward 11", 11, 14},
		{"Forward 12", 12, 20},
		{"Forward 13", 13, 15},
		{"Forward 14", 14, 23},
		{"Forward 15", 15, 25},
		{"Forward 16", 16, 8},
		{"Forward 17", 17, 24},
		{"Forward 18", 18, 21},
		{"Forward 19", 19, 19},
		{"Forward 20", 20, 16},
		{"Forward 21", 21, 1},
		{"Forward 22", 22, 9},
		{"Forward 23", 23, 2},
		{"Forward 24", 24, 18},
		{"Forward 25", 25, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := sut.Forward(tt.i)
			assert.Equal(t, tt.want, actual)
		})
	}
}

func TestRotor_Backward(t *testing.T) {
	sut := &Rotor{
		number:      I,
		ringSetting: RingSetting(3),
		position:    RotorPosition(2),
	}
	tests := []struct {
		name string
		i    int
		want int
	}{
		{"Backward 0", 0, 10},
		{"Backward 1", 1, 21},
		{"Backward 2", 2, 23},
		{"Backward 3", 3, 25},
		{"Backward 4", 4, 7},
		{"Backward 5", 5, 1},
		{"Backward 6", 6, 4},
		{"Backward 7", 7, 6},
		{"Backward 8", 8, 16},
		{"Backward 9", 9, 22},
		{"Backward 10", 10, 0},
		{"Backward 11", 11, 2},
		{"Backward 12", 12, 5},
		{"Backward 13", 13, 3},
		{"Backward 14", 14, 11},
		{"Backward 15", 15, 13},
		{"Backward 16", 16, 20},
		{"Backward 17", 17, 8},
		{"Backward 18", 18, 24},
		{"Backward 19", 19, 19},
		{"Backward 20", 20, 12},
		{"Backward 21", 21, 18},
		{"Backward 22", 22, 9},
		{"Backward 23", 23, 14},
		{"Backward 24", 24, 17},
		{"Backward 25", 25, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := sut.Backward(tt.i)
			assert.Equal(t, tt.want, actual)
		})
	}
}

func TestRotor_TurnOver(t *testing.T) {
	type fields struct {
		number        RotorNumber
		ringSetting   RingSetting
		rotorPosition RotorPosition
	}
	tests := []struct {
		name   string
		fields fields
		want   RotorPosition
	}{
		{
			name: "Case 0",
			fields: fields{
				I,
				RingSetting(1),
				0,
			},
			want: RotorPosition(1),
		},
		{
			name: "Case 1",
			fields: fields{
				I,
				RingSetting(1),
				1,
			},
			want: RotorPosition(2),
		},
		{
			name: "Case 24",
			fields: fields{
				I,
				RingSetting(2),
				24,
			},
			want: RotorPosition(25),
		},
		{
			name: "Case 25",
			fields: fields{
				I,
				RingSetting(3),
				25,
			},
			want: RotorPosition(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rotor{
				number:      tt.fields.number,
				ringSetting: tt.fields.ringSetting,
				position:    tt.fields.rotorPosition,
			}
			r.TurnOver()
			assert.Equal(t, tt.want, r.position)
		})
	}
}

func TestRotor_IsAtNotch(t *testing.T) {
	type fields struct {
		number        RotorNumber
		ringSetting   RingSetting
		rotorPosition RotorPosition
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
				RingSetting(2),
				16,
			},
			want: true,
		},
		{
			name: "Is not a notch",
			fields: fields{
				I,
				RingSetting(2),
				17,
			},
			want: false,
		},
		{
			name: "Is a notch",
			fields: fields{
				VI,
				RingSetting(2),
				0,
			},
			want: true,
		},
		{
			name: "Is not a notch",
			fields: fields{
				VI,
				RingSetting(2),
				17,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rotor{
				number:      tt.fields.number,
				ringSetting: tt.fields.ringSetting,
				position:    tt.fields.rotorPosition,
			}
			assert.Equalf(t, tt.want, r.IsAtNotch(), "IsAtNotch()")
		})
	}
}
