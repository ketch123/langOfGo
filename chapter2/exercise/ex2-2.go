package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ketch123/langOfGo/chapter2/exercise/anyconv"
	//"./anyconv"
)

func main() {
	var nums []string
	if len(os.Args) == 1 {
		fmt.Println("usage: cf [unit] [num]...")
		fmt.Println("unit: temp | weight | length")
		os.Exit(0)
	} else if len(os.Args) == 2 {
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		text := stdin.Text()
		nums = strings.Fields(text)
	} else {
		nums = os.Args[2:]
	}
	unit := os.Args[1]

	for _, num := range nums[:] {
		t, err := strconv.ParseFloat(num, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %v\n", err)
			os.Exit(1)
		}
		switch unit {
		case "temp", "Temp", "TEMP":
			f := anyconv.Fahrenheit(t)
			c := anyconv.Celsius(t)
			fmt.Printf("%s = %s, %s = %s\n", f, anyconv.FToC(f), c, anyconv.CToF(c))
		case "weight", "Weight", "WEIGHT":
			kg := anyconv.Kg(t)
			lb := anyconv.Pound(t)
			fmt.Printf("%s = %s, %s = %s\n", kg, anyconv.KgToPl(kg), lb, anyconv.PlToKg(lb))
		case "length", "Length", "LENGTH":
			m := anyconv.Meter(t)
			ft := anyconv.Feet(t)
			fmt.Printf("%s = %s, %s = %s\n", m, anyconv.MToFt(m), ft, anyconv.FtToM(ft))
		default:
			fmt.Println("unsupported unit")
			fmt.Println("unit: temp | weight | length")
			os.Exit(0)
		}
	}
}
