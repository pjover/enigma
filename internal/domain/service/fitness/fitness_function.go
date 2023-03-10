package fitness

const epsilon = 3e-10

type FitnessFunction interface {
	score(text string) float32
}
