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
