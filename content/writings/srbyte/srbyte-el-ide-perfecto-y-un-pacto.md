---
title: El IDE perfecto... y un pacto.
date: 2007-01-16
author: Rodrigo A.
tags: linux, novell, microsoft, gnu
draft: false
post_id: blog-3515952828243908885.post-5501699388329849125
---

Como ya había mencionado, la migración a herramientas GNU/Linux puede ser una tarea un poco abrumadora y extenúan algunas veces para el que viene del mundo del software propietario, no es para "asustar" a nadie, pero es realidad. Pues como hace poco también cambie mi mentalidad por tecnologías multiplataforma y por el desarrollo de software libre con herramientas "OpenSource"; me vi en la necesidad buscar un IDE (porque a eso estoy acostumbrado) para mi sistema libre. Y es que durante la travesía, me percate porque es más "fácil" desarrollar en Windows. Microsoft, con su afamada serie Visual Studio, provee una herramienta "unificada" para el desarrollo sencillo y (algunas veces) practico de las ventanas. Un API sencilla, un lenguaje sencillo como VB (que ha hecho más por el auge del mundo de la programación que los lenguajes de OO), un IDE que se apega al deseo natural de todo programador de convertirse en un diseñador de interfaces... aun cuando ocurran crímenes como este:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEj06EeIq4Uy_dKLFef7T6M3MSx4ajQqJe_J5LirzARc3TBvSM0aIowKLU_Od2AJTuvmviKs99UB42hsmV3N3WfGAmqT8yoL0_9GnZfoGpaoUsdvZlc2wg01MfxhZTTSrbU2dgRGb52-mnc/s400/wgetgui-screenshot.png)    
Librerías "simples", motores de bases de datos (aunque terriblemente lentos) fáciles de usar, depuradores incluidos y soporte para plugins. Aun así, en este sistema propietario... no tengo muchas opciones como desarrollador. Se me enseña que es lo que debo de usar, a desarrollar con la terrible mentalidad de licencias propietarias que no solo hacen daño a una persona, sino a una economía como la de nuestro pulgarcito de América.

En un sistema libre, hay tanto que usar que no se sabe donde comenzar. Pero se tienen varias opciones según el lenguaje que prefiera, listo algunos a continuación:

> C/C++: Anjuta, KDevelop
> VisualBasic: Gambas
> FreePascal: Lazarus
> Mono (.Net): MonoDevelop
Pero si bien, estos entornos solucionan algunos problemas de adaptación, no son solución alguna para el desarrollo de aplicaciones multiplataforma bajo un lema de:

> "compile once, run everywhere..."
Verán, no me gusta enredarme entre el (¿tedioso? o ¿problemático?, probablemente por MIS practicas de programación) proceso de re-compilar una solución informatica completa en un sistema y en otro. Evidentemente si utilizo herramientas que de por si son multiplataforma como: OpenAL, OpenGL, SDL, MySQL, SQLite, PostgreSQL, GTK, QT, etc...no tendría muchos problemas. Pero aun así tendría que cambiar el lema anterior por:

> "run once, compile everywhere...many times"
Definitivamente una idea no tan atractiva. Ya que mi deseo es producir, indiferente del SO (sea libre o propietario), y que mi producto pueda ser usado en cualquier SO. Bajo esta idea de producir y distribuir sin dificultades, simplemente puedo usar:

> Java: Eclipse, NetBeans, XDevelop, JDevelop
> Python: Boa Constructor, Eric, Stani's Python Editor
Como IDE para aplicaciones que seguramente podre compilar y luego correr en cualquier SO (y que funcionen de acuerdo a mis necesidades, después de todo los mismos IDE funcionan en Windows o GNU/Linux). Mono es prometedor, pero existen ciertas diferencias entre el y .Net, esperemos que el "pacto" entre Novell y Microsoft lleve a nuevos rumbos la virtualización y la interoperabilidad... aunque lo mas probable sea que Micro introduzca su modelo de negocio tan nocivo para los países subdesarrollados como este. Tiempos interesantes deparan al software libre.