// Uc converts its numeric argument to Celsius, Fahrenheit
// Feet, Meters, Pounds and Kilograms
package main

import (
	"fmt"
	"os"
	"strconv"

	"local/ch2/ex2/unitconv"
)

func printTempConv(i float64) {
	f := unitconv.Fahrenheit(i)
	c := unitconv.Celsius(i)
	k := unitconv.Kelvin(i)
	fmt.Printf("%s = %s, %s = %s, %s = %s\n",
		f, unitconv.FToC(f), c, unitconv.CToF(c), k, unitconv.KToC(k))
}

func printLenConv(i float64) {
	f := unitconv.Feet(i)
	m := unitconv.Meter(i)
	fmt.Printf("%s = %s, %s = %s\n",
		f, unitconv.FToM(f), m, unitconv.MToF(m))
}

func printWeightConv(i float64) {
	p := unitconv.Pound(i)
	k := unitconv.Kilogram(i)
	fmt.Printf("%s = %s, %s = %s\n",
		p, unitconv.PToK(p), k, unitconv.KToP(k))
}

func printConvs(arg string) {
	i, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "uc: %v\n", err)
		os.Exit(1)
	}
	printTempConv(i)
	printLenConv(i)
	printWeightConv(i)
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			printConvs(arg)
		}
	} else {
		var arg string
		_, err := fmt.Scanf("%s", &arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "uc: %v\n", err)
			os.Exit(1)
		}
		printConvs(arg)
	}
}
