package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		case weight, Weight, WEIGHT:
			kg := tempconv.Kg(t)
			lb := tempconv.Pound(t)
			fmt.Printf("%s = %s, %s = %s\n", f, tempconv.KgToPl(f), c, tempconv.PlToKg(c))
		case length, Length, LENGTH:
			m := tempconv.Meter(t)
			ft := tempconv.Feet(t)
			fmt.Printf("%s = %s, %s = %s\n", f, tempconv.MToFt(f), c, tempconv.FtToM(c))
		default:
			fmt.Fprint("unsupported unit")
			fmt.Fprint("unit: temp | weight | length")
			os.exit(1)
		}
	}
}
