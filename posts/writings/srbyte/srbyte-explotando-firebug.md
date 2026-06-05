---
title: Explotando Firebug
date: 2010-01-19
author: Robertux
tags: firebug, debug, firefox, programacion, javascript, unit test
draft: false
post_id: blog-3515952828243908885.post-8123051320625798973
---

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEikxmMCiSMWLCKcTwqpyVbHzXuE0fNIwNeUouWjr1O_EnsfmtllJvH17TOavNe-3qmjsxH2cjAR52xSEIld0RxHuxnb84VXFh11ex8RUY5y7woCWaN4dFQhTh6DpH0J1W5OCODqbwcMO9s/s400/firebug-logo.jpeg)    

Esta es la continuación del [post de herramientas de desarrollo](https://www.srbyte.com/2009/11/explotando-la-web-developer-toolbar.html) que puedes utilizar usando add-ons de Firefox.

Firebug ([link](https://getfirebug.com/)) es uno de los add-ons para Firefox mas complejos y mas populares en los círculos de desarrolladores. Te permite modificar a tu antojo la estructura del DOM de cualquier pagina en la que estés navegando, modificar sus clases CSS y depurar sus scripts, [monitorear http requests](https://www.srbyte.com/2010/01/monitorea-los-ajax-httprequests-con.html), ejecutar scripts sobre las paginas web, perfilar los tiempos de carga/ejecución de scripts, etc. aplicando todos los cambios realizados en caliente, visualizándolos al muy estilo WYSIWYG aunque obviamente todos estos cambios que se realizan no se aplican en el servidor. Entre sus usos mas comunes es depurar código javascript o aplicar cambios en el HTML/CSS asegurándose que se verán correctamente ya que estas realizando estos cambios directamente en el navegador.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiF01co48_Wm-yR9CeA_BoprJjEZ0HWX7njYr2ldF9-ojy5ArlGZZs8taUe9fwY-SfUWCRBtQ0oaHnqale9cxfoMb58NhEVR8RYL6Ct1etSlhFsefE-zFWdG9LzFaxpldUxs76c2jpEpFc/s400/firebug-editing.png)    
"Editando en caliente las propiedades CSS de la pagina de google.com desde Firebug"

La característica especial que distingue a este add-on es la posibilidad de ser extendido mediante mas add-ons (sub add-ons) los cuales le permiten hacer muchas mas tareas de las que actualmente permite.

CodeBurner for Firebug ([link](https://addons.mozilla.org/es-ES/firefox/addon/10273)) Este add-on te muestra una nueva pestaña después de la pestaña DOM, con la palabra "Code Example" en la cual, cuando estas inspeccionando el HTML de un sitio y tienes dudas sobre algún tag, este te muestra una pequeña referencia con la descripción de dicho tag y un sencillo ejemplo.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiOc5c32ld5WgFBkasIVujVAhUFWX4mVK2xtc4DHdKWzZ9J0VE5lGAMVo0M5Y7NW-NHN4eGHTxJprISBbJDx4NeELK9i02ob3-m8zkGGExNFYwb53pqo4dkjtlhiCd1KYcMxpHuvFuwTQo/s400/codeburner.png)    
"Visualizando una referencia sobre el tag desde Code burner, el tag se encuentra en la pagina principal de flickr.com, como se puede ver en la imagen"

FireFinder for Firebug ([link](https://addons.mozilla.org/es-ES/firefox/addon/11905)) Este add-on te permite hacer búsquedas de cualquier elemento HTML dentro del sitio que estas visitando. Esta búsqueda puede ser del texto contenido dentro de alguna etiqueta o mediante el uso de selectores similares a los usados en [prototype](https://www.prototypejs.org/) o [jquery](https://jquery.com/). Una vez encontrado el elemento, te lo remarca dentro de la pagina web con lineas punteadas rojas y te muestra un vinculo llamado "inspect" el cual te permite ubicar este elemento dentro de la pestaña "HTML" de Firebug. Como bonus también te permite generar un vinculo para FriendlyFire usando [jsbin](https://jsbin.com/), con el cual compartes el código HTML mediante una URL para poder ser visaulizado y editado por otros, muy al estilo de [pastebin](https://pastebin.com/).

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjLsjDtTKOH6v98cIsLNkGH7mt923iqV9sfL_0qpyqXbA0XXkOFtTL4aipCG_YwdvfzFIKXHQblSKEo2BJ4momqDrj_tKdtksq2TNJwOnxke6jlQI7qyEJdLW1DiWPAXp-RfLOVwPGz_O4/s400/firefinder.png)    
"FireFinder, remarcando el logotipo de twitter en la pagina de twitter.com, el cual fue buscado usando el selector #header2>a.logo"

Firecookie ([link](https://addons.mozilla.org/es-ES/firefox/addon/6683)) Este add-on de brinda la capacidad de monitorear y modificar las cookies que están siendo ocupadas en los sitios que visitas.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjdTtWgZNbRSjAdqRp8qyD_KDYa1MNF8dPMgbcYnAFTfRWP5hR3dIE6OIwrUoMrq2UJaiXKaHMgSECgoAKeUFmVNyTQkjRWeRVwfuq94U663JmTbgRbqogdIuAQZ9fi5ZLkQZC_OobtpGE/s400/firecookie.png)    
"FireCookie, mostrando las cookies que se almacenan localmente al entrar al sitio delicious.com"

Firiepicker ([link](https://addons.mozilla.org/es-ES/firefox/addon/15032)) Es un sencillo add-on que te muestra un selector de color (color picker) cuando dentro de una pagina web estas inspeccionando las propiedades css relacionadas con el color de un elemento. De esta forma puedes seleccionar el color de manera visual y este te escribe su respectivo codigo RGB.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhqsbOWrz7GFlVk4fn4Fpj3WdNcbr7UZ43HPfT3k5peSimeTMWJx9Q_JfjPbh02qq8tL2VqFbaE5OfFYqWwY53VGT-5hu-NViAxcUPbr83IgPf6w36SDd3e2QAgI8ZzXA59ivWlEAIAtBE/s400/firepicker.png)    
"Seleccionando un nuevo color usando FirePicker para la propiedad CSS background-color, en uno de los elementos del sitio deviantart.com"

Fireunit ([link](https://fireunit.org/)) Con Fireunit puedes escribir [Unit Tests](https://www.srbyte.com/2008/12/herramientas-del-programador-unit.html) para Javascript y ejecutar/visualizar sus resultados desde una pestaña en Firebug. Esto lo haces desde la nueva pestaña llamada "Tests" que te aparece en Firebug al instalar este add-on. El modelo de código para estos unit tests aparece en el mismísimo código fuente del sitio de fireunit, cuando lo inspeccionas con Firebug. Dicho código se muestra a continuación:

```
// Some examples of using FireUnit if ( typeof fireunit === "object" ) { // Simple true-like/false-like testing fireunit.ok( true, "I'm going to pass!" ); fireunit.ok( false, "I'm going to fail!" ); // Compare two strings - shows a diff of the // results if they're different fireunit.compare( "The lazy fox jumped over the log.", "The lazy brown fox jumped the log.", "Are these two strings the same?" ); // Compare a string using a regular expression fireunit.reCompare( "The .* fox jumped the log.", "The lazy brown fox jumped the log.", "Compare a string using a RegExp." ); // Display the total results fireunit.testDone(); }

```

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhUgX5nxX1fSQ_hxTK5U5QnIdQlAWmhGx2Qs2n9g_JHnn7xNYXnDGW-y0DxDp6ce6B5nNTE-gpeBQtVKl5XOwx0nhHlRJOt7x55Li1J5ikBiM6WWkL_0q7_zeOPmpa4wAGMSCKxVEgoqFY/s400/fireunit.png)    
"Visualizando con FireUnit el resultado de los unit tests de ejemplo que vienen en el sitio fireunit.org"