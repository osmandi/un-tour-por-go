package main

import "fmt"

const (
	// Crea un gran número al cambiar 1 bit a la izquierda 100 lugares
	// En otras palabras, el número binario que es 1 seguido de 100 ceros
	Big = 1 << 100

	// Cambia de nuevo a la derecha 99 lugares, por lo que terminamos con 1 << 1 o

	// Cambia de nuevo a la derecha 99 lugares, por lo que terminamos con 1 << 1 o 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(float64(Big))
	fmt.Printf("Small: %d", Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
