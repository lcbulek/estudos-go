package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	cuE        = "e"
	cuVirgula  = ","
	cuCentavo  = "Centavo"
	cuCentavos = "Centavos"
	cuReal     = "Real"
	cuReais    = "Reais"
)

var (
	unidades = [20]string{"", "Um", "Dois", "Três", "Quatro", "Cinco", "Seis", "Sete", "Oito", "Nove",
		"Dez", "Onze", "Doze", "Treze", "Quatorze", "Quinze", "Dezesseis", "Dezessete", "Dezoito", "Dezenove"}

	dezenas = [8]string{"Vinte", "Trinta", "Quarenta", "Cinquenta", "Sessenta", "Setenta", "Oitenta", "Noventa"}

	centenas = [9]string{"Cento", "Duzentos", "Trezentos", "Quatrocentos", "Quinhentos", "Seiscentos", "Setecentos",
		"Oitocentos", "Novecentos"}

	cem = "Cem"

	mil = "Mil"

	milhar = [3]string{"Milhão", "Bilhão", "Trilhão"}

	milhares = [3]string{"Milhões", "Bilhões", "Trilhões"}
)

func getMoneyInWords(f float64) []string {
	switch {
	case f < 20:
		return []string{unidades[int(f)]}
	case f < 100:
		return getDezenas(f)
	case f == 100:
		return []string{cem}
	case f < 1000:
		return getCentenas(f)
	case f < 1000000:
		return getMilhares(f)
	case f < 1000000000:
		return getMilhoes(f)
	case f < 1000000000000:
		return getBilhoes(f)
	case f < 1000000000000000:
		return getTrilhoes(f)
	}
	return nil
}

func getTrilhoes(f float64) (words []string) {
	var w []string
	value := int(f / 1000000000000)
	if value == 1 {
		words = append(words, "Um")
		words = append(words, milhar[2])
	} else {
		w = getMoneyInWords(float64(value))
		words = append(w, milhares[2])
	}

	f = f - float64(value*1000000000000)
	if f > 1 {
		words = append(words, cuVirgula)
		words = append(words, getMoneyInWords(f)...)
	}
	if f == 1 {
		words = append(words, cuE)
		words = append(words, getMoneyInWords(f)...)
	}
	return words
}

func getBilhoes(f float64) (words []string) {
	var w []string
	value := int(f / 1000000000)
	if value == 1 {
		words = append(words, "Um")
		words = append(words, milhar[1])
	} else {
		w = getMoneyInWords(float64(value))
		words = append(w, milhares[1])
	}

	f = f - float64(value*1000000000)
	if f > 1 {
		words = append(words, cuVirgula)
		words = append(words, getMoneyInWords(f)...)
	}
	if f == 1 {
		words = append(words, cuE)
		words = append(words, getMoneyInWords(f)...)
	}

	return words
}

func getMilhoes(f float64) (words []string) {
	var w []string
	value := int(f / 1000000)
	if value == 1 {
		words = append(words, "Um")
		words = append(words, milhar[0])
	} else {
		w = getMoneyInWords(float64(value))
		words = append(w, milhares[0])
	}

	f = f - float64(value*1000000)
	if f > 1 {
		words = append(words, cuVirgula)
		words = append(words, getMoneyInWords(f)...)
	}
	if f == 1 {
		words = append(words, cuE)
		words = append(words, getMoneyInWords(f)...)
	}
	return words
}

func getMilhares(f float64) (words []string) {
	var w []string
	value := int(f / 1000)
	if value == 1 {
		words = append(words, "Um")
		words = append(words, mil)
	} else {
		w = getMoneyInWords(float64(value))
		words = append(w, mil)
	}

	f = f - float64(value*1000)
	if f > 0 {
		words = append(words, cuE)
		words = append(words, getMoneyInWords(f)...)
	}

	return words
}

func getCentenas(f float64) []string {
	value := centenas[int(f/100)-1]
	words := []string{value}

	mod := math.Mod(f, 100)
	if mod != 0 {
		sobra := getMoneyInWords(mod)
		words = append(words, cuE)
		words = append(words, sobra...)
	}
	return words
}

func getDezenas(f float64) []string {
	value := dezenas[int(f-20)/10]
	words := []string{value}

	mod := math.Mod(f, 10)
	if mod != 0 {
		words = append(words, cuE)
		words = append(words, unidades[int(mod)])
	}
	return words
}

func limparAcasa(words []string) string {
	s := strings.Join(words, " ")
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
	words := []string{""}

	if valor == 0 {
		os.Exit(0)
	}

	// obter o valor por extenso da parte inteira
	words = getMoneyInWords(inteiro)

	if inteiro > 1 {
		words = append(words, cuReais)
	} else if inteiro == 1 {
		words = append(words, cuReal)
	}

	if inteiro >= 1 && fracional >= 1 {
		words = append(words, cuE)
	}

	// obter o valor por extenso da fração
	words = append(words, getMoneyInWords(fracional)...)

	if fracional > 1 {
		words = append(words, cuCentavos)
	} else if fracional == 1 {
		words = append(words, cuCentavo)
	}

	if *minusU {
		fmt.Println(strings.ToUpper(limparAcasa(words)))
	} else {
		fmt.Println(limparAcasa(words))
	}
}
