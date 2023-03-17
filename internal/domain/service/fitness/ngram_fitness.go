package fitness

import (
	"encoding/csv"
	"io"
	"math"
	"os"
	"strconv"
)

type NgramFitness struct {
	size         int
	length       int
	fileName     string
	computeIndex func([]int) int
	scores       []float32
}

type NgramFitnessType uint

const (
	Bigram NgramFitnessType = iota
	Trigram
	Quadram
)

func NewNgramFitness(ngramFitnessType NgramFitnessType) (*NgramFitness, error) {
	var ngramFitness *NgramFitness
	switch ngramFitnessType {
	case Bigram:
		ngramFitness = &NgramFitness{
			size:     2,
			length:   826,
			fileName: "bigrams.csv",
			computeIndex: func(index []int) int {
				return (index[0] << 5) | index[1]
			},
		}
	case Trigram:
		ngramFitness = &NgramFitness{
			size:     3,
			length:   26426,
			fileName: "trigrams.csv",
			computeIndex: func(index []int) int {
				return (index[0] << 10) | (index[1] << 5) | index[2]
			},
		}
	case Quadram:
		ngramFitness = &NgramFitness{
			size:     4,
			length:   845626,
			fileName: "quadgrams.csv",
			computeIndex: func(index []int) int {
				return (index[0] << 15) | (index[1] << 10) | (index[2] << 5) | index[3]
			},
		}
	}
	return buildNgramFitness(ngramFitness)
}

func buildNgramFitness(ngramFitness *NgramFitness) (*NgramFitness, error) {
	ngramFitness.scores = buildScoresWithDefaultValues(ngramFitness.length)

	file, err := os.Open(ngramFitness.fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	r := csv.NewReader(file)
	r.FieldsPerRecord = 2
	r.TrimLeadingSpace = true

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		i := ngramFitness.computeIndex(parseIndex(record[0]))
		f, _ := strconv.ParseFloat(record[1], 32)
		ngramFitness.scores[i] = float32(f)
	}
	return ngramFitness, nil
}

func buildScoresWithDefaultValues(length int) []float32 {
	defaultValue := float32(math.Log10(epsilon))
	var scores = make([]float32, length)
	for i := range scores {
		scores[i] = defaultValue
	}
	return scores
}

func parseIndex(text string) []int {
	var arr []int
	for i := 0; i < len(text); i++ {
		a := int(text[i]) - 65
		arr = append(arr, a)
	}
	return arr
}

func (f *NgramFitness) Score(text string) float32 {
	fitness := float32(0)
	var index = make([]int, f.size)
	index[0] = 0
	for i := 1; i < f.size; i++ {
		index[i] = int(text[i-1]) - 65
	}
	for i := f.size - 1; i < len(text); i++ {
		index = f.updateIndex(text, index, i)
		index := f.computeIndex(index)
		fitness += f.scores[index]
	}
	return fitness
}

func (f *NgramFitness) updateIndex(text string, index []int, i int) []int {
	for j := 0; j < f.size-1; j++ {
		index[j] = index[j+1]
	}
	index[f.size-1] = int(text[i]) - 65
	return index
}
