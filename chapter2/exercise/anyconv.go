package anyconv

import "fmt"

type Celsius float64
type Fahremheit float64
type Feet float64
type Meter float64
type Kg float64
type Pound float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	Boiling       Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahremheit) String() string { return fmt.Sprintf("%g°F", f) }
func (ft Feet) String() string      { return fmt.Sprintf("%gft", ft) }
func (m Meter) String() string      { return fmt.Sprintf("%gm", m) }
func (kg Kg) String() string        { return fmt.Sprintf("%gkg", kg) }
func (p Pound) String() string      { return fmt.Sprintf("%glb") }
