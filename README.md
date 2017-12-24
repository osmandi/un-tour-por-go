# Tutorial de tour por Go

Fuente: [Un tour por Go](https://go-tour-es.appspot.com)

Para ejecutar localmente (dice que hay más ejercicios, pero lo veo igual :P), una vez instalado Go. Instalar Gotour (inglés)
```
go get code.google.com/p/go-tour/gotour
```

En castellano
```
go get github.com/rcostu/go-tour-es/gotoures
```

Para correr la versión en inglés `gotour` y en español (no me funciona) `gotoures`

# Paquetes
Los programas comienzan su ejecución en el paquete main.

Por convención, el nombre del paquete es el mismo que el último elemento de la ruta de importación.

# Importación

Hay dos maneras de importar paquetes
```
import "fmt"
import "math"
```

o

```
import (
    "fmt"
    "math"
)
```

# Identificadores exportados
Tras importar un paquete, puedes hacer referencia a los identificadores que exporta.

En Go, un identificador es exportado si empieza por una mayúscula.

# Funciones
Una función puede tener cero o más argumentos.

Observa que el tipo se indica después del nombre de la variable.

Cuando dos o más parámetros consecutivos de la función son del mismo tipo, puedes omitir uno de ellos.

Ejemplo de 
```
x int, y int
```

a

```
x, y int
```

# Múltiples resultados

Una función puede devolver varios resultados

# Resultados nombrados
Las funciones tienen parámetros; en Go los resultados pueden ser nombrados y actuar como variables. Se les denomina "variables de retorno"

Si las variables de retorno tienen un nombre, una sentencia "return" sin argumentos devuelve el valor actual de dichas variables.

# Variables
La sentencia **var** declara una lista de variables; como en la lista de argumentos de las funciones, el tipo se indica al final.

# Variables inicializadas
La declaración de variables permite inicializaciones, una por variable.

Si se inicializa una variable, el tipo puede omitirse; la variable adoptará el tipo de valor con el que ha sido inicializada.

# Declaración implicita de variables
Dentro de una función, puede utilizarse la sentencia de asignación **:=** en lugar de la declaración **var**.

(Fuera de una función, todas las declaraciones de variables comienzan con la palabra clave ***var**+ y el operando ***:=*** no está disponible).

# Tipos básicos
Los tipos básicos en Go son:
```
bool

string

int  int8  int16  int32  int64
uint uint8 unit16 uint32 uint64 uintprt //solo números positivos

byte // alias para uint8
    // representa un punto Unicode

float32 float64 // Decimales

complex64 complex128 // Números complejos
```

# Constantes
Las constantes se declaran como las variables, pero con la palabra ***const***.

Las constantes pueden ser cadenas, booleanas o numéricas.

# Constantes numéricas
Las constantes numéricas son valores de alta precisión.

Una constante sin un tipo definido tiene el tipo necesitado según el contexto en el que se declara.

Intenta también imprimir el valor needInt(Big)

# For

Go tiene solo un operando para definir bucles, los bucles for.

El bucle for básico es muy parecido al que se utiliza en C o Javam salgo que los () desaperecen, ni siquiera son opcionales y las {} son obligatoras.

# For-while
Un bucle while de C se transforma en un bucle for en Go

# For-eterno
Si omites la condición del bucle, es un bucle infinito de manera que un bucle infinito se escribe de manera compacta.

# If
La instrucción if es similar a la sentencia en C o Java, salvo que los paréntesis () desaparecen (ni siquiera son opcionales) y las llaves {} son obligatorias.

# If con instrucción inicial
Al igual que en la sentencia for, la sentencia if puede empezar con una instrucción de inicialización que se ejecutará antes de evaluar la condición.

Las variables declaradas por la instruccion de inicialización son únicamente visibles en el ámbito del if.

(intenta usar v en la última sentencia return) => La variable ***v*** es indefinida, porque se define en el if

# If y else
Las variables declaradas dentro de la instrucción de inicicalización de un if son también visibles dentro de los bloques else.

# Ejercicio: Bucles y funciones
Una forma sencila de jugar con funciones y bucles es implementar la funcionalidad de la raiz cuadrada utilizando el método Newton.

En este caso el método de Newton aproxima Sqrt(x) tomando un punto inicial z y repitiendo:

```
z=z-((z^2-x)/2z)
```

Para empezar, simplemente repite el cálculo 10 veces y mira cómo de cerca estas de la solución para distintos valores (1,2,3,...).

Después cambia la condición del buble para cuando el valor deje de cambiar (o solo cambie con un delta muy pequeño). Mira si esto ocurre con más  menos iteraciones. ¿Cómo estás de cerca comparando con math.Sqr?

Pista: Para declarar e inicalizar un valor decimal dale un valor decimal o utiliza la conversión:
z:=float64(0)
z:=0.0

# Estructuras y tipos

Una estructura (struct) es un registro de variables dentro de un mismo tipo.

(Y una declaración type declara un nuevo tipo de datos)

# Campos de una estructura

Los campos de una estructura son accesibles mediante el operador . (punto)

# Punteros

Go posee punteros, pero no tiene aritmética de punteros (como C).

Los campos de las estructuras pueden accederse a través de un puntero a una estructura.

La indirección del puntero es transparente al programador.

# Estructuras literales

Una estructura literal denota una nueva instancia de la estructura que muestra los valores de sus campos.

Puedes mostrar solo un subconjunto de los campos utilizando la sintaxis Name: (y el orden de los campos nombrados es irrelevante)

El prefijo especial & construye un puntero al espacio donde la nueva estructura se aloja.

# La función "new"

La expresión new(T) aloja en memoria un valor T inicializado a 0 y retorna un puntero al mismo tiempo.

```
var t *T = new(T)

o

t := new(T)
```
