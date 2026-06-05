---
title: Haciendo Limpieza de Fin de Año
date: 2008-12-29
author: Robertux
tags: desorden, administrar, desktop, herramienta, escritorio, buscar, ubuntu, debian, año nuevo
draft: false
post_id: blog-3515952828243908885.post-8676424220067513930
---

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgP15xLb2lraOW3UA9o0yIvrM5UKNzVVtREq5LBbh6UB8LInOAAdmCJxY2HpDNeWt1O9SIidWyXIKUp_wQWMj6Oaw6GY2aSe8409cEInI6dbmDYIaoZP-LlWk5nc6gIOOr-P_rKR4g2_CI/s400/iconsmess.jpg)    
"Desorden de iconos en un escritorio Mac"

La mayoría de personas somos desordenados por costumbre, especialmente si nos referimos a la información que almacenamos en nuestras computadoras. Por defecto, guardamos todo en el escritorio, luego vamos creando carpetas temporales, luego quizá carpetas por fecha o sino, vamos usando números correlativos. De cualquier forma, esto no ayuda a mantener un orden en nuestros archivos y programas que viven dentro de nuestra PC.

Muchos tienen por costumbre ordenar sus cuartos o sus escritorios, hacer una limpieza general de sus casas, pintar o arreglar los imperfectos aquí y allá. Otros hasta toman como superstición el hecho de tener la casa limpia y ordenada para el final del año, creyendo que de esa manera, el orden y la limpieza les durará todo el año siguiente.

Aún sabiendo que [puede existir cierta eficiencia en el desorden](https://www.srbyte.com/2007/03/es-la-gente-desordenada-ms-productiva.html), muchos otros optamos por hacer una limpieza interna en nuestras computadoras. No solamente quitarle el polvo, ordenar las cajas de CDs apilados por ahí y quitar las notas y papeles que ya no ocupamos sino ir mas allá y ordenar-limpiar el software y los archivos instalados en ella.

Para Windows existen varios programas que ayudan a esta tarea, como por ejemplo el [Perfect Uninstaller](https://www.freedownloadmanager.org/downloads/Perfect_Uninstaller_55181_p/). Linux también cuenta con su set de herramientas para administrar lo que guardas por ahí, darte cuenta del espacio que ocupa y si realmente está siendo útil o si solamente fue parte de un programa que instalaste una vez y luego quedó ahí abandonado y a su suerte, por error.

Entre estos programas tenemos:

Baobab: También conocido como Disk Usage Analyzer. Es un programa que te permite hacer un escaneo de toda la jerarquía de carpetas que compone tu computadora, clasificándola en base al tamaño que ocupa dentro del disco para que puedas saber qué archivos son los más grandes y si no son tan necesarios, borrarlos de una buena vez. La interfaz es muy sencilla e intuitiva, como pueden ver. Este programa viene instalado por defecto en Ubuntu o Debian, o especialmente si utilizas el entorno Gnome. Sino, puedes instalarlo mediante el comando apt-get install baobab como root.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhvvQAO5HMeI2GmYeG6nOvmqgxo7BMVxLVYPF2EOiVjp529_8WVACXc2Q7Wift7Vp02av9418jkmjpd4LADPBEAKeN336ZbFFezFXJZLTNfG1cgEWVB78I69ffcvCZewyaFwlv3V36wKH4/s400/baobab.png)    
"Captura de Pantalla de Baobab, mostrando de forma tabular y gráfica el espacio que ocupan las carpetas y subcarpetas del sistema"

Cruft: Es un programa para la consola que escanea en busca de archivos innecesarios que alguna vez fueron instalados junto con alguna aplicación pero esta no existe mas. Cabe mencionar que aún se encuentra en versión pre-alfa, pero según las pruebas realizadas de manera personal, demostró ser bastante satisfactorio y efectivo. Para instalarlo, basta ejecutar apt-get install cruft. Luego, nada más ejecutas el comando escribiendo cruft también como root, o sino, cruft -h para ver el resto de opciones disponibles.

Apt: El mismo gestor que, en Debian y Ubuntu te permiten instalar paquetes, también te permite remover los paquetes innecesarios. Los programadores de Linux, a diferencia de los de windows, hacen uso de librerías creadas por otros para la mayoría de programas, de manera que se crean dependencias, se reusa código/funcionalidad y hace que los programas ocupen menos espacio en disco. El sistema APT, además, utiliza archivos empaquetados (.deb) como instaladores para cada programa que utilizas, el cual guarda en un caché por si deseas reinstalarlo, de manera que no tengas que volver a descargarlo.

Para borrar los paquetes de instalación (.deb) que se han descargado y quedan en el caché, puedes ejecutar el comando apt-get clean o apt-get autoclean como root. Para borrar librerías o dependencias que ya no están siendo utilizadas por algún programa, ejecutas el comando apt-get autoremove también como root.

Si tienen dudas o fallos con los comandos antes mencionados, déjanos un comentario.