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

# Slices

Un slice apunta a un array de valores y posee un tamaño fijo.

[]T es un slice con elementos

# Creando Slices con make

Los slices son creados con la función *make*. Funciona alojando un array inicializado a 0 y retornando un slice que apunta a ese array:

```
a := make([]int, 5) // len(a)=5
```
 
Los slices tienen un tamaño y una capacidad. La capacidad de un slice es el tamaño máximo que el slice puede crecer dentro del array al que apunta.

Para especificar una capacidad basta con pasar un tercer argumento a *make*:

```
b:= make([]int,0,5) //leb(b)=0, cap(b)=5
```

Los slices pueden crecer reasignándose (por encima de su capacidad):

```
b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:] // len(b)=4, cap(b)=4
```

# Slices nil
El valor por defecto de un slice es nil.

Un slice nil tiene un tamaño y una longitud de 0.

# Range
La forma *range* de un bucle *for* itera sobre elementos de un slice o un map.

> Son como los foreach en Java.

Para descartar un return de una función, usar "_" donde corresponde el lugar de la variable.

# Ejercicio: Slices
Implementa la función Pic. Debería devolver un slice de tamaño dy, siendo cada uno de los elementos un slice de dx enteros sin signo de 8 bits. Cuando ejecutes el programa, mostrará tu dibujo, interpretando los números enteros como una escala de grises (bueno, escala de azules).

La elección de la imagen es de tu elección. Algunas funciones interesantes pueden ser x^y, (x+y)/2, x*y-

(Necesitas usar un bucle para reservar memoria para cada []uint8 dentro de la matriz [][]uint8)

(Usa uint8(intValue) para convertir entre tipos.)


# Maps
Un map relaciona claves y valores.

Los Maps debe crearse con la función *make* (nunca con *new*) antes de su uso; el map *nil* está vacío y no puede ser asignado.

# Maps literales
Los map literales son como las estructuras literales, pero las claves son obligatorias

> Se utiliza de la misma manera que los slices para comprobar si un valor existe, en caso de que sí se obtiene un valor y el ok.

# Ejercicio Maps

Implementa *ContadorPalabras*. Debería devolver un map con el número de veces que una "palabra" aparece en la cadena *s*. La función *wc.Test* ejecuta un caso de prueba ejecutando la función implementada e imprime éxito o fallo.

Puedes encontrar ayuda en [strings.Fields](http://golang.org/pkg/strings/#Fields)

# Funciones con variables

Las funciones también son valores

# Funciones son clausuras

Y las funciones son clausuras completas.

La función *adder* retorna una clausura (o función anónima). Cada clausura está vinculada a su variable *sum* correspondiente.

# Ejercicio: La calusura de Fibonacci

Vamos a divertirnos un poco con las funciones.

Implementa una función de fibonacci que devuelva una función (o clausura) que devuelva los sucesivos números de fibonacci.

# Switch

Probablemente ya sabes cómo iba a ser la cláusula *switch*.

El cuerpo de un caso sale automáticamente de la cláusula *switch*, a menos que termina con una sentencia *fallthrough*, que provocaría quesiguiera ejecutando el siguiente caso contemplado.


> Se recomienda esta cláusula para esos casos donde necesitas usar más de dos sentencias if.

# Orden de evaluación de un switch

Los casos de un *Switch* evalúan los casos de arriba a abajo, parando cuando se encuentra un caso satisfactorio.

# Switch sin condición
Un *switch* sin condición es lo mismo que *switch true*.

Esta construcción puede ser una manera clara de escribir cadenas largas if-then-else.

# Ejercicio avanzado: Raíces cúbicas

Veamos cuál es el soporte de Go para números complejos mediante los tipos complex64 y complex128. Para raíces cuadradas el método de Newton se basa en repeticiones.

z = (z)-((z^3-x)/3z^2)

Busca la raiz cúbica de 2, para asegurarte que el algoritmo funciona. Existe una función Pow en el paquete match/cmplx.

# Métodos

Go no tiene clases. De todas formas, puedes definir métodos para tipos struct.

El receptor del método aparece en su propia lista de argumentos entre la palabra reservada *func* y el nombre del método.

# Métodos (continuacion)

De hecho, puedes definir un método para cualquier tipo que definas en tu paquete, no solamente estructuras.

No puedes definir un método de un tipo de otro paquete o de un tipo básico.

# Métodos con punteros a receptores

Los métodos pueden asociarse con un tipo o un puntero a un tipo declarado.

Acabamos de ver dos métodos *Abs*. Uno para el puntero a un vértice "Vertex" y otro para el tipo MyFloat.

Hay dos razones para usar un receptor de tipo puntero:

- Primero, para evitar copiar el valor en cada llamada al método (más eficiente si el tipo usado es una estructura grande). 

- Segundo, de tal forma que el método pueda modificar el valor a la que apunta el receptor.

Intent cambiar las declaraciones de los métodos *Abs* y *Scale* para usar Vertex como el receptor, en lugar de Vertex(con asterisco).

El método *Scale* no tiene efecto cuando v es un vértice "Vertex". Scale cambia *v*. Cuando "v" es un tipo (no un puntero) el método ve una copia del vértice "Vertex" y no puede mutar el valor original.

*Abs* funciona de cualquier forma. Solo lee v. No importa si está leyendo el valor original (a través del puntero) o una copia del valor.

# Interfaces

Una interfaz es un tipo de datos definido como un conjunto de métodos.

Un tipo interface puede contener cualquier tipo que implemente esos métodos.

# Las interfaces se implementan implícitamente

Un tipo implementa una interfaz simplemente implementando los métodos.

*No hay declaración explícita*

Las interfaces implícitas desacoplan la implementación de paquetes entre los paquetes que definen las interfaces: Ninguno depende de otro.

También favorece la definición de interfaces precisas, porque no tienes que encontrar todas las implementaciones y etiquetarlas con el nuevo nombre de la interfaz.

El paquete io define Reader y Writer; tú no tienes que hacerlo.

Es decir, los paquetes internos de Go que utilicemos para crear interfaces no necesitamos declararlos.

# Errores

Un error es cualquier cosa que se defina a sí mismo como una cadena de error. La idea está en la interfaz predefinida *error*, con su único método, *Error*, que devuelve una cadena de caracteres:

```
type error interface{
    Error() string
}
```

Las funciones de impresión del paquete *fmt* sabe cómo llamar al método cuando se le pide imprimir un error.
