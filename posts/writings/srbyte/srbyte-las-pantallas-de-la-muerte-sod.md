---
title: Las Pantallas de la Muerte (SOD)
date: 2008-08-08
author: Rodrigo A.
tags: apple, linux, windows, ipod, iphone, mac
draft: false
post_id: blog-3515952828243908885.post-3617676976153617986
---

Tenia rato de no tener una "Pantalla de la Muerte" asechando en mi monitor, y eso me inspiro y motivo a realizar esta entrada. Muchas personas conocen este fenómeno como la famosa "Pantalla Azul", mas, el nombre completo es: Blue Screen of Death (Pantalla Azul de la Muerte, BSOD por sus siglas en ingles). Este terrible fenómeno, sucede cuando el sistema operativo (tipicamente Microsoft Windows, aunque nadie se salva realmente) tiene un error que requiere que todas las operaciones se detengan inmediatamente, por eso mismo también se le dice "Stop Error". Usualmente sucede cuando el error en el sistema es tan fatal, que se requiere que se apague la computadora para prevenir cualquier tipo de daño en el hardware, así que no solo sirven para causar daños psicológicos irreparables, sino que también cuidan tu PC.

En la familia de Windows NT, estos errores suceden cuando los controladores de los dispositivos están mal codificados/programados o cuando el hardware funciona incorrectamente.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhiLenrfq-jweir6Cq7CrXoGzus0kC5tn6XUAFTfbWX8eVegAASpCy-NFr320Go_0bKTRN7zECQmpVe6iSMMv43O1od8hUxFNgTglymQJQIFz3w58v6L9AUD-QTK5BDd2qB72RpxVYX2B8/s400/Windows_XP_BSOD.png)    
"BSOD de Windows NT"

En la familia de Windows 9x (95, 98, 98SE, Me, etc...) eran los DLLs el problema, o bugs en el Kernel:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhBOqqtxkmUcOoGMve1P3-aS0e-mM6T7yY_K2mb2GA6usfcit74_aMiQjX6PHSAzlL49V399tq5GD3RkUMw-PfHxV5pBgLannnx3kRfroSftXUyWBGveZsvWr4-_ke4lrU_m-Vq0fbqY8o/s400/Windows_9X_BSOD.png)    
"BSOD de Windows 9x"

Inclusive puede suceder, que si se golpea -fuerte- el CPU, la maquina se reinicia y muestra la BSOD.

Por cierto los "amigos" de Microsoft, han incluido una característica en Windows que permite que la pantalla azul aparezca manualmente. Para habilitarla, con fines "educativos" claro, hay que añadir un valor al registro de Windows (usando regedit.exe) y seguir las instrucciones de [este sitio](https://pcsupport.about.com/od/tipstricks/ht/makebsodxp.htm). Luego, después de reiniciar la maquina, una BSOD aparecerá cuando el usuario presione la tecla SCROLL LOCK (PAUSA INTER) dos veces mientras se sostiene la tecla derecha CTRL. Como ya dije, esta característica es para fines útiles y educativos, sirve para obtener "memory dumps" de la maquina en un estado especifico. También puedes provocar un BSOD al terminar el proceso csrss.exe y winlogon.exe . Aunque si son curiosos, ya habrán notado que con el Task Manager no se puede terminar esos procesos, pero si se puede con software de terceros ;) Y por si fuera poco, también hay un salva pantallas de la famosa pantalla azul, y la pueden descargar [aqui](https://technet.microsoft.com/en-us/sysinternals/bb897558.aspx).

Bien, si dejamos el mundo de Winbugs Windows, veremos que hay muchos dispositivos también tienen esta "característica". La Xbox 360, cuando esta consola experimenta una "fallo general de hardware", tres luces rojas aparecen alrededor del botón de encendido.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgfT-CU8a_4owihjk3yuqdtrXHMEXNqKtj004TD8Si-D2Qw6FD-VPZTt4PcG99aNG9enEowsk_erumeSvRTMtWCduJJXH-pa_v4CepS9OpXK7N1QX3eWTVAfnKo-R-9mbHMaQCiJB74p_w/s400/Xbox360-ringofdeath.jpg)    
"Xbox 360 anillo de la muerte"

