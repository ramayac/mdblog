---
title: ¿Por que aprender Python?
date: 2009-01-23
author: Rodrigo A.
tags: programacion, python
draft: false
post_id: blog-3515952828243908885.post-7008671535066090185
---

Si eres un Power User, Sys Admin, o simplemente realizas toneladas de trabajo frente a la computadora, probablemente te hayas topado con que hay ocasiones en las que se necesita automatizar ciertas tareas tediosas. Por ejemplo, el de buscar una cadena de texto repetida en varios documentos.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgsKiPBhslyaPNZjJmp5rX-bK0_cciV28KidNlkzLTGc_yxBrXe5b4YMSzHFLh8RwZHwVtQC69pY0rnYtrJ1aAbjvtkNCDRKcZcX-BmUFWP4rx2lgzbZQ0Gq2P6Kb51JAg040SXeNmG1vcZ/s320/300px-PowerUser.jpg)    
"Yo soy un PowerUser, por eso uso Linux!!!"

En el trabajo, en uno de los proyectos en los que estoy asignado, me vi en la necesidad de realizar la tarea mencionada; buscar una cadena de texto (una función) en cada archivo de un modulo, y obtener el nombre del archivo en donde sucede la ocurrencia. Usualmente esto lo podría hacer en Linux, con una sencilla combinación de los comandos: "grep" y "ls" en un script. Pero claro, no estoy usando Linux en mi trabajo, así que me vi "forzado" a utilizar una (basca de) función de búsqueda en el IDE que estamos usando. Como no obtuve los resultados adecuados, en eso recorde una herramienta multiplataforma, opensource, y mas robusta que el chistoso "command prompt" de Windows...

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgVmKgldh1TU8nqsfX2zDhlrYRXU6Gt1-oZiS98p6skp0jEMVqhnkgQn6D9lafkbZjnahCW6ZFQ-xnWhWQd3WjU5nQtbIkVFLIZT1-mh-tEFq8wW7TYex3NPPxyA2DqMGAF6IjwjaFENEnV/s320/python3.jpg)    

- Imaginen que necesitan buscar todas las veces que aparece el nombre de tu exnovi@ en las conversaciones que se guardan del MSN, ¿como harias esto?
- Y si quieres mover todos los archivos de fotos que tienes regados por toda tu carpeta de documentos, ¿como lo haces?
- Y para eliminar todos los espacios vacíos de los nombres de tus archivos de música, o los caracteres extraños, ¿que herramienta usarías?
- Si bien, (yo se que) existen herramientas para realizar estas tareas, a veces ni la mas completa GUI (Interfaz de Usuario) puede asistirnos con tareas masivas de movimiento de datos, respaldo, eliminación de archivos vacíos (archivos de 0 KB), o de nombres extraños ( ###$%__song.mp3 por ejemplo).
- Y que te parece el caso en que tal vez necesitas una base de datos pequeña y personalizada, realizar un simple juego o realizar una GUI especial.
- Puede ocurrir que eres un developer de C/C++/Java, y el ciclo usual de codificar/compilar/depurar/recompilar te parece muy lento; ¿que sucede si deseas hacer un programa que se conecte a una base de datos remota y ejecute pruebas automatizadas?

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgKAKM1e5yGPzATfeZQEQmr6QU7ouIvqDsN8WlXGxRRS9XZipEUrg3OPB6ev6SAiaNpfeMnWZVaxZnVU3XSHwDEpHfi77sSjPI75awNNCl6YbgZt-Rq2RhBJ7hv6s9eSkxRVzEeeGIXGyUs/s320/big-robot.jpg)    
"Si te sientes como un robot realizando tareas repetitivas, tal vez sea tiempo de considerar una herramienta que automatice tus tareas!"

- ¿O una sencilla aplicación que pruebe rangos de IP que tengan conexión a Internet, y que cambie tu IP cada cierto tiempo, para poder navegar "tranquilo" en el trabajo? (más información sobre esto luego).

En cualquiera de estos casos y en muchos otros más, Python es el lenguaje de programación para ti.

¿Realmente crees que es necesario tener una aplicación diferente para cada una de estas tareas? ¿Necesitas una herramienta diferente para cada problema, o es mejor tener una compacta y elegante Navaja Suiza?

Python es fácil de usar, y disponible (como la mayoría del Software Libre) para Windows, Linux y Mac OS X, y probablemente te ayudara a realizar cada uno de estos trabajos más velozmente.

Si bien es fácil,también es un verdadero y muy robusto lenguaje de programación, que ofrece mucha mas estructura (entiendase legibilidad) y soporte para programas largos, más de lo que cualquier lenguaje de bash script te permitiría. Python también ofrece mucho mas chequeo de errores de los que soporta C, y siendo un lenguaje de alto nivel, tiene tipos de datos bastante útiles, como arreglos flexibles y diccionarios.

Y como Python tiene tipos de datos más generales, es aplicable a un dominio de problemas mucho más amplio que Awk o Perl, y aun así, muchas cosas son tan fáciles de implementar en Python como en esos lenguajes.

Python permite dividir tu programa en módulos que pueden ser reutilizados. Posee una larga colección de módulos estándares (I/O, llamadas al sistema, sockets, PyGame, etc).

Como es un lenguaje interpretado, no necesitas compilarlo y linkearlo (como C/C++). Y se compila "al vuelo" y/o bytecode como Java o .NET

El interprete de Python se puede utilizar interactivamente, lo que facilita experimentar con las características del lenguaje, hacer programas "para una sola vez", o para utilizarlo como una avanzada calculadora científica, je je.

Python permite escribir programas de manera compacta y legible. Y casi siempre, un programa escrito en Python, sera mas corto que su equivalente en C, C++ o Java. Es extensible, se pueden añadir nuevos módulos al runtime de Python con solo saber como programar en C.

Pero lo realmente importante, y por lo que vale la pena mencionar a Python en este momento, es que hace poco salio la nueva versión de Python, la 3.0 Y lo radical de esta versión, es que es TOTALMENTE incompatible con Python 2.6 y menores. Muchos se sorprenderán de semejante movida, pero el creador de Python: [Guido Van Rossum](https://www.python.org/~guido/), tiene buenas razones para haberla hecho, entre ellas: mejorar la sintaxis del lenguaje, cambiar un buen par de tipos de datos, y proveer una plataforma para convertir a Python, en el lenguaje interpretado de alto nivel de mayor popularidad en el mercado. Si estabas esperando "el momento adecuado" para aprender un lenguaje de programación, sin lugar a dudas este es el momento para aprender Pyhton (3.0)!!!

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgkvO1do_4xLorF72Vm9_ysFoTDStwUv0V5Igy1uj47uXcFXlBKeuERbwkKDeAAIpMyX40owiDVxjNC7aSfa_3XDmO6L1gQzkthNEe_ulNQ0MpLCYlUbWqiFWp9-VH0xFQqQXpoDh5Psuaf/s320/studying-boh.jpg)    
"¿Cansado de lenguajes estáticos y compilados?, quizas sea hora de cambiar a un lenguaje de ALTO nivel."

Ya sea que estés comenzando, o simplemente estés cansado de Perl, de Awk, o de los ridículos archivos batch de Windows, etc.

Así que... a bajar Python:
[
> https://www.python.org/download/releases/3.0/
](https://www.python.org/download/releases/3.0/) Y la documentación de Python 3.0 (Py3K) puedes encontrarla aquí:
[
> https://docs.python.org/3.0/
](https://docs.python.org/3.0/) Saludos!