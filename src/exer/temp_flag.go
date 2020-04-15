package main

import (
	"fmt"
	"flag"
)

type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit((c - 32) * 5 / 9) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

type celsiusFlag struct { Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temp %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temp")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}