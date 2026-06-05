---
title: Bootstrap: Tus Aplicaciones Web con la UI de Twitter
date: 2011-11-10
author: Robertux
tags: twitter, github, css, programacion, bootstrap, less, javascript
draft: false
post_id: blog-3515952828243908885.post-8764236823225518557
---

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEi709Y7Y8TlHtmFsSRxAoauuest4TJ_z1eWHTO_5aO8KjBNBqtHHLRawk0tJoTTvA5NlmWjnkaCYaWz61FhHQ1jmtzhMX-iWImvjvf6wu_sYijm-6wtkQHxBQCsWHKGhewt5uEgbv_kqLQ/s320/Screen+shot+2011-11-09+at+9.10.43+PM.png)    

"Portada del Toolkit en Github"

Luego que twitter [actualizara por fin su interfaz web](https://twitter.com/newtwitter), unos se dedicaron a hablar maravillas de la misma mientras que otros se dedicaron a acabársela con cuanto bug le encontraban. En lo personal, a pesar de la tradanza en su carga y falta de una que otra verbosidad, me parece muy bien diseñada y elaborada.

Recientemente buscaba un sitio para validar en línea expresiones regulares y me topé con [este sitio en RegExPlanet](https://www.regexplanet.com/simple/), cuya interfaz me pareció muy similar pero no fue sino hasta que hurgué en su CSS y Javascript y busqué sus referencias que me topé con Bootstrap y fue cuando recordé la familiaridad del sitio, estaba basado en la UI de twitter haciendo uso del toolkit denominado [Bootstrap](https://twitter.github.com/bootstrap/).

Tal como lo mencionan en su portal de Github, Bootstrap es un conjunto de archivos CSS (Y sus respectivos LESS), HTML y Javascript con los cuales te ahorras mucho trabajo armando layouts, estilizando inputs y utilizando componentes un poco más complejos ya construidos a base de HTML como son el caso de la barra de menú (o barra de navegación, como quieran llamarle) en la parte superior con posición Fixed, tabs, formularios modales, tooltips, etc.

Pueden descargar este toolkit en [su página de Github](https://github.com/twitter/bootstrap) junto con la documentación la cual te explica lo fácil que es integrar esta UI en tus sitios web.

Basada completamente en la Apache License versión 2.0 lo cual significa que no tendrás que preocuparte por que alguien te quiera demandar por utilizar los mismos componentes UI con los que está hecha la interfaz actual de twitter.

Les adjunto además [un ejemplo editable elaborado en JsBin](https://jsbin.com/ubaruf/21) mostrando cómo con solamente agregar el CSS de Bootstrap automáticamente estiliza los input y con algunos divs con clases específicas es posible construir el resto de la estructura básica de un formulario web.