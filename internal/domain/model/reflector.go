package model

import (
	"fmt"
)

type Reflector uint

const (
	A Reflector = iota
	B
	C
)

func NewReflector(code string) (Reflector, error) {
	switch code {
	case "A":
		return A, nil
	case "B":
		return B, nil
	case "C":
		return C, nil
	default:
		return 0, fmt.Errorf("invalid reflector %s", code)
	}
}

var reflectorEncodingValues = []string{
	"ZYXWVUTSRQPONMLKJIHGFEDCBA",
	"YRUHQSLDPXNGOKMIEBFZCWVJAT",
	"FVPJIAOYEDRZXWGCTKUQSBNMHL",
}

func (r Reflector) encoding() string {
	return reflectorEncodingValues[r]
}

var reflectorForwardWiringValues = [][]int{
	{25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
	{24, 17, 20, 7, 16, 18, 11, 3, 15, 23, 13, 6, 14, 10, 12, 8, 4, 1, 5, 25, 2, 22, 21, 9, 0, 19},
	{5, 21, 15, 9, 8, 0, 14, 24, 4, 3, 17, 25, 23, 22, 6, 2, 19, 10, 20, 16, 18, 1, 13, 12, 7, 11},
}

func (r Reflector) ForwardWiring() []int {
	return reflectorForwardWiringValues[r]
}

var reflectorInverseWiringValues = [][]int{
	{25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
	{24, 17, 20, 7, 16, 18, 11, 3, 15, 23, 13, 6, 14, 10, 12, 8, 4, 1, 5, 25, 2, 22, 21, 9, 0, 19},
	{5, 21, 15, 9, 8, 0, 14, 24, 4, 3, 17, 25, 23, 22, 6, 2, 19, 10, 20, 16, 18, 1, 13, 12, 7, 11},
}

func (r Reflector) InverseWiring() []int {
	return reflectorInverseWiringValues[r]
}
