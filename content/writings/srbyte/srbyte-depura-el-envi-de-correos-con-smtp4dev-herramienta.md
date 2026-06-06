---
title: Depura el envió de correos con "smtp4dev" [HERRAMIENTA]
date: 2013-02-07
author: Rodrigo A.
tags: smtp, herramienta, mail, smtp4dev, correo, programacion, programador, codeplex
draft: false
post_id: blog-3515952828243908885.post-8291044391070901596
---

Casi siempre que se hacen aplicaciones web, surge la necesidad de notificar al usuario por correo sobre algún suceso, ya sea registro, suscripción u otro. Así que como developer, siempre nos vemos en la necesitad de probar código para enviar correo. La practica nos enseña que no hay que esperar hasta último momento (ya en producción) para probar el envió de correos, pero usualmente probar estas funcionalidades es un dolor de cabeza, especialmente en Windows (¿quien ya ha configurado un servidor smtp en Windows?).

Pues para aliviar esta carga, les quiero compartir esta sencilla herramienta, que francamente me salvo la vida muchísimas veces para probar envíos de correo, se llama "smtp4dev".

Realmente no hay mucho que ver o explicar, smtp4dev es un sencillo programa que se aloja en el taskbar y escucha el puerto 25, por cualquier correo entrante.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEioieZ-7pO_KYChvWHPcm4tmJHvoCbvqQJyko6H1FWDiJLkeaFGQjOtzjNOSx3Mx4G1qcXPld64S9AFtUIszg2hla3ITUmYwbzaH1iZKciMdhh3Wtldl31dINsj1HSfH2vDZoiypeh6VPFz/s320/mainform.png)    
"Bandeja de entrada de smtp4dev"

Entonces si usas Java Mail o la funcion mail() de PHP para enviar un correo, smtp4dev lo recibe y te despliega una notificación de recibido, el correo "recibido" lo pueden abrir con Thunderbird o Outlook sin mayor problema.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhP3BGJEIoJgLQfwwcgi6WGtcO4vwqwjGvTW0Knrb8r_EKDpHsLHw7r0VxQKNp5D0chGO1XG5YZLVesYLA6sTJP_XwTmouVWdxGKoUasIbrcJ39asVoz5VRYD8YM0Cv5OmIeuu5f2JiaqDu/s1600/notify.png)    
"Notificación de correo entrante"

Simple, funcional y open source ¿qué más se puede pedir?, pueden descargar smtp4dev acá:

> [https://smtp4dev.codeplex.com/](https://smtp4dev.codeplex.com/)

Espero que les sirva, tanto como me ha servido (y lo sigue haciendo!)

Saludos!