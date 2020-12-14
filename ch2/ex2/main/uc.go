// Uc converts its numeric argument to Celsius, Fahrenheit
// Feet, Meters, Pounds and Kilograms
package main

import (
	"fmt"
	"os"
	"strconv"

	"local/ch2/ex2"
)

func printTempConv(t float64) {
	f := ex2.Fahrenheit(t)
	c := ex2.Celsius(t)
	k := ex2.Kelvin(t)
	fmt.Printf("%s = %s, %s = %s, %s = %s\n",
		f, ex2.FToC(f), c, ex2.CToF(c), k, ex2.KToC(k))
}

func printLenConv(l float64) {
	f := ex2.Feet(l)
	m := ex2.Meter(l)
	fmt.Printf("%s = %s, %s = %s\n",
		f, ex2.FToM(f), m, ex2.MToF(m))

}

func printWeightConv(w float64) {
	p := ex2.Pound(w)
	k := ex2.Kilogram(w)
	fmt.Printf("%s = %s, %s = %s\n",
		p, ex2.PToK(p), k, ex2.KToP(k))
}

func main() {
	for _, arg := range os.Args[1:] {
		i, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "uc: %v\n", err)
			os.Exit(1)
		}
		printTempConv(i)
		printLenConv(i)
		printWeightConv(i)
	}
}
