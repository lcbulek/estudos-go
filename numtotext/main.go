package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	numbers = [19]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve",
		"thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tenths = [9]string{"ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
)

func NumberToText(f float64) []string {
	w := []string{""}
	switch {
	case f >= 1 && f <= 19:
		return []string{numbers[int(f)-1]}
	case f >= 20 && f <= 99:
		{
			w = []string{tenths[int(f/10)-1]}
			w = append(w, NumberToText(math.Mod(f, 10))...)
			return w
		}
	case f >= 100 && f <= 999:
		{
			w = []string{numbers[int(f/100)-1]}
			w = append(w, "hundred")
			w = append(w, NumberToText(math.Mod(f, 100))...)
			return w
		}
	case f >= 1000 && f <= 999999:
		{
			w = NumberToText(f / 1000)
			w = append(w, "thousand")
			w = append(w, NumberToText(math.Mod(f, 1000))...)
			return w
		}
	case f >= 1000000 && f <= 999999999:
		{
			w = NumberToText(f / 1000000)
			w = append(w, "million")
			w = append(w, NumberToText(math.Mod(f, 1000000))...)
			return w
		}
	case f >= 1000000000 && f <= 999999999999:
		{
			w = NumberToText(f / 1000000000)
			w = append(w, "billion")
			w = append(w, NumberToText(math.Mod(f, 1000000000))...)
			return w
		}
	}
	return nil
}

func main() {
	valor, _ := strconv.ParseFloat(os.Args[1], 64)
	inteiro, fracional := math.Modf(valor)
	fracional = math.RoundToEven(fracional * 100)

	fmt.Println(inteiro, fracional)

	w := NumberToText(inteiro)
	if inteiro == 1 {
		w = append(w, "dollar")
	} else {
		w = append(w, "dollars")
	}
	if fracional > 0 {
		w = append(w, "and")
		w = append(w, NumberToText(fracional)...)
		if fracional == 1 {
			w = append(w, "cent")
		} else {
			w = append(w, "cents")
		}
	}
	fmt.Println(w)
}
