---
title: Adaptando tu Sitio Para Móviles - Parte 2
date: 2010-02-08
author: Robertux
tags: navegador, framework, movil, ui, web, programacion, javascript
draft: false
post_id: blog-3515952828243908885.post-6122207666345663019
---

Esta es la segunda parte del post anterior de [Adaptando tu Sitio Para Moviles - Parte 1](https://www.srbyte.com/2010/01/adaptando-tu-sitio-para-moviles-parte-1.html), en el cual hablaba de los servicios que tea automatizan este proceso. Ahora explicaremos algunas librerías utiles para crear nuestros sitios móviles por nuestra cuenta.

Para crear una versión de tu sitio orientada a un dispositivo móvil debes considerar ciertas características de la navegación en móviles:

Ancho de banda: Entre estas, el hecho de que estos dispositivos se conectan por WiFi o 3g y el ancho de banda disponible puede ser mucho menor al de una computadora personal, por lo cual tendrás que evitar cargar las páginas con muchas imágenes y mucho contenido.

Contenido no soportado: Debes considerar también que no todos los dispositivos móviles soportan contenido rico como Flash (ehem, iPhones, iPods e iPads) así como tampoco no todos soportan Java, Silverlight, Active X o Adobe Air, entre otros. Hay muchos sitios basados en Flash que ni siquiera con la versión Lite, incluida en algunos móviles, pueden renderizar las animaciones de tu sitio. Especialmente porque, a diferencia de las imagenes y el contenido HTML habitual, las películas de Flash no se ajustan automáticamente al tamaño de la pantalla de tu dispositivo. Los que tienden a utilizar Flash para darle una gran vistosidad a su sitio como por ejemplo restaurantes, hoteles o páginas de artistas podrían considerar realizar estas mismas animaciones usando librerías Javascript como jQuery, Motools o Dojo y aun así no es recomendable colmar la página de tanto adorno y animación.

Procesamiento: El procesador de los dispositivos también te impedirá ejecutar mucho Javascript en tus páginas, es mejor dejar la mayor parte del procesamiento de información en el servidor web. Recuerda que en la actualidad los smartphones más avanzados poseen un procesador de no mas de 1Gz y la verdad es que en mi opinión, no necesitan mas de eso ya que tu smartphone no es una herramienta de trabajo que tendrá que realizar muchos cálculos en tu sitio. Su uso primordial es consumir/producir información, no procesarla.

Dimensiones de la pantalla: Las dimensiones de las pantallas en los dispositivos móviles, si bien hoy en día abarcan toda la superficie de los mismos, es mucho más reducida que la de un monitor de una PC(480x320px en un iPhone, 800x480 en un Nokia N900, 360x480 en un Blackberry Storm y 800x480 en un Google Nexus One). Esto te obligará a distribuir el contenido de tu sitio en un espacio más reducido quizá dividiéndolo en diferentes páginas o eliminándolo si no es tan necesario. En un blog, por ejemplo, ya no se deben incluir los típicos widgets contadores de visitas, shoutboxes, secciones de últimos comentarios, etc. Si se desea hacerlo, podria utilizarse una pagina aparte que sirva como "About".

Interacción con el usuario: Un usuario comunmente interactua con un Smartphone haciendo uso de sus dedos o un Stylus, haciendo "taps" en lugar de clics para abrir vínculos y haciendo "swipe" para mover las scrollbars y visualizar más contenido. Hay ciertos eventos que no se podrán producir desde un Smartphone, por ejemplo el drag&drop (basado en el evento mouseover), combinaciones de teclas con Ctrl-Shift-Alt, scrolling, etc. En [esta tabla](https://www.quirksmode.org/m/table.html) pueden encontrar una completa referencia sobre los eventos, tags HTML y atributos CSS soportados por la mayoria de navegadores web para móviles.

Librerías

En mi opinión personal, el dispositivo que más ha popularizado su sistema operativo, el cual además revolucionó las interfaces gráficas para dispositivos móviles es hasta la fecha el iPhone. Si deseas que tu aplicación web tenga la apariencia y animaciones similares a las de este dispositivo puedes hacer uso de la libreria [jQTouch](https://www.jqtouch.com/), la cual es en realidad un plugin para el popular framework de javascript [jQuery](https://jquery.com/) que simula la apariencia y animaciones propias de las interfaces de un iPhone OS.

[Acá](https://www.jqtouch.com/preview/demos/main/#home) puedes ver un demo de las caracteristicas de jQTouch (visualizarlo desde un smartphone) y a continuación puedes ver un vídeo de las features principales de dicha libreria:

"Features generales de jQTouch"

Segun comentan en su sitio, la [YUI Library](https://developer.yahoo.com/yui/grids/#mobile) también funciona perfectamente en lo que ellos denominan browsers Grado A.

Otra buena libreria de javascript especializada en el desarrollo de sitios web para móviles es [WebAppNet](https://webapp-net.com/). Esta es otra libreria que simula la UI y animaciones del iPhone OS. Puedes ver un demo de sus features (desde un móvil preferiblemente) en [este sitio](https://demo.webapp-net.com/). A continuación un preview de este demo, visualizado desde un iPod Touch:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiSZIvCGSESKfQPwJotA02_5olqeotazNmZVAIZpzAgvKKzxK3jOlUjNJXtrIM-XpOk3gzPYl70crU9D4LxP1S3omzFIenroiEruopLoD7c4nkQ8rKNHSSMAW8TNrWQusVFp4fN6Wo1ea8/s400/WebAppNet-demo-itouch.jpg)    
"Aplicación Demo que muestra la apariencia y features de la libreria WebAppNet, visualizada desde un iPod Touch"

Por último los invito a leer el siguiente [post recomendado](https://woorkup.com/2010/01/10/best-practices-to-develop-perfect-websites-for-iphone-and-mobile-devices/) para desarrolladores de sitios web para iPhone y otros móviles, con muchas otras ideas a tomar en cuenta a la hora de realizar nuestros sitios. En [este otro sitio](https://patterns.littlespringsdesign.com/index.php/Main_Page) tambien pueden encontrar documentacion, patrones de diseño y guias para crear sitios para móviles. Además, no te olvides de comprobar la compatibilidad de tu sitio mediante [este test](https://www.w3.org/2008/06/mobile-test/) del w3c especifíco para sitios web orientados a smartphones.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEj72yKflF5B6F1VBOcf3Ih8zzggkjbe1T_eMmgMwOeXhzudBSSpYKjFzNF2osY0I-vLkyjs52K4DODeXDYJB4WWFytGWBki0lsdMsGW-2OSTD3OtaMgDTBORGszsxXpv_EgD_W86qbLYww/s400/WebCompatibilityTest-itouch.jpg)    
"Resultados de Mobile Safari en el Test de Compatibilidad de Browsers para Móviles, del w3c"

Si necesitas ofrecerle a tus usuarios algo mas que contenido y los navegadores para móviles no te lo permiten, siempre puedes optar por programar una aplicación específica para smartphones haciendo uso de los SDKs que provee cada productor de smartphones.