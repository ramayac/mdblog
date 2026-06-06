---
title: Internacionalización en tus Sitios Web con jQuery
date: 2010-11-09
author: Robertux
tags: internacionalizacion, i18n, jquery
draft: false
post_id: blog-3515952828243908885.post-8230686576696307559
---

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhLhGb1_i96QmtpHMmCToIJVJXaU2-LLTbXAwaewEj8DJKQMwmv3hK98zO2kZHBhUHC0TEmpCvgOWJsrzz_1kGA3qa5cg9drNkaTeNwo1PZuIbiinG4XwwIO-hksOnrnHlpCstywtgCLy0/s400/i18n.png)    
"Si esperas recibir usuarios de países con otros idiomas, es una buena recomendación poder mostrar el contenido de tu sitio en diferentes idiomas, tal como lo hace last.fm"

Internacionalización o i18n es la característica de los sistemas que les permite adaptar su contenido a diferentes idiomas según las regiones donde se desee mostrar. Muchas veces no solamente te permiten traducir el idioma del contenido sino también monedas, fechas y ajustes de zonas horarias. Dependiendo del lenguaje de programación usado, pueda que este ya brinde soluciones para poder permitir la internacionalización de tu sitio web. Java EE por ejemplo ofrece los Message Bundles combinados con tags JSP para llevar a cabo este fin.

Así como es posible hacerlo desde server side, en esta ocasión les traigo un ejemplo realizado desde client side, usando el muy popular framework de javascript conocido como [jQuery](https://jquery.com/), al cual le he agregado un plugin denominado [jQuery.i18n.properties](https://codingwithcoffee.com/?p=272) el cual me permite implementar la internacionalización en un sitio web de una forma muy similar a como se realiza en java.

La ventaja de usar una solucion client side? me permite cambiar el idioma del contenido del sitio sin necesidad de salirme (logout) y volver a entrar(login), ni siquiera es necesario refrescar la pagina! todo se realiza en tiempo real a la velocidad de un Ajax Callback.

Para demostrarles lo mencionado anteriormente he desarrollado un ejemplo utilizando jquery, jquery.i18n.properties, un pequeño script de javascript y un par de servlets de java. Todo esto dentro de un proyecto web en eclipse y además no pueden faltar los Message Bundles. Los Message Bundles son archivos de texto con extensión .properties que contienen propiedades de tipo clave=valor en los cuales se almacenan los mensajes que se mostrarán en pantalla de acuerdo al idioma seleccionado. De tal manera que existe un archivo llamado locale_es.properties para los mensajes en español, otro llamado locale_en.properties para los mensajes en inglés y otro más denominado locale_fr.properties para los mensajes en francés.

[https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhLhGb1_i96QmtpHMmCToIJVJXaU2-LLTbXAwaewEj8DJKQMwmv3hK98zO2kZHBhUHC0TEmpCvgOWJsrzz_1kGA3qa5cg9drNkaTeNwo1PZuIbiinG4XwwIO-hksOnrnHlpCstywtgCLy0/s1600/i18n.png](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhLhGb1_i96QmtpHMmCToIJVJXaU2-LLTbXAwaewEj8DJKQMwmv3hK98zO2kZHBhUHC0TEmpCvgOWJsrzz_1kGA3qa5cg9drNkaTeNwo1PZuIbiinG4XwwIO-hksOnrnHlpCstywtgCLy0/s1600/i18n.png)![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhdC4JgsLMElLdlMOCgizavAvHLWJYCtp2clok2aouoXroCS2k3zTkbke8eeswEh_iPZNZv2QDDShRuLRW8lZOK7gV85Vl8DjCujNmVoxonuUUmPThI5UtAXMC0mP2PlJVdvwvejAXp4H0/s400/Screenshot-5.png)    
"En el panel izquierdo pueden ver los archivos utilizados en este proyecto. Servlets, javascripts, css, html y Message Bundles. En el panel principal a la derecha pueden ver parte del código HTML escrito."

Se ha escrito una pequeña página de login. Para internacionalizar los elementos les he colocado la clase "i18n" y como id el nombre de la clave que en el Message Bundle contiene el mensaje a mostrar. Se han agregado tres links en la parte superior que invocan a la función javascript para cambiar de idioma. Esto llama al plugin i18n.properties quien mediante un ajax callback obtiene el Message Bundle según el idioma deseado y lo carga en un mapa con los datos en forma de clave=valor.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEikG1F_6rUE4oKLPK4kBR34zrnaVxynY7qMbDC0vntrPKdRjH5UsvYfPv0P9txVTibxkVNoeGUVsUaYCpUTD52rmBs2LHq7mCq_YROliy85rSRC27TAvAzHZwqoVZ3ZgUIC85qhsRbt7rQ/s400/Screenshot-4.png)  
"Las funciones principales del archivo core.js con las cuales se lleva a cabo la internacionalización"

Luego de tener el mapa, se asignan los valores a cada elemento con la clase "i18n" y la clave según el ID del elemento. Luego de ello se escribió un Fake login test que devuelve respuestas random de manera que se prueben tanto mensajes de error como mensajes de éxito, ambos internacionalizados.

A continuación unas capturas de pantalla del sitio web en funcionamiento:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgetsNKBKDaWXL56e6wCh1WrKszCkq8shdZTzqUYCK2dgSf1Pv_rfUHqEVINHSDh1DBKlt5rSLP422GCHI2mlofcecfzQYDxj2p-p1eUWINZCO7DAm3bDkD1l1V-sne2679-heK6jj1iLc/s400/Screenshot-3.png)    
"Sitio web mostrando mensaje de error en el login, todo en idioma inglés"

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhy1ZvMugD-R_-W45nE1IUQq8NXDnoNzzmJi674jCcSvSideq1gk2IXIhR7ynyUTslhuGNPaAxFwNC6TFmEZ73UzXaKyqEBrlLiXaC2FJOLXXMz4j2XEso-EGCkVKnYHMeEbwP66WsE0qY/s400/Screenshot-2.png)    
"Sitio web mostrando mensaje de error en el login, todo en idioma español"

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgHIAHy149rTh1kbpUsLMIk_8j_KtRaRN0Okz3Cc-5_epFHU79BA-A_Rd8OP0nFoSHg3fPPKxVmsKKBtLWU8cJQUCF7gHvRWfAJY46BehDU2tG8FGPsKOV0Ds9F_uXyuPa2KguKVciSFA8/s400/Screenshot-1.png)    
"Sitio web mostrando mensaje de éxito en el login, todo en francés"

En [este link](https://dl.dropbox.com/u/3393841/InternationalLoginEclipseProject.zip) pueden descargar el código fuente del proyecto.