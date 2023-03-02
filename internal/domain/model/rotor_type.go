package model

import (
	"fmt"
	"strings"
)

type RotorType uint

const (
	I RotorType = iota
	II
	III
	IV
	V
	VI
	VII
	VIII
)

var stringValues = []string{
	"I",
	"II",
	"III",
	"IV",
	"V",
	"VI",
	"VII",
	"VIII",
}

func NewRotorType(value string) (RotorType, error) {

	value = strings.ToUpper(value)
	for i, val := range stringValues {
		if val == value {
			return RotorType(i), nil
		}
	}
	return RotorType(0), fmt.Errorf("'%s' is an invalid enigma rotor type value", value)
}

func (r RotorType) String() string {
	return stringValues[r]
}

var encodingValues = []string{
	"EKMFLGDQVZNTOWYHXUSPAIBRCJ",
	"AJDKSIRUXBLHWTMCQGZNPYFVOE",
	"BDFHJLCPRTXVZNYEIWGAKMUSQO",
	"ESOVPZJAYQUIRHXLNFTGKDCMWB",
	"VZBRGITYUPSDNHLXAWMJQOFECK",
	"JPGVOUMFYQBENHZRDKASXLICTW",
	"NZJHGRCXMYSWBOUFAIVLPEKQDT",
	"FKQHTLXOCBJSPDZRAMEWNIUYGV",
}

func (r RotorType) Encoding() string {
	return encodingValues[r]
}

func firstNotchTypeFunc(notchPosition int, rotorPosition int) bool {
	return notchPosition == rotorPosition
}

func secondNotchTypeFunc(_ int, rotorPosition int) bool {
	return rotorPosition == 12 || rotorPosition == 25
}

var isAtNotchFunctions = []func(int, int) bool{
	firstNotchTypeFunc,
	firstNotchTypeFunc,
	firstNotchTypeFunc,
	firstNotchTypeFunc,
	firstNotchTypeFunc,
	secondNotchTypeFunc,
	secondNotchTypeFunc,
	secondNotchTypeFunc,
}

func (r RotorType) IsAtNotch(notchPosition int, rotorPosition int) bool {
	return isAtNotchFunctions[r](notchPosition, rotorPosition)
}
