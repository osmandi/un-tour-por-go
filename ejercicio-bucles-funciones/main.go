package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	var zi float64 = 1 // valor inicial
	var z float64 = 0  // valor siguiente
	var run bool = true
	var corridas int = 0
	for run {
		corridas += 1
		z = zi - ((zi*zi - x) / (2 * zi))
		err := math.Abs((zi-z)/zi) * 100
		//fmt.Printf("Z0 es %.2f y Z es , corridas int%.2f Error: %.2f%%\n", zi, z, err)
		zi = z
		if err < 0.0001 {

			run = false

		}

	}

	return z, corridas
}

func main() {
	fmt.Println("Este programa determina la raiz cuadrada de un número")
	fmt.Print("Ingresa un valor entero o decimal: ")
	var valor float64
	fmt.Scanf("%f", &valor)
	resultado, corrida := Sqrt(valor)
	fmt.Printf("Resultado por método newton: %.2f Resultado por librería mat.Sqrt: %.2f en %d corridas\n",
		resultado, math.Sqrt(valor), corrida)
}