El nintendo DS también tiene una pantalla de la muerte, que sucede cuando se remueve una tarjeta DS o un cartucho de GBA mientras se esta en el menú de selección de Pictochat. El color de la pantalla esta basado en la versión del Firmware del DS:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhrsV1B_Vu6i7UX4zmVQSIwwrm-G_jh-DSLcrTzty_OyABxOAd-AIQG2EPva1AO_hdkyxNuJEspWuKgY7oyvBRVmr5raqEvVooVCdgHuxJvv9YQsIPC7-lyUqvg01hj90ppR6nm33Yo42Q/s400/DS_BSoD.JPG)    
"Pantalla de la muerte del nintendo DS"

Parece que entre la gente que ha comprado el nuevo iPhone, esta apareciendo un nuevo rumor: sobre algun problema del software 2.0 y 2.0.1 del iPhone o de alguna aplicación de maliciosa (para el iPhone) que daña el sistema operativo del mismo, solo que este error... es permanente :-S (hay usuarios que reportan que la característica de restauración no funciona).

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhhwKH5q6oD88vGLwnnJWW1908R12b9yTiJ_JscU0hdTbsakgBcCk9EAuh5_3GU83fZlQluEWm4nM8cjcGiMZGWeHL3vDdV0jHz7iYq2Hq0Vuga5LxGVMDjWP2vb4_OMbhYxZ0oTr1WEJ8/s400/iphone-boot-screen.jpg)    
"Pantalla de la Muerte para el iPhone"

Los productos de Apple tampoco se salvan, antes la mayoría de Mac tenían esta imagen:

![image](https://1.bp.blogspot.com/_ayvorITawE4/SJu6ZxBsVjI/AAAAAAAABAM/jNl5fhL8XrI/s400/Sad_mac.png)    
"Pantalla de la Muerte, Sad Mac"

Los primeros iPod, también tienen una especie de "Sad Mac":

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEh-JMQ5SY6z4SMfX_pHuA7Q79FV1mTWBGEtEX7bjrvUid2OPjJ7XL_2-5QTgwmTOhmKkjwjMilVvAMxMppeyTYlqx_fuPXJFo5fx__-2mQmY0r8I3yW-Id47W2Rz7PGpM6RZxaH3xDmjQE/s400/Sad_iPod.png)    
"Sad iPod"

Y desde OS X, las Mac tiene una "elegante" pantalla de la muerte:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEghDOhc5xhQacSXdXSxiGH84coW58CqyaJOQWrSrLnx1OtGeyz-CGptsRd-lAneUW_l75jEjm_SdL2BSJC1iBN4czzXdpLZTjJQDTojra0epwUuiEY0EmOAFue6yzKNpg1EMsa6Mt9-C38/s400/MacOSX_kernel_panic.png)    
"Pantalla de la Muerte de Mac OS X"

También GNU/Linux tiene el temido "Kernel Panic", que en lo personal, solo me ha sucedido dos veces: una por un CD de instalación rayado, y otra por que yo compile mal el kernel. Pero también tiene pantalla de la muerte:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEioCWcM0588kqJnDjmgiBqH7DNMe_6auz2cupateRalqwP_kTF98d2evljYh5VRFeD5Br2-xtO_ojG5irmVOFeskZoaDTVjT9l3LL3X9bb4CHVQuIA1IY1JvQqSfVc18eA9iy6Jjv6in-I/s400/Kernel-panic.jpg)    
"Kernel Panic!"

El problema de un SOD para el usuario, es que sus mensajes de error son algo inexplicables/cripticos/misteriosos y no ayudan en nada a los usuarios a conocer el "porque" del error sucedido, lo que contribuye al pánico y la ansiedad (jaja).

Invito a los lectores a recordar las pantallas de la muerte a lo largo de su vida, ¿En que sistema operativo lo han tenido? ¿Cuantas veces? ¿Que otros aparatos han visto con pantallas de la muerte: cajeros automáticos, reproductores de DVD, el Wii, o en el cable?