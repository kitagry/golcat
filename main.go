package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/gookit/color"
)

func rainbow(i int) color.RGBColor {
	freq := 0.03

	var (
		red   = math.Sin(freq*float64(i))*127 + 128
		green = math.Sin(freq*float64(i)+2*math.Pi/3)*127 + 128
		blue  = math.Sin(freq*float64(i)+4*math.Pi/3)*127 + 128
	)

	return color.RGB(uint8(red), uint8(green), uint8(blue))
}

func Println(text string, i int) {
	for index, t := range text {
		rainbow(index + i*3).Print(string(t))
	}
	fmt.Println("")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	i := 0
	for scanner.Scan() {
		Println(scanner.Text(), i)
		i++
	}
}
