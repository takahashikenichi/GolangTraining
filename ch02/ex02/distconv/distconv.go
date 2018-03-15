package distconv

import "fmt"

type Meter float64
type Feet float64

const (
	weightCoeffience = 3.2808
)

func MToF(m Meter) Feet {
	return Feet(m * weightCoeffience)
}

func FToM(f Feet) Meter {
	return Meter(f / weightCoeffience)
}

func (m Meter) String() string {
	return fmt.Sprintf("%gM", m)
}

func (f Feet) String() string {
	return fmt.Sprintf("%gF", f)
}
