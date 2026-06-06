---
title: Monitorea los Ajax HttpRequests con Fiddler
date: 2010-01-13
author: Rodrigo A.
tags: firebug, firefox, explorer, ajax, fiddler, javascript, internet
draft: false
post_id: blog-3515952828243908885.post-1437424478541686193
---

Como programador de sitios web basados en Ajax, tiendo a necesitar una herramienta que me permita verificar exactamente que le envió al server en un XmlHttpRequest vía Ajax. Los parámetros que van en el request, la URL a la que lo envió, la respuesta recibida etc. Así puedo determinar si un error se genero porque se enviaron datos erróneos al server desde el Ajax Request o si fue el server quien interpreto o proceso mal los datos.

Mi herramienta para realizar estos monitoreos fue Firebug, un add-on para Firefox que te permite depurar Javascript, inspeccionar/manipular el DOM y también inspeccionar/manipular los estilos de un sitio. Este además visualiza en su consola los diferentes HttpRequests que se hacen al server.

Como muchos sabrán, no siempre basta con realizar pruebas y depuraciones de tu sitio usando el browser Firefox ya que comúnmente los errores ser producen al navegar en tu sitio desde Internet Explorer o algún otro y como hasta donde he podido investigar, ningún otro navegador cuenta con un add-on como el de Firefox para visualizar estas peticiones de Ajax, no hay otra manera de hacerlo mas que con una herramienta externa al browser.

Buscando en la web encontré la herramienta llamada [Fiddler](https://www.fiddler2.com/fiddler2/). Este es un proxy web hecho con el framework Microsoft.Net para captura y depuración del trafico http que se genera en tu computadora. Este trafico es comúnmente generado por los browsers al navegar dentro de sitios web, dar click sobre algun hipervinculo de alguna pagina o cuando un evento javascript dispara un request XmlHttp.

![image](https://docs.google.com/File?id=djh57d7_34fwqwxjfj_b)  
"Captura de pantalla de Fiddler, monitoreando el trafico http local. Visualizando en la parte derecha el detalle de un Ajax Request de la pagina de twitter.com"

Como pueden ver en la captura de pantalla, Fiddler estomáticamente captura todos los requests http y los lista en la columna de la izquierda, mostrando los detalles de cada request en la columna de la izquierda, como por ejemplo los headers de la petición, parámetros y la respuesta obtenida.

Hay un pequeño problema con Fiddler cuando intentas monitorear el trafico de un server alojado en tu localhost. Esto es porque, como ellos mismos mencionan en su sitio: "Internet Explorer and the .NET Framework are hardcoded not to send requests for Localhost through any proxies" por lo cual, nos vemos en la necesidad de no utilizar la palabra "localhost" ni su IP equivalente 127.0.0.1 para acceder desde nuestro browser a los sitios alojados en nuestro server. Como alternativa podemos usar en su lugar, el nombre de nuestro equipo. Por ejemplo:

en lugar de: https://localhost:8080/miSitioWeb/index.jsp

tenemos que escribir: https://laptux:8080/miSitioWeb/index.jsp

Asumiendo que el nombre de mi equipo es "laptux".