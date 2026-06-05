---
title: Programación Ágil con Eclipse
date: 2010-09-30
author: Rodrigo A.
tags: java, galileo, helios, programacion, eclipse
draft: false
post_id: blog-3515952828243908885.post-3669810229883386281
---

Muchas veces hacemos uso de un IDE tan poderoso como eclipse y no aprovechamos todas sus ventajas porque quizás no sabíamos que este las tenía. Desde la versión Galileo y ahora con la nueva versión Helios, el [entorno de desarrollo Eclipse](https://www.eclipse.org/) ha agregado muchas mejoras a su IDE para hacernos la vida más fácil a la hora de programar en Java, especialmente cuando tienes que realizar grandes y repetitivas porciones de código. A continuación describimos algunas de ellas:

Source > Generate Getters and Setters: Es una opción que se encuentra desde versiones anteriores de eclipse, quizá desde la versión Europa pero muchos quizá aun no la conocían. Al momento de elaborar sus clases java pueden escribir únicamente sus miembros privados y desde esta opción pueden generar automáticamente los getters y setters para estos miembros.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEicZu-Azz4w-u_LgxvHUqD3DORFjSC808zuW-uV5tcNEyC4M-DsoSd2jq2Er58qjsX8zzuah69Xx0lAZ9yu-u7wSxWnI62W5gG-jY0Wg1UvaNob5R_EPx83bneRRHRMJvHPUEdnIaRq-sk/s400/generateGetters.png)    
"Opción Generate getters y setters en Eclipse"

Source > Generate Constructor Using Fields: Similar a la opción anterior. Una vez tengan todos sus miembros privados de la clase y desean crear un constructor que reciba parámetros para inicializarlos, basta con acceder a esta opción del menú y este se crea automáticamente.

Source > Generate toString: Esta opción es nueva para mi y está disponible desde la versión Galileo de Eclipse. Muchas veces para efectos de logging queremos imprimir en pantalla todos los valores que poseen los miembros de una clase pero si esta clase tiene demasiados miembros es muy difícil escribir un método toString() para mostrarlos todos. Este método nos automatiza esta tarea.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhsCoixAH0QZz3c1nmguvSK1TnHy29NGToiVagcM5bzl2xzIngOtS5poJoFZm_KMQErzRQSeyo9SroJJdT7wIur4ETg0BpKAeov45mvbgVXCIjuaAU1DIawPFBHXL5YDIGyFCr9lENkxko/s400/toString.png)    
"Source > Generate toString() genera un método toString() mostrando los valores de los campos de la clase"

Autocompletado: Siglas en lugar de nombres de clases: Muchas veces tenemos que escribir nombres de clases realmente largos y al escribir las primeras tres letras el autocompletado (Ctrl + espacio) nos da demasiadas opciones como para encontrar la que buscamos. Recientemente me enteré que puedo digitar las siglas de una clase Java y Eclipse te muestra las clases cuyas letras formen esas siglas. Por ejemplo, en lugar de digitar EntityManagerFactory, puedes escribir nada más EMF:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjVJxHRnphAk2tW9A3xRbS_brcy09QKyK7lBNjSxKNTkgfa74vfFrwPX5gWfEgQQgxNrBKeAHN_GHVXOt6OWUPK67UPoiiEdn1u2S2zQo86YIg9eH5PVhYE_6LcS_RyzSWWmSlDzAD8-rQ/s400/autocomplete.png)    
"Autocompletando el nombre de la clase EntityManagerFactory con sus siglas EMF"

Si tu autocompletado es demasiado lento y tienes muchas clases cargadas en tu classpath que nunca utilizarás, puedes filtrarlas para que tu autocompletado cargue más rápido. En las preferencias (Window > Preferences) selecciona Java > Appearance > Type filters y filtra los paquetes de clases que no deseas que te aparezcan en el autocompletado.

Export > Runnable Jar File: Ahora Eclipse nos permite exportar nuestros archivos Jar empaquetados (o enlazados, según configuración) con todas las dependencias necesarias disponibles para la ejecución del Jar a generar.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjGBZ_1disjaZJ1t2O-x673JMGmqM9fjeeO624PScQYHPt67WH8c0TQGBBLau_nbEtJX1AwOi7m1VFfudJUES1FwOg-IxsNUqneHqAvh0eVW9UN_MT4ppjTHRgK7zt-YSIL5CMDIHpZcfI/s400/exportRunnableJar.png)    
"Opciones para exportar el Runnable Jar"

Source > Surround With: Otra tarea mecánica es quizá agregar for, while, if o try/catch. Eclipse puede ayudarnos a encerrar sentencias de código en cláusulas for, while, if o try/catch mediante la opción Source > Surround With. Si se trata de un try/catch, el IDE generalmente nos remarca las líneas de código que puedan disparar una excepción y en un tooltip nos mostrará un vínculo directo a la opción Surround With Try/catch.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEg1k63OBaOxkCZfPp9ewR5FRL9IqQgtDylw6fngMjtmC8A9Sp76XbCpgxhyO7dFNrmHx-HAVnFpEZzwLawnWuy0tSyFziOJ4UmCpNLzgDO098FFlQSiiOPyU0xNb9VS9uv3nxqcfrp764Y/s400/surroundWith.png)    
"Cuando una sentencia puede disparar una excepción, el IDE nos puede encerrar automáticamente esta sentencia en un bloque try/catch"

Atajos de Teclado: Eclipse es aun mas fácil de usar si te memorizas sus atajos de teclado. En [este vínculo](https://eclipse-tools.sourceforge.net/Keyboard_shortcuts_(3.0).pdf) puedes descargar un PDF con los shortcuts más comunes para Eclipse desde la versión 3.0 en adelante. Algunos comunes son como por ejemplo:

- Ctrl + T para búsqueda de una clase Java dentro de tu proyecto y buildpath
- Ctrl + O para buscar un miembro dentro de una clase
- Alt + Shift + X para ejecutar tu proyecto (con las variaciones J para proyecto java, T para unit test, R para ejecutar en servidor de aplicaciones, etc)
- Alt + Shift + Z (teniendo seleccionadas las líneas a encerrar) para mostrar las opciones de Surround With
- Ctrl + Shift + C para comentar una o varias líneas
- F3 para ir a la declaración de una clase o un miembro de una clase

Además de las características propias del IDE, Eclipse tiene la ventaja de poseer una variedad de plugins para extender su funcionamiento. Un ejemplo de ello es el plugin para [Grepcode](https://grepcode.com/), un buscador online para examinar código fuente de clases Java.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEj2zIWXs-oQanH_6iRK26Qm8j-obQoEX3N9tBl12UW00tnCfVvlrakUWnpH4j4sbxbespQgwOFQwYhzpOazayIYGD8bpN_viUEn8D5HR53w6WebiKuAk6Zolo26TD2G61ZQR7jdzGhIQkk/s400/grepcode.png)    
"Grepcode mostrando el código fuente de la clase EntityManager"