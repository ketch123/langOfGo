package tempconv

import "fmt"

type Celsius float64
type Fahremheit float64
type Kelvinscale float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	Boiling       Celsius = 100
)

func (c Celsius) String() string     { return fmt.Sprintf("%g°C", c) }
func (f Fahremheit) String() string  { return fmt.Sprintf("%g°F", f) }
func (k Kelvinscale) String() string { return fmt.Sprintf("%g°K"), k }
