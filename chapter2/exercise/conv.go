package anyconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func FtToM(ft Feet) Meter { return Meter(f / 3.2808) }
func MToFt(m Meter) Feet  { return Feet(m * 3.2808) }

func KgToPl(kg Kg) Pound { return Pound(kg * 2.2046) }
func PlToKg(pl Pound) Kg { return Pound(pl / 2.2046) }
