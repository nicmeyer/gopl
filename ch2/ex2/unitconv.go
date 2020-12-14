package unitconv

import "fmt"

// temperature.
type Celsius float64
type Fahrenheit float64
type Kelvin float64

// length.
type Feet float64
type Meters float64

// weight.
type Pounds float64
type Kilograms float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// temperature String methods.
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

// length String methods.
func (f Feet) String() string   { return fmt.Sprintf("%gft", f) }
func (m Meters) String() string { return fmt.Sprintf("%gm", m) }

// weight String methods.
func (p Pounds) String() string    { return fmt.Sprintf("%glb", p) }
func (k Kilograms) String() string { return fmt.Sprintf("%gkg", k) }
