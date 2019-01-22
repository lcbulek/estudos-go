package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	cuCentavo  = "Centavo"
	cuCentavos = "Centavos"
	cuReal     = "Real"
	cuReais    = "Reais"
	cuE        = "e"
	cuVirgula  = ","
	cuDe       = "de"
	cuCem      = "Cem"
	cuUmMil    = "Um Mil"
	cuMilhar   = "Mil"
	cuMilhao   = "Um Milhão"
	cuMilhoes  = "Milhões"
	cuBilhao   = "Um Bilhão"
	cuBilhoes  = "Bilhões"
)

var (
	unidades = [19]string{"Um", "Dois", "Três", "Quatro", "Cinco", "Seis", "Sete", "Oito", "Nove", "Dez", "Onze", "Doze",
		"Treze", "Quatorze", "Quinze", "Dezesseis", "Dezessete", "Dezoito", "Dezenove"}
	dezenas  = [9]string{"Dez", "Vinte", "Trinta", "Quarenta", "Cinquenta", "Sessenta", "Setenta", "Oitenta", "Noventa"}
	centenas = [9]string{"Cento", "Duzentos", "Trezentos", "Quatrocentos", "Quinhentos", "Seiscentos", "Setecentos",
		"Oitocentos", "Novecentos"}
)

func NumberToText(f float64) []string {
	w := []string{""}
	switch {
	case f >= 1 && f <= 19:
		return []string{unidades[int(f)-1]}
	case f >= 20 && f <= 99:
		{
			mod := math.Mod(f, 10)
			w = []string{dezenas[int(f/10)-1]}

			if mod > 0 {
				w = append(w, cuE)
			}
			w = append(w, NumberToText(mod)...)
			return w
		}
	case f >= 100 && f <= 999:
		{
			if f == 100 {
				w = append(w, cuCem)
				return w
			}
			mod := math.Mod(f, 100)
			w = []string{centenas[int(f/100)-1]}

			if mod > 0 {
				w = append(w, cuE)
			}
			w = append(w, NumberToText(mod)...)
			return w
		}
	case f >= 1000 && f <= 999999:
		{
			w := []string{""}
			value := int(f / 1000)

			if value == 1 {
				w = append(w, cuUmMil)
			} else {
				w = append(w, NumberToText(float64(value))...)
				w = append(w, cuMilhar)
			}

			mod := math.Mod(f, 1000)

			if mod <= 10 && mod >= 1 {
				w = append(w, cuE)
			} else if mod > 10 {
				w = append(w, cuVirgula)
			}

			w = append(w, NumberToText(math.Mod(f, 1000))...)
			return w
		}
	case f >= 1000000 && f <= 999999999:
		{
			w := []string{""}
			value := int(f / 1000000)
			if value == 1 {
				w = append(w, cuMilhao)
			} else {
				w = append(w, NumberToText(float64(value))...)
				w = append(w, cuMilhoes)
			}
			mod := math.Mod(f, 1000000)

			if mod == 0 {
				w = append(w, cuDe)
			} else if mod <= 1000 {
				w = append(w, cuE)
			} else if mod > 1000 {
				w = append(w, cuVirgula)
			}

			w = append(w, NumberToText(math.Mod(f, 1000000))...)
			return w
		}
	case f >= 1000000000 && f <= 999999999999:
		{
			w := []string{""}
			value := int(f / 1000000000)
			if value == 1 {
				w = append(w, cuBilhao)
			} else {
				w = append(w, NumberToText(float64(value))...)
				w = append(w, cuBilhoes)
			}
			mod := math.Mod(f, 1000000000)

			if mod == 0 {
				w = append(w, cuDe)
			} else if mod <= 1000000 {
				w = append(w, cuE)
			} else if mod > 1000000 {
				w = append(w, cuVirgula)
			}
			w = append(w, NumberToText(math.Mod(f, 1000000000))...)
			return w
		}
	}
	return nil
}

func limparAcasa(words []string) string {
	s := strings.Join(words, " ")
	s = strings.Replace(s, " ,  ", ", ", -1)
	s = strings.Replace(s, " , ", ", ", -1)
	s = strings.Trim(s, " ")
	return s
}

func main() {
	minusU := flag.Bool("u", false, "Uppercase")

	flag.Parse()
	flags := flag.Args()

	valor, _ := strconv.ParseFloat(flags[0], 64)
	inteiro, fracional := math.Modf(valor)
	fracional = math.RoundToEven(fracional * 100)

	w := []string{}

	w = NumberToText(inteiro)
	if inteiro == 1 {
		w = append(w, cuReal)
	} else {
		w = append(w, cuReais)
	}

	if fracional > 0 {
		w = append(w, cuE)
		w = append(w, NumberToText(fracional)...)
		if fracional == 1 {
			w = append(w, cuCentavo)
		} else {
			w = append(w, cuCentavos)
		}
	}

	fmt.Println(w)

	if *minusU {
		fmt.Println(strings.ToUpper(limparAcasa(w)))
	} else {
		fmt.Println(limparAcasa(w))
	}
}
