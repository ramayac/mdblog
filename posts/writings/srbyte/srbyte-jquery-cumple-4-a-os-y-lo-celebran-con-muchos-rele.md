---
title: JQuery Cumple 4 Años y lo Celebran con Muchos Releases
date: 2010-01-15
author: Rodrigo A.
tags: framework, programacion, html, jquery, javascript
draft: false
post_id: blog-3515952828243908885.post-8786999579869923825
---

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgrkh0cxEBi_V8a_oPjbhWKsJvHgS0C0S_2O4bqAbdXQ7KHWxgEFlNdMzVySEj6jvMbloj3fNxnWPM61yndXc0DeMmbnSi65MJpeEYDTk8uJHQDGw_LCOC7pqMfJ4YQyuO1uwNmBVGlJtM/s400/jQuerySite.png)    
"Entre las novedades de jQuery presentan 14 días de releases continuos"

El día de ayer, 15 de Enero, jQuery celebró su cuarto cumpleaños de haber sido liberado y desde entonces los desarrolladores de este popular framework lo están celebrando a lo grande con la liberación de la verisón 1.4, un nuevo sitiopara la API, mejoras en el performance, QA en vivo, etc. En total serán 14 días durante los cuales tendremos nuevas sorpresas y releases relacionados con este framework, los cuales pueden consultar día a día en el sitio [jquery14.com](https://jquery14.com/).

[jQuery](https://jquery.com/) es un framework javascript open source que te ayuda principalmente a interactuar con el DOM de una página web de una forma muy práctica y con muy pocas líneas de código fuente. Además simplifica la comunicación vía ajax, el manejo de eventos, la creación y manipulación de controles para crear interfaces gráficas ricas, entre otras cosas.

Su gran popularidad y su capacidad de extensión mediante plugins ha permitido que muchos programadores lo hayan enriquecido con controles muy útiles como el [datagrids](https://www.flexigrid.info/), [lightboxes](https://leandrovieira.com/projects/jquery/lightbox/), [gráficos](https://code.google.com/p/flot/), [skins](https://www.dfc-e.com/metiers/multimedia/opensource/jqtransform/), [color pickers](https://acko.net/dev/farbtastic), por mencionar unos cuantos entre otros cientos de plugins que existen por ahí en la web.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjaUL6D2Tud3WNkuCDwA2sA9lCAIilUgyG_H_LqZXcsLDEjoZM8G6jUbI4OIWQ0EHiorm_Jqk0qUtEhvboIdkyg4ZS-45kAQl6IjdfqCnzlym6LR6S2q-5VL2YteRJvzyqwDXrs4FBTYcQ/s400/flexigrid.png)    
"El plugin Flexigrid es de los más complejos y útiles que he utilizado con jQuery"

Lo que mas me encanta de usar jQuery es su facilidad para manipular el DOM. Como ejemplo, podemos ver el siguiente código fuente que utiliza sintaxis JSON de la nueva versión 1.4 para crear un DIV con id, propiedades CSS y una función para manejar su evento click para luego agregar este div al body de la página, todo de una sola vez:

```
jQuery("<div/>", { id: "foo", css: { height: "50px", width: "50px", color: "blue", backgroundColor: "#ccc" }, click: function() { $(this).css("backgroundColor", "red"); } }).appendTo("body");

```

Como muchos proyectos open source, jQuery sobrevive gracias a las donaciones altruistas hechas por los usuarios. Como parte de las sorpresas que los desarrolladores de jQuery nos han traido, también tienes la oportunidad de recibir un e-book gratuito por tu [donación al proyecto](https://jquery14.com/donate).