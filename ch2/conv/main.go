// exercise 2.2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/wngtk/gopl-solutions/ch2/tempconv"
)

type Inch float64
type Meter float64

func (in Inch) String() string { return fmt.Sprintf("%g in", in) }
func (m Meter) String() string { return fmt.Sprintf("%g m", m) }

func InchToMeter(in Inch) Meter {
	return Meter(in * 0.0254)
}

func MeterToInch(m Meter) Inch {
	return Inch(m / 0.0254)
}

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf2: %v\n", err)
			os.Exit(1)
		}

		convert(t)
	}

	if len(os.Args) == 1 {
		in := bufio.NewScanner(os.Stdin)
		for in.Scan() {
			t, err := strconv.ParseFloat(in.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf2: %v\n", err)
				os.Exit(1)
			}

			convert(t)
		}
	}
}

func convert(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)

	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))

	i := Inch(t)
	m := Meter(t)
	fmt.Printf("%s = %s, %s = %s\n", i, InchToMeter(i), m, MeterToInch(m))
}
