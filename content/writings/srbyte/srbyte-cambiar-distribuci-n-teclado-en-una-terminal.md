---
title: Cambiar distribución teclado en una Terminal
date: 2010-05-10
author: Rodrigo A.
tags: linux, teclado
draft: false
post_id: blog-3515952828243908885.post-6010140104916114549
---

Mini nota:

Justo hace una media hora, necesitaba iniciar una aplicación en Linux, con el mapa de teclado en Ingles. Como la aplicación únicamente funciona en modo de pantalla completa, no podía cambiarle la distribución del teclado usando el "Indicador de Teclado" de Gnome ([del que pueden obtener más información aquí](https://library.gnome.org/users/user-guide/2.27/prefs-keyboard.html.en)). ¿La solución? muy sencilla, iniciar una terminal y digitar:

> #setxkbmap us
La instrucción anterior cambia la distribución del teclado en la terminal, y ahora se puede iniciar el programa con el mapa de teclado especificado. Seguramente hay muchas otras formas de hacerlo, pero si llevas prisa, o la aplicación con la que estas trabajando "no se deja", puedes utilizar este sencillo comando.