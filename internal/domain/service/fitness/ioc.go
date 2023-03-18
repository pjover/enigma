package fitness

func NewIoc() FitnessFunction {
	return ioc{}
}

type ioc struct {
}

func (i ioc) score(text string) float32 {
	histogram := make([]int, 26)
	for _, c := range text {
		histogram[c-'A']++
	}

	n := len(text)
	total := 0.0
	for _, v := range histogram {
		total += float64(v * (v - 1))
	}

	return float32(total / float64(n*(n-1)))
}
