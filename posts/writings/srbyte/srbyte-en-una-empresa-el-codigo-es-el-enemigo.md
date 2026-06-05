---
title: En una empresa: el codigo es el enemigo.
date: 2009-01-05
author: Rodrigo A.
tags: codigo, desarrollo, programacion, programador, empresa, opinion
draft: false
post_id: blog-3515952828243908885.post-7890007720728715217
---

Voy a escribir un par de artículos a la semana, hace poco conseguí trabajo y me esta consumiendo, pero las experiencias y habilidades son invaluables.

El código en exceso es malo. Requiere mantenimiento periódico. Posee errores ... que deben ser encontrados, depurados y mitigados. Añadir funcionalidad extra, implica que el código antiguo se tiene que adaptar.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEg-qy5yiZc67TxH2IhCzcywS-JneJfLORIp3Hd3KM1ux4t_dAMRBQxzvKXgcJ0OoQmeJNkXua-EqGG9DltmWsaSh5x-LTmd0KzXXC7IU3nFuN0yktAKxlQev0d1_-JSzUvZPclA6_7h2ymD/s320/iStock_000000237891XSmall_3.jpg)    
Mientras hay mas código escrito:

- Existen mas lugares en donde se esconden[los infames bugs (que se combaten depurando o escribiendo más código)](https://www.srbyte.com/2008/12/herramientas-del-programador-unit.html)
- El proceso de compilación tarda mas tiempo
- Los Checkouts del [repositorio de control de versiones](https://www.srbyte.com/2008/03/programemos-mejor-subversion.html) también se alargan
- Un empleado nuevo tardara más tiempo en comprender el sistema
- Es mas complicado refactorizar código
- Y también es mas difícil encontrar documentación precisa de: que hace qué, cuándo, dónde y por qué.
![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhXbXL94H9sN5ZSck0BT4AJe1KA_tRPtYgnVwNZQCTzqe59rD_lmBDVQVlh06nHJ6yw9jxIxKhlDh5uRsaKSCgmEo1JGqVbeKhcHryZeVlkvmv6bDFdkIVG1KKsI2ttiIU7ooypul0rYb4Z/s320/beautiful-code.png)    
"Este libro debería leerse en la Universidad..."

Más código, significa menos flexibilidad y funcionalidad; esto para muchos es una paradoja, pero la mayoría de veces, una solución simple, rápida y elegante, es mejor que una "mega super función general que hace de todo".

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjqVPyyXLVZIf10Lf_dd4UDLcjiUEsXkVusMiqtrQZsofAz6LGxiyqpRs5UHD0j0XQ49wvt0MYHqdY_vO0EBiPwD5LwmQPOQ10WkABlDtkZrS3lfCA_0wBSGZyY2Kkmm7jtVQXWQGvp9uDq/s320/crappy-code-book-cover.jpg)    
"Este es más común..."

El código es escrito por ingenieros de sistemas, técnicos programadores o consultores; digamos simplemente que son desarrolladores. Producir más código, requiere mas desarrolladores. Varios desarrolladores deben comunicarse. Un desarrollador tiene un costo de "canal de comunicación" de n^2, como pueden ver, las comunicaciones incrementan exponencialmente con cada desarrollador en el proyecto ( 1^2=2, 2^2=4, 3^2=9, etc...). Añadamos un poco de papeleo (burocracia) a los canales de comunicación:

- Control de actividades de desarrollo, por metas y diarias.
- Documentos que constan la finalización de las actividades programadas.
... si hago esto por cada desarrollador... termino con un enorme costo organización, en función del tiempo gastado en cada desarrollador, o en función de un nuevo puesto de trabajo para que alguien realice esta tarea.

¿No se debería hacer todo lo posible para[incrementar la productividad del individuo](https://www.srbyte.com/2007/02/programando-mejor-parte-iii.html) en términos del buen código que este escribe? La idea es simple: Escribir menos código para hacer algo (y si se puede, que sea buen código). Si tiene que escribir menos código, se contratan menos personas, y se reducen los costos de comunicación.

Después de todo, una buena empresa debe ser eficiente y eficaz, ¿no?. Si algo ya existe, úselo. Si hay mejores tecnologías, procure utilizarlas. Y si sus desarrolladores le dicen que hay que utilizar una nueva tecnología, procure prestar atención a lo que dicen, y si es factible hágalo.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiySMM8NO1eBY4Xscvgim3D_UGJNEvwhE2ofhMu0EeMxmXu3LqGqnmx_e48d1AgDQhE6MtRr2TZMKskcg4i6mGQV-9uv9nYkzDFCo_xyAUGEqQhQkexUyEVCMlvK2mbnhAY-GOLrIJJ2XHx/s320/reclining.jpg)    
"Buen equipo, buena silla, buen escritorio... son necesarios para producir buen código."

La comodidad de un desarrollador de software no puede ser discutida. Estos merecen buenas sillas, estar cómodos en sus cubículos o escritorios, café cerca y como máximo, ocho horas de trabajo diarias. El trabajo de un desarrollador de software, no solo es uno de los mas estresantes, sino también es uno de los mejor remunerados. Y si no trata bien a sus desarrolladores, estos producen código enmarañado, descuidado e irresponsable, hackeado para que funcione (ley del llegue)... malo en pocas palabras. Codificar mal, siempre es producir mas código del que se necesita.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhvu0Rp0weI2rVAhUiB4AUs8Uazf_wrDx4S-v6QlOX97Rrlk131uoAEq0YrJUHQX_0W5OtOPtIDvS_4JpAsGCe3oF8fVtBKGkXeBNXRynXI7Us-np7Ef3DkDE67Vkl-qS3l6OOcu3jxNBXT/s320/no_hacking.gif)    
Repitan conmigo:
> "desarrolladores infelices, producen mal código, que aumenta mis costos".
En la empresa en la que estoy trabajando, la mayoría de desarrolladores piensan que los "usuarios" son los enemigos. Para mi, el código es el enemigo, y para la empresa también.

Una cosa más, para los desarrolladores que leen este articulo... sigan como consejo la sabiduría innegable de xkcd: ![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgKFM3eZ2b_5XNfgtxKO3OSw4nNAt7UzcjEKi8Ti6QOkJ_P1tKNU-kSeDlCUKt6FLFjB9dHIh0dCjMAlTeSN53_JMPbBV3X4jb4CrEcUT34_R03tJxxZ3u0RPxs8wiHb38Uyv_EcWuYSw06/s320/goto.png)