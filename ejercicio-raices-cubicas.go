package main

import (
	"fmt"
	"math"
)

func raizCubica(x float64) (float64, int) {
	var zi float64 = 1 // valor inicial
	var z float64 = 0  // valor iterable
	var run bool = true
	var corridas int = 0
	for run {
		corridas += 1
		z = zi - ((math.Pow(zi, 3) - x) / (3 * math.Pow(zi, 2)))
		err := math.Abs((zi-z)/zi) * 100
		zi = z
		if err < 0.0001 {
			run = false
		}
	}

	return z, corridas
}

func main() {
	var valor float64
	fmt.Println("Ingresa un valor para obtener la raíz cuadrada")
	fmt.Scanf("%f", &valor)
	resultado, corridas := raizCubica(valor)
	fmt.Printf("Resultado por método newton: %0.2f en %d corridas", resultado, corridas)
}
