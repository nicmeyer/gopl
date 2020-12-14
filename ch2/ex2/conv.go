package unitconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k) + AbsoluteZeroC }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// FToM converts a Feet lenght to Meters.
func FToM(f Feet) Meters { return Meters(f / 3.2808) }

// MToF converts a Meters length to Feet.
func MToF(m Meters) Feet { return Feet(m * 3.2808) }

// PToK converts a Pounds weight to Kilogram.
func PToK(p Pounds) Kilograms { return Kilograms(p * 0.4536) }

// KToP converts a Kilograms weight to Pounds.
func KToP(k Kilograms) Pounds { return Pounds(k / 0.4536) }
