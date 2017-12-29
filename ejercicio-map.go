package main

/*
Implementa ContadorPalabras.
Debería devolver un map con el número de
veces que una "palabra" aparece en la cadena s.
La función wc.Test ejecuta un caso de prueba ejecutando
la función implementada e imprime éxito o fallo.
*/

import (
	"strings"

	"golang.org/x/tour/wc"
)

func ContadorPalabras(s string) map[string]int {

	s = "I am learning Go!"
	countMap := strings.Fields(s)
	m := make(map[string]int)
	for i := range countMap {
		m[countMap[i]] += 1
	}

	return m
}

func main() {
	wc.Test(ContadorPalabras)
}
