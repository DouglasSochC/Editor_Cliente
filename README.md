
# Manual de Usuario

Se ha creado un traductor que englobe dos de los principales lenguajes utilizados hoy en dia, los cuales son Javascript y Python, esto es a partir del lenguaje Java, el cual es el lenguaje base.

La idea principal del programa es que se genere un análisis léxico y sintáctico de cada uno de los distintos archivos de entrada los cuales están definidos en el lenguaje Java. Y como salida se espera, lo siguiente:

- Lenguaje Traducido
- Tockens Generados
- Posibles Errores Léxicos
- Posibles Errores Sintácticos
- Grafo del Análisis Sintáctico

Por lo tanto se tiene lo siguiente:

#### Pantalla inicial:
Cuando se ingresa a la aplicación se mostrara la siguiente pantalla, el cual posteriormente se describirá cada uno de sus componentes y como estos funcionan.

![enter image description here](http://drive.google.com/uc?export=view&id=1dVywFX6G6ZvcbsUqH_jc2235IEf9L1wz)

- Agregar Pestaña:

	Como su nombre lo indica, este nos permite generar nuevas pestañas en el cual se puede ingresar nuevo código, hay que tomar en cuenta que cada pestaña tiene un comportamiento independiente.
- Guardar Código:

	Su funcionamiento no es mas que solo almacenar el código que se encuentra expuesto en el área de texto, que tiene por nombre "Archivo de Entrada", este se vera posteriormente.
- Traducir a JS:

	Esta opción realiza varias funciones al momento de analizar un archivo de entrada, los cuales son:
	- El lenguaje traducido
	- Retorno de tockens generados
	- Muestra los posibles errores léxicos
	- Muestra los posibles errores sintácticos

	Hay que tomar en cuenta, que el lenguaje traducido, y los tockens generados se pueden descargar en un archivo .pdf, y los posibles errores que se pueden encontrar, estos se desplegaran en la consola.  

	También hay que tomar en cuenta que para poder empezar a interactuar con el programa, se debe de generar una pestaña, nueva, ya que un caso contrario este puede poseer un error.

- Generar Grafo:

	Como su nombre lo indica, este genera un grafo de análisis sintáctico, del lenguaje java, y lo desplegara en la misma que se analizo, solo que en la parte baja de la consola.

- Traducir a Python: 

El funcionamiento que este realiza no es mas que solo la comunicación con el servidor, que debería de traducir la entrada del texto de tipo java a el lenguaje de python.

#### Estructura al Utilizar las Opciones:

![Uso de las Opciones](http://drive.google.com/uc?export=view&id=1nPM3fd8RuUebfTu-2rtNrGW8Yoamf3PC)

- Archivo de Entrada:

	Esta caja de texto nos permite ingresar el código Java para que posteriormente pueda ser analizo, a través de esta área de texto, es como se puede utilizar el programa por completo.
- Consola:

	Esta caja de texto desplegara todos los errores que se pueden hallar durante el análisis del lenguaje Java.

Al verificar la imagen se puede visualizar como debe de ser el ingreso de los archivos de entrada, al programa, y como a pesar de que cada pestaña tiene el mismo nombre, al moverse por cada uno de estos se tiene comportamientos completamente diferentes
