package src

type Fork struct {
	Left string
	Right string
}

func (f *Fork) String() string {
	return f.Left + f.Right
}

func NewFork(left, right string) *Fork {
	return &Fork{Left: left, Right: right}
}

type Position struct {
	Start string
	Position string
	AtEnd bool
}

func (p *Position) String() string {
	return p.Position
}

func CheckAtEnd(positions []*Position) bool {
	for _, p := range positions {
		if !p.AtEnd {
			return false
		}
	}
	return true
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
			t := b
			b = a % b
			a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
			result = LCM(result, integers[i])
	}

	return result
}
