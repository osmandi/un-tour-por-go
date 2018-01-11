package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt struct {
	Value float64
	Err   string
}

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("No se puede sacar la raíz cuadrada de %.2f, %s",
		e.Value, e.Err)
}

func Sqrt(x float64) (float64, error) {
	var zi float64 = 1 // valor inicial
	var z float64 = 0  // valor siguiente
	var run bool = true

	if x < 0 {
		return 0, &ErrNegativeSqrt{x, "Números negativos no soportado"}
	}

	for run {
		z = zi - ((zi*zi - x) / (2 * zi))
		err := math.Abs((zi-z)/zi) * 100
		//fmt.Printf("Z0 es %.2f y Z es , corridas int%.2f Error: %.2f%%\n", zi, z, err)
		zi = z
		if err < 0.0001 {
			run = false
		}
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(-4))
}
