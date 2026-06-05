---
title: Los 25 Errores de Software Más Peligrosos del 2011
date: 2011-07-11
author: Robertux
tags: programacion, seguridad, sql
draft: false
post_id: blog-3515952828243908885.post-5692024609438567604
---

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhqUqbzXBfjeL79-3mRb7cyRF5IlNDXc8w8fP1K9tuVQLh3X_mbyrNUaiSV2RryeY_rRkeBszYUFRfYM8y-o0qDoXZr738fJiNnWwtN5iJ7acjKdiZgdzZmTWSgwBSwM60Uy2gNErdM6nI/s400/sql_injection.gif)    
"SQL Injection. Una de las más peligrosas y comunes vulnerabilidades en los sistemas"

Recientemente me entero de una comunidad online de programadores y especialistas en seguridad que cada año se encargan de enumerar las vulnerabilidades de sistemas informáticos que más daño han realizado y que más se han popularizado entre los sistemas para que cada uno tome las medidas necesarias al momento de desarrollar sus sistemas especialmente cuando manejamos datos que al ser vulnerados causarían pérdidas significativas de información en nuestra empresa.

Esta comunidad, denominada [Common Weakness Enumeration](https://cwe.mitre.org/about/index.html), liberó recientemente la [lista de los 25 errores de software](https://cwe.mitre.org/top25/index.html) que más se han utilizado y que más daño han causado durante el presente año de manera que podamos tomarlos en cuenta para evitar que nuestros sistemas sean atacados.

Además de la lista, en el sitio puedes encontrar información detallada del error, las variaciones que pudieran existir del mismo, código fuente de ejemplo de utilización del ataque y posibles alternativas de corrección a las vulnerabilidades que este aprovecha. Es una fuente bastante completa de ítems a tomar en cuenta al momento de diseñar la seguridad de nuestro sitio. De hecho, si en su equipo de desarrollo o calidad no poseen un checklist de vulnerabilidades a cubrir, les recomiendo usar este sitio como referencia.

Me remito a describir en español las tres primeras vulnerabilidades de la lista:

1. La inadecuada neutralización de comandos especiales utilizados en SQL (A.K.A. SQL Injection). La número uno de la lista. Tal como lo ilustra la imagen del post, si utilizamos SQL simple para validar nuestras páginas de login, algo como por ejemplo:

"SELECT * FROM SYSTEM.USERS WHERE NAME = '" + inputBox1.text + "' AND PASSWORD = '" + inputBox2.text + "'";

cualquiera podría escribir cualquier comando SQL dentro de los textboxes y con suficiente astucia obtener acceso total al sistema saltándose la validación del login o peor aún, obteniendo acceso total a la base de datos.

Esto me recuerda a una vieja ilustración que XKCD elaboró para el día de las madres:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEh-G-eFU9xTYHe_DClGECVq7zjQDAkDxwG5zL7YNW88AclmmDgfQSpgFYucavniIes3wjlf8hvjLB0M6TxN3der7oi8UQ9ULA6qDt5rsVjIXNEoim03Z65uMvYrqrLyC7pwmdoIL8HUxMU/s400/exploits_of_a_mom.png)    
"...'Little Bobby Tables' xD"

En el blog de [El Rincon de Cisko](https://ciskosv.blogspot.com/) pueden encontrar más información en español acerca de este tema.

2. La inadecuada neutralización de comandos especiales utilizados en un sistema operativo (A.K.A. OS Command Injection). Aplica el mismo principio que el SQL Injection pero en este caso nos referimos a comandos del sistema operativo. Pueda ser menos común que nuestros programas ejecuten comandos directamente sobre el sistema operativo pero aunque sea menos común, esta vulnerabilidad es mucho más peligrosa que el SQL Injection ya que en el sistema operativo pueden residir no solamente el servidor de aplicaciones, servidor web y servidor de bases de datos sino muchos otros servicios sin mencionar que una vez que el atacante ha logrado ingresar a un servidor, es mucho más fácil que este logre acceso a cualquier otro servidor dentro de la intranet, especialmente si nuestro servidor de aplicaciones posee privilegios de administrador (root) lo cual es mencionado en otro elemento de la lista: "11. Ejecución de aplicaciones/servicios con privilegios mayores a los necesarios".

3. Copia de búffers sin la comprobación del tamaño del dato de entrada (Buffer Overflow). Este es nuevo para mí. Según la documentación del sitio leo que aplica mayormente para lenguajes de programación que no poseen un manejo automático de la memoria utilizada, como por ejemplo C y C++ ya que estos no revisan previamente si la porción de memoria donde se están escribiendo los datos está reservada para su fin o si su tamaño excede lo esperado y termina escribiéndose información en la memoria reservada del sistema operativo, por ejemplo.

La recomenadación para evitar esta vulnerabilidad consiste en verificar previamente el tamaño de los búffers de entrada, utilizar librerías o frameworks que realizan esta comprobación por nosotros o utilizar lenguajes de alto nivel que poseen un manejo automático de la memoria a utilizar por aplicación o grupo de aplicaciones.

Les recomiendo nuevamente y con gran énfasis leer el resto de la [lista de vulnerabilidades](https://cwe.mitre.org/top25/index.html) y sus respectivos detalles técnicos, causas y soluciones para que como programadores y administradores de sistemas podamos dormir bien por las noches sin preocuparnos porque el operador nos va a llamar en la madrugada avisándonos que vulneraron un server y robaron información confidencial y valiosa de nuestra empresa. Mas aún teniendo en cuenta los recientes ataques de LulzSec, AntiSec y Anonymus.

vía [Java Code Geeks](https://www.javacodegeeks.com/2011/07/top-25-most-dangerous-software-errors.html).