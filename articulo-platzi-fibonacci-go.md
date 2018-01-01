# ¿Qué es la secuencia de Fibonacci?
---

Es básicamente la secuencia de un conjunto de números que empieza por el 0 y el 1 y los siguientes son la suma de sus dos números anteriores.

Por ejemplo: **0**, **1**, 1, 2, 3, 5...

La primera vez que escuché de tan curiosa secuencia fue cuando ví la película del código Da Vinci:

[![Alt text](https://img.youtube.com/vi/UFbS1ssEmKc/0.jpg)](https://www.youtube.com/watch?v=UFbS1ssEmKc)

Volví a a ver esta secuencia cuando estaba entrando en la documentación de go en [golang.org](https://golang.org).

# ¿Qué es una función?

Pues en este punto ya debes saber cómo hacer para retornar una variable a través de una función.

Del mismo modo, las funciones también pueden retornar una función.

# Algoritmo de fibonacci con Go

Empezamos declarando el paquete main y librería a usar
```
package main

import "fmt"
```

Decidí declarar un slice **a** como variable global para guardar allí los valores y poder tener la secuencia.
```
var a []int
```

Luego la función clausura
```
func fibonacci() func(int) int {
	return func(x int) int {
		if x >= 2 {
			xf := a[x-1] + a[x-2]
			a = append(a, xf)
			return xf
		}

		a = append(a, x)
		return (a[x])
	}
}
```

Recibirá un valor representado en **x** que es el contador del bucle for dentro de la función maín.

Si dicho valor es 0 o 1, lo guardará en el slice gloabl **a** y lo retornará.

Si es 2 o mayor, sumará los dos valores anteriores guardados en el slice **a**, guardará este valor con la función **append(a, xf)** y finalmente lo retornará.

Para finalizar la función main
```
func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f(i))
	}
}
```

Se crea una variable que instancie la función **fibonacci()**, lo pase a un buble for quien se encargará de crear la secuencia dependiente del contador del for.

> Nota: Puedes cambiar el **20** del buble for por cualquier otro número y tendrás la secuencia de fibonacci.

# Conclusión

El código completo lo tienes aquí
```
package main

import "fmt"

// Fibonacci es una función que devuelve una
// función que devuelve un int.

var a []int

func fibonacci() func(int) int {
	return func(x int) int {
		if x >= 2 {
			xf := a[x-1] + a[x-2]
			a = append(a, xf)
			return xf
		}

		a = append(a, x)
		return (a[x])
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f(i))
	}
}
```

Y probarlo en la plataforma de Go [Aquí](https://play.golang.org/p/TaqzGyyO9Tc)

Este es el código de [golang.org](https://golang.org) respecto a la secuencia de Fibonacci
```
package main

import "fmt"

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()
	// Function calls are evaluated left-to-right.
	fmt.Println(f(), f(), f(), f(), f())
}
```

Que puedes probar [aquí](https://play.golang.org/p/RWLFDNhpKfE)

Como siempre ha sido un placer, te deseo un feliz y exitoso año nuevo 2018.

Cuéntame, ¿cómo crees que este código puede ser mejorado?