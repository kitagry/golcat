package main

import (
	"math"

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

func Print(text string) {
	for i, t := range text {
		rainbow(i).Print(string(t))
	}
}

func main() {
	text := "あいうえおカキクケコさしすせそ\nたちつてとナニヌネノはひふへほ"
	Print(text)
}
