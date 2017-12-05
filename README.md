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


