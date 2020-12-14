// Uc converts its numeric argument to Celsius, Fahrenheit
// Feet, Meters, Pounds and Kilograms
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nicmeyer/gopl/ch2/ex2/unitconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		i, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "uc: %v\n", err)
			os.Exit(1)
		}
		{
			f := unitconv.Fahrenheit(i)
			c := unitconv.Celsius(i)
			k := unitconv.Kelvin(i)
			fmt.Printf("%s = %s, %s = %s, %s = %s\n",
				f, unitconv.FToC(f), c, unitconv.CToF(c), k, unitconv.KToC(k))
		}
		{
			f := unitconv.Feet(i)
			m := unitconv.Meter(i)
			fmt.Printf("%s = %s, %s = %s\n",
				f, unitconv.FToM(f), m, unitconv.MToF(m))
		}
		{
			p := unitconv.Pound(i)
			k := unitconv.Kilogram(i)
			fmt.Printf("%s = %s, %s = %s\n",
				p, unitconv.PToK(p), k, unitconv.KToP(k))
		}

	}
}
