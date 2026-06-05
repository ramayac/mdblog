---
title: Aprendiendo a usar Canvas en HTML5
date: 2010-04-07
author: Robertux
tags: canvas, css, html5, firefox, mozilla, html, javascript, jquery
draft: false
post_id: blog-3515952828243908885.post-4533753731997979926
---

HTML5, la nueva versión del lenguaje de marcado para crear páginas web está ya soportado por los últimos navegadores. Esta nueva versión trae muchas novedades a la web y características especiales que antes solo era posible llevarlas a cabo usando Flash. Entre los tags más mencionados son el tag Vídeo, para incrustar y reproducir vídeos nativamente sin necesidad de Flash y el tag Canvas que permite crear y animar gráficos 2D y hasta 3D de forma nativa.

En este artículo demostraremos el uso del tag Canvas, realizando dibujos y un mini paint utilizando CSS y Javascript.

Antes de empezar, cabe mencionar que por el momento no todos los navegadores soportan todas las características de HTML5. Por ejemplo, Internet Explorer es de los que aún no soportan el tag Canvas y por lo tanto no podrán apreciar el demo de canvas que hemos publicado en este post. A continuación una tabla comparativa mostrando las características de HTML5 que soporta cada navegador a la fecha:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEigzB50TCpEpbxTo2L-5QtNqvfczlanSAaImZ6mwulrXXQGNGUqtdcRKlluLBtity6ZxNvZxXNhGueUD-Y_5z7WUitujNv1g8MjLJSyUPo3UimJsEDV10Ku6M8YMuB1gZ4Xl68LISZ1eYA/s400/html5-table.png)  
"Tabla comparativa mostrando las características de HTML5 soportadas por cada navegador"

Primeramente, debemos crear nuestra página HTML con su respectivo body dentro del cual debemos ubicar nuestro tag Canvas. Este tag no requiere atributos especiales mas que su height y width. Entre la apertura y cierre del tag Canvas puedes incluir cualquier otro contenido HTML el cual le será mostrado al usuario si este utiliza un navegador que no soporte el tag. Esto puede servir como un espacio para mostrar un mensaje indicando la falta de soporte por falta del navegador. También es útil asignarle un ID a dicho tag, para luego poder identificarlo al momento de manipularlo desde Javascript:

"Código HTML mostrando el tag Canvas y algunos de sus atributos"

Una vez que tienes el tag, es hora de manipularlo desde Javascript! empecemos por obtener el Context, el cual es el objeto que posee los métodos para dibujar dentro del canvas. Entre estos métodos se encuentran fillFect(), arc(), moveTo(), stroke(), etc. Este objeto Context además tiene una propiedad llamada fillStyle con la cual puedes definirle el color en formato rgb o rgba con el cual se realizarán los trazos.

"Código Javascript obteniendo el objeto Context asociado al objeto DOM Canvas para luego establecer un color vía RGB y pintar un rectángulo"

A continuación mostramos dos ejemplos un poco más elaborados. En el primero, mostramos cómo dibujar una cara feliz y en el segundo cómo asociar los eventos del ratón para pintar sobre el canvas por medio del click. Para ello utilizamos la aplicación web [jsFiddle](https://jsfiddle.net/) que te permite incrustar HTML, CSS y Javascript al mismo tiempo que se puede apreciar el resultado de dicho código. Además se hace uso de jQuery para simplificar el acceso al objeto DOM Canvas:

"Código HTML, CSS y Javascript para dibujar una cara feliz en el Canvas. Clic en el tab Result para ver el resultado"

"Código HTML, CSS y Javascript para crear una superficie de dibujo en la cual se pinta un cuadro negro por cada clic del ratón. Clic en el tab Result para ver e interactuar con el resultado"

[En este link](https://mugtug.com/sketchpad/) pueden probar Sketchpad, un ejemplo mucho más elaborado de una aplicación web simulando las características de paint.

También en el [Mozilla Developer Center](https://developer.mozilla.org/en/HTML/Canvas) pueden encontrar más información y tutoriales sobre el tag canvas y el resto de características HTML5 actualmente soportadas por Mozilla Firefox.

En mi opinión personal, HTML5 permitirá sustituir Flash en un futuro no muy lejano por lo menos en las características que antes solían verse en páginas web como animaciones, reproducción de medios y videojuegos en línea. Es importante aprender desde ya lo que se puede hacer con él y así estar preparados para un futuro en el que este estándar sea utilizado por la mayoría de sitios y la crezca demanda de programadores con dichos conocimientos.