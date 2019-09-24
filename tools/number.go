package tools

// NumberGenerator is a predictable number generator
type NumberGenerator struct {
	seed          int
	currentNumber int
	min           int
	max           int
}

// NewNumberGenerator creates a new number generator with a specific seed
func NewNumberGenerator(seed int, min int, max int) NumberGenerator {
	return NumberGenerator{
		seed:          seed,
		currentNumber: 1,
		min:           min,
		max:           max,
	}
}

// NextPositive returns the next positive value of the generator
func (n *NumberGenerator) NextPositive() int {
	n.currentNumber = (1680748*n.currentNumber/n.seed + 111)

	if n.currentNumber < 0 {
		return (n.currentNumber%(n.max-n.min) + n.min) * -1
	}

	return n.currentNumber%(n.max-n.min) + n.min
}
