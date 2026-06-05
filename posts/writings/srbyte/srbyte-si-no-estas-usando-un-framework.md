---
title: Si no estas usando un framework...
date: 2009-03-27
author: Rodrigo A.
tags: desarrollo, practicas, framework, software
draft: false
post_id: blog-3515952828243908885.post-6638478040019983651
---

Si no estas usando un framework para desarrollar software, probablemente estés re-inventando la rueda. La única validez que tiene el proceso de "re-inventar la rueda" es para conocer como funciona. Pero si te están pidiendo elaborar un sistema realmente amplio para una empresa, con una cobertura del 70% de las operaciones, y esa empresa NO se dedica a realizar software... entonces es una seria estupidez desarrollar software sin un framework.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjjYozeSIREWJ_HYS0a21DvLYk3KMiyeF2jEQJVGUM6SZ9apTBICRmsSPukcYOTerM0inAns6-Kij2G047994Ti-fGdBeMrRpZZGOLuqpAJt-DYMqd00DN8tH7bUb7Qeik7UbcV4ySGkcuD/s320/4359967kid-on-bike-square-wheel.jpg)    
"Cuidad, si re-inventas la rueda, podrías terminar así..."

Un framework, es una estructura de soporte definida mediante la cual, se desarrolla y organiza lógicamente una pieza de software. Entonces, el programa, modulo o sistema en cuestión, se apoya en un framework (en el marco de trabajo) para ser desarrollado. Un framework incluirá soporte para generar programas, una serie de librerías organizadas en módulos lógicos, y un lenguaje (que se ejecuta en una maquina virtual usualmente) para ayudar a desarrollar y unificar los módulos que conformen el proyecto. El uso de un framework, es tan obligatorio en un proyecto de grandes magnitudes, como lo son las típicas practicas de desarrollo de software:

- [Unit Testing](https://www.srbyte.com/2008/12/herramientas-del-programador-unit.html)
- Code Coverage
- [MVC](https://www.srbyte.com/2008/10/cosas-que-todo-desarrollador-web-debe.html)
- [Version Control](https://www.srbyte.com/2008/03/programemos-mejor-subversion.html)
- Automated Deployment
Para los estudiantes, o novatos en el tema, permitanme dar un sencillo ejemplo: Imaginemos por un breve momento, que se quiere construir (de acuerdo a un plano) una casa de dos plantas, para tres personas, con cochera y jardín. Este es el problema:

> construir la casa de acuerdo a las especificaciones (del plano), construirla en un tiempo planificado y con la mayor eficiencia posible.
Bien, hay un par de caminos que podemos tomar, el primero consiste en: 1. Comprar ladrillos, cemento, arena, pintura, tejas, piso cerámico, defensas, hierro, contratar mano de obra calificada, comenzar la construcción....

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEh_9HcfBHXQW0CHwRFmM9Z0qZ8m3uHVDGwU5ZpNbIPa5hJpxjzM5qTyvIOCTa91CE72gVpOsVsvE85u_ORBFpQWmeHjPXDEhBGpAOykdruKbJ98-xbZB9fIgxW3JJn15duD7anquseSEup4/s320/custom_home.jpg)    
"Niiiiceee...."

Al menos, eso es lo lógico, ¿no? Son los pasos "normales", lo "sano". La misma lógica es aplicable al mundo del software.... sin embargo, existen personas que rechazan esta "linea de pensamiento" y cometen errores que en el contexto del ejemplo anterior, seria:

2. Construir la casa, desde CERO (literalmente), de manera que vamos a inventar nuestro propio cemento, pintura, techo, defensas, pisos. Vamos a usar herramientas que no estamos seguros si las vamos a aprovechar y vamos a gastar en cosas completamente innecesarias... como... flamingos rosados de mármol tallados a mano importados de la India para el jardín (wtf?).

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhKoGXIfc84eCLi5RYVbdv8HT8arfFzPjC8uFhlobgOc9zPr0xYfqzp44vEX5eF10lP2Jn9_rqQsm_IWTEg1VnZ4GWio1pgD4e5B8L3fdWq7OGF1t1Tb22jYdzux6Zp-WF941giN-lFULqj/s320/crazy_house_2.jpg)    
"NOT so nice. NOT AT ALL."

Y para colmo, creen que se puede realizar la casa, en el mismo tiempo que lo haría de la manera "normal". Se que para muchos lectores, esto suena ridículo, ¡¿verdad?! suena muy ridículo. Es más, seria increíble pensar, que alguien proponga soluciones así.... pero les tengo una noticia...

> Este es el caso que sucede SUCEDE en las empresas, y también SUCEDE en los equipos de desarrollo.
Casi a diario, se esta perpetrando este crimen. No solo es una perdida de tiempo, sino que también es un insulto al desarrollador de software que si tiene buenas practicas ([o que intenta tenerlas](https://www.srbyte.com/2008/07/consejos-practicos-de-desarrollo-de.html)). ¿Saben que es lo peor? muchos cometen el error de creer que con solo usar Java o .NET (y solamente eso), ya están usando un framework suficiente para hacer un RIA ([Rich Internet Application](https://en.wikipedia.org/wiki/Rich_Internet_application)) en 3 meses. Pues déjenme decirle, que si ese fuera el caso, no existiría Spring, Struts, IceFaces, MyFaces, ASP.NET MVC, Adobe Flex, RoR, Etc... Lo terrible es que este suceso se sigue perpetrando, y es por simple ignorancia...

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEieb6haLUASLULgHaBPsTbDw0ptJoxwC5WODdt2i0f0Jp9KRbml5uLjKFTyACrNPkpruH8SWUmPreHHT4qywUqTJe9ky2BPfxncMwep92k1LmyqCtUkhA3eloZFNQihJy-bWGkEIMBxS0_T/s320/simio-que-ignoras.jpg)    
"¡Que pena ser ignorante!"

Empresas, por que rayos ¿no capacitan a sus gerentes de informática?, ¿a sus arquitectos de software?, ¿a sus team leaders?, ¿Por que no escuchan las voces de los programadores, que de cara a semejantes atrocidades, alzan la voz inmediatamente y proponen el cambio?
[Ya lo dije antes](https://www.srbyte.com/2009/02/acerca-del-micromanagement.html), pero no me molesta volverlo a repetir...
> "si un equipo de desarrollo, sumergido hasta el cuello de dificultades técnicas, pretende producir software: simplemente vamos a obtener software que es REFLEJO de las dificultades, problemas, ignorancia, malos requerimientos y mala administración en los que se vio envuelto el mismo".
Instruyace, no deje de leer, no se estanque, procure estar a la par (o solo un centímetro atrás) de la tecnología, y procure, sobre todas las cosas, escuchar a sus colegas, pero a los que saben.

Pero sobre todas las cosas, por favor, si no estas usando un framework.... ya es hora de que comience a usar uno.

¿Y tu, qué frameworks usas?