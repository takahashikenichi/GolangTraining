package weightconv

import "fmt"

type Pound float64
type KiloGram float64

const (
	coeffience = 0.45359237
)

func PToK(p Pound) KiloGram {
	return KiloGram(p * coeffience)
}

func KToP(k KiloGram) Pound {
	return Pound(k / coeffience)
}

func (p Pound) String() string {
	return fmt.Sprintf("%glb", p)
}

func (k KiloGram) String() string {
	return fmt.Sprintf("%gkg", k)
}
					}
