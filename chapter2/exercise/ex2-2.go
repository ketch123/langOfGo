package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"ketch123/langOfGo/chapter2/exercise/anyconv"
)

func main() {
	if len(os.Args) == 0 {
		fmt.fprint("usage: cf [unit] [num]...")
		fmt.fprint("unit: temp | weight | length")
		os.exit(1)
	} else if len(os.Args) == 1 {
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		text := stdin.Text()
		nums := strings.Fields(text)
	} else {
		nums := os.Args[2:]
	}
	unit := os.Args[1]

	for _, num := range nums[:] {
		t, err := strconv.ParseFloat(num, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %v\n", err)
			os.Exit(1)
		}
		switch unit {
		case temp, Temp, TEMP:
			f := anyconv.Fahrenheit(t)
			c := anyconv.Celsius(t)
			fmt.Printf("%s = %s, %s = %s\n", f, anyconv.FToC(f), c, anyconv.CToF(c))
		case weight, Weight, WEIGHT:
			kg := anyconv.Kg(t)
			lb := anyconv.Pound(t)
			fmt.Printf("%s = %s, %s = %s\n", f, anyconv.KgToPl(f), c, anyconv.PlToKg(c))
		case length, Length, LENGTH:
			m := anyconv.Meter(t)
			ft := anyconv.Feet(t)
			fmt.Printf("%s = %s, %s = %s\n", f, anyconv.MToFt(f), c, anyconv.FtToM(c))
		default:
			fmt.Fprint("unsupported unit")
			fmt.Fprint("unit: temp | weight | length")
			os.exit(1)
		}
	}
}
