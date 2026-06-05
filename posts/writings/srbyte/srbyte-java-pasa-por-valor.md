---
title: Java "pasa por valor"
date: 2010-02-23
author: Rodrigo A.
tags: pasar por valor, java, pasar por referencia
draft: false
post_id: blog-3515952828243908885.post-3677642860504595223
---

Nota: Aclaro que no soy un experto, y que me encuentro en constante aprendizaje, es por eso mismo que escribo esta breve nota que es más que todo un recordatorio personal, y que quizás la sirva a más de algún programador para aclarar el concepto.

Hace unos dias me puse a depurar un problema, que tenia que ver con pasar una cadena como parámetro en un método (primer error, porque las cadenas en Java son inmutables, cosa de la que hablare más adelante). En el metodo se modificaba el contenido de la cadena, e incorrectamente supuse que esa modificación se mantendría al terminar el método, por aquella frase (que muchos equivocadamente citan), que dice que en Java:

> "Los tipos de datos primitivos se pasan por valor, y los objetos por referencia".
Vaya sorpresa me lleve, al darme cuenta que esta afirmacion era incorrecta. Los terminos "pasar por valor" y "pasar por referencia", tienen significados muy precisos, y usualmente se abusa de sus definiciones, o en mi caso se entienden o asumen equivocadamente.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjj05zL-mCqSmaDC1ICXjW1yw0cLiwd4akH7y7xCr6pNnr-i5vDctz63j0DR4Qe7LsvSPVAaQGlA6anvlTcitWLlnTMinHRq2gaRSpnbvULjdQVrEeXyVPvPlmPKioCP3TbnaOSwTkS773o/s320/confusionpirates.jpg)    
"Pass by what?"

Buscando aclarar mi duda, encontré en el sitio [Javadude.com](https://javadude.com/articles/passbyvalue.htm), una excelente explicación con ejemplos prácticos que demuestran claramente que Java "pasa por referencia", es una lectura casi obligatoria que recomiendo mucho. Cito las definiciones de los términos en cuestión, encontradas en el sitio que mencione anteriormente con traducciones propias (cualquier sugerencia para mejorar la legibilidad del texto es bienvenida) por motivos aclaratorios:

Pasar por Valor (Pass-by-value): "El parámetro actual (o expresión) que se pasa como argumento al método se evalúa completamente, y el valor resultante de esta evaluación se copia en la ubicación de memoria que tenia el parámetro anterior a la ejecución del método. Esa ubicación es típicamente una porción de memoria en la pila de ejecucion de la aplicación (que es como Java lo maneja), pero otros lenguajes pueden elegir una forma de almacenamiento de parámetros diferente."

Texto Original: "The actual parameter (or argument expression) is fully evaluated and the resulting value is copied into a location being used to hold the formal parameter's value during method/function execution. That location is typically a chunk of memory on the runtime stack for the application (which is how Java handles it), but other languages could choose parameter storage differently."

Pasar por referencia (Pass-by-reference) "El parámetro (en el método) actúa meramente como un alias para el parámetro actual (que esta afuera del método). En cualquier momento que el método utiliza el parámetro (para leerlo o escribir en el), se esta utilizando en realidad el parámetro actual (el que esta afuera del método). Texto original: "The formal parameter merely acts as an alias for the actual parameter. Anytime the method/function uses the formal parameter (for reading or writing), it is actually using the actual parameter." Para puntualizar, Java es estrictamente "Pasar por valor", como C (no como C++, que si soporta pasar por valor y por referencia, vean esta [comparacion en la semantica entra Java y C++](https://en.wikipedia.org/wiki/Comparison_of_Java_and_C%2B%2B#Semantics). Y esto lo podemos encontrar en la Java Language Specification (JSL) en: [https://java.sun.com/docs/books/jls/third_edition/html/classes.html#8.4.1](https://java.sun.com/docs/books/jls/third_edition/html/classes.html#8.4.1): > "Cuando un método o constructor se invoca, los valores del argumento o expresión actual inicializan nuevas variables como parámetro, cada una del tipo de dato declarado, antes de la ejecucion del cuerpo del método o constructor. El identificador que aparece en el "DeclaratorId" puede ser usado como un simple nombre en el cuerpo del metodo o constructor para referirse al parámetro formal." > Texto original: "When the method or constructor is invoked, the values of the actual argument expressions initialize newly created parameter variables, each of the declared Type, before execution of the body of the method or constructor. The Identifier that appears in the DeclaratorId may be used as a simple name in the body of the method or constructor to refer to the formal parameter." En pocas palabras: > Java tiene apuntadores y ademas es estrictamente "Pasa por Valor". Tener la idea de que Java pasa datos por referencia, es más común de lo que imagine, me tome la libertad de preguntar sobre esta caracteristica a varios amigos programadores con más experiencia que yo (5) y solo uno tenia claro el concepto.... ![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjBXl2Blbz5NlGfwsl9Z518xVBoDvLSmhdo_HK053LawTA026uH_Sz6zGjxzkNYPTe073WTKiqsAHdrSa_YIbOPOvDur0mCHhabQDtG8nCA2htcrbMfgIm0uTV0YDaxa3ApN4xXB_mMrMeG/s200/javaevil.png)     "... de pronto senti que Java era maligno."

Y no es que ellos no sepan como funciona Java, porque a la hora de programar, o de aclarar la idea, era fácil ver como algunos terminaban asintiendo y cambiando de parecer. Es simplemente una característica que se olvida, y que creo que es peligroso que eso suceda (que se olvide), por lo que me tomo la libertad de escribir un poco sobre ella, esperando que a más de alguno le sirva, o que si ya lo sabían que simplemente lo recuerden.

No se olviden de leer el articulo de [Javadude.com](https://javadude.com/articles/passbyvalue.htm) es realmente muy, pero muy bueno, especialmente si quieren ver ejemplos aclaratorios.

Saludos!