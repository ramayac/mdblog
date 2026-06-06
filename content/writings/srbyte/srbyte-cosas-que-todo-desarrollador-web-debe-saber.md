---
title: Cosas que todo Desarrollador Web debe saber
date: 2008-10-24
author: Rodrigo A.
tags: software, web, desarrollo, tips
draft: false
post_id: blog-3515952828243908885.post-4233864346734014638
---

Dejo una lista de cuatro cosas que todo desarrollador web debe saber, también queremos saber si desean agregar algo o profundizar sobre algún ítem especifico:

1. Las Tablas ya son historia... Existió un tiempo en que todo mundo usaba tablas para crear el diseño de una pagina web. Las tablas no fueron pensadas para el diseño de una pagina, sino mas bien para presentar "datos tabulares". La verdadera forma de diseñar una pagina web, su disposición, apariencia y demás, es a través de CSS (Cascade Style Sheets). CSS no solo ayuda aclarando el uso de las etiquetas en HTML, sino que también ayuda de la siguiente forma:

- Los buscadores se llevan mejor con CSS que con las tablas.
- Permite mucha más flexibilidad y posiciones que las tablas.
- Ayuda reduciendo código HTML/texto, lo que implica menor tiempo de carga.
- Permite mostrar el principal contenido del sitio mientras las imágenes o gráficos en general cargan de manera paralela.

2. El buen rendimiento del sitio es obligatorio...
![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhz7PJdFZ5owgaQJAwglkSMV5EhkFuJXCTbrh6v0Mz-Fls6htGUfIGiszNMUfVJrAt6xfAQIe29phMIfEHKZH7dLyUO5hjgqhZ_BDHQ7WUanLan1WHI4T8HCgvlJzQSlTYFCY_bHYkBmYc/s320/compress_files_icon.jpg)    En la etapa de desarrollo de un sitio web, los desarrolladores usualmente no se preocupan mucho por los tiempos de carga y "round trip" (lo que tarda la información del sitio en llegar al cliente y lo que tarda en regresar). Ahora el buen rendimiento de un sitio web es una parte integral en el desarrollo web (y claro, en cualquier tipo de desarrollo de software). Algunas cosas que se deben de considerar en la etapa de planeación del proyecto o de desarrollo del mismo son:

- Minimiza, une y comprime los archivos de CSS y que contengan JavaScript; a esta técnica se le llama: [Compresion HTTP](https://www.google.com/search?num=30&hl=en&safe=active&client=iceweasel-a&rls=org.debian%3Aen-US%3Aunofficial&q=http+compression&btnG=Search&aq=f&oq=) (que hace uso de Content Encoding y Transfer Encoding).
- Usa CSS Sprites, que consiste en crear un solo archivo con todas las imágenes que se necesitan en la pagina, y luego, se extraen indicando la posición de cada una. Un ejemplo sencillo es esta imagen que cualquiera puede encontrar entre los recursos que utiliza Gmail:[https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhwAfzjjRm5S5aVU-ibWiMvXT6iQp6XoQSLIwQu1yjcGamj31pPzXo9v2NJdlgkz_UTUBSxAzlDsJ9y4hwix_OLF3Rjn71z_8dqtKKoblZj3FXKdNdTegv-Vqm8Elr2anu7SG9b8ewO_AU/s1600-h/gmail-icons_ns2a.png](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhwAfzjjRm5S5aVU-ibWiMvXT6iQp6XoQSLIwQu1yjcGamj31pPzXo9v2NJdlgkz_UTUBSxAzlDsJ9y4hwix_OLF3Rjn71z_8dqtKKoblZj3FXKdNdTegv-Vqm8Elr2anu7SG9b8ewO_AU/s1600-h/gmail-icons_ns2a.png) ![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhwAfzjjRm5S5aVU-ibWiMvXT6iQp6XoQSLIwQu1yjcGamj31pPzXo9v2NJdlgkz_UTUBSxAzlDsJ9y4hwix_OLF3Rjn71z_8dqtKKoblZj3FXKdNdTegv-Vqm8Elr2anu7SG9b8ewO_AU/s320/gmail-icons_ns2a.png)    
"CSS Sprite, usado por Gmail."
[https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhwAfzjjRm5S5aVU-ibWiMvXT6iQp6XoQSLIwQu1yjcGamj31pPzXo9v2NJdlgkz_UTUBSxAzlDsJ9y4hwix_OLF3Rjn71z_8dqtKKoblZj3FXKdNdTegv-Vqm8Elr2anu7SG9b8ewO_AU/s1600-h/gmail-icons_ns2a.png](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhwAfzjjRm5S5aVU-ibWiMvXT6iQp6XoQSLIwQu1yjcGamj31pPzXo9v2NJdlgkz_UTUBSxAzlDsJ9y4hwix_OLF3Rjn71z_8dqtKKoblZj3FXKdNdTegv-Vqm8Elr2anu7SG9b8ewO_AU/s1600-h/gmail-icons_ns2a.png)... de esta se extraen los iconos de estado para el chat de Gmail.
- Coloca todas las inclusiones de CSS al principio de los archivos HTML, y todos los JavaScripts justo al final, y evita tener CSS y JavaScript en medio del código de la página.
- Si trabajas en una empresa grande, procura utilizar CDN (Content Delivery Network), es decir: tener varios servidores dedicados a un contenido en especifico, como solo imágenes, animaciones flash, streaming de vídeo, etc. O tener varios para diferentes áreas geográficas (usando mirrors, GeoIP, etc).

3. Diseño amigable para los buscadores...
![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhQP9Agf4uSPa-mPE5FZzCvf0uQ6OThBZU_CwERA46Io3psYeWDoaZuhIzxlAfWlx9HQ6ef6iABO6MelLQP9beXlbHWDvtaZx5laWxZR6dQUbtSKvvFVPkJOayEDySx8b-DORgTOePpBgM/s320/548130658_b776b25bf6.jpg)    En vez de pensar constantemente en SEO (Search Engine Optimization), hay que detenerse un poco y hacer que el contenido del sitio web, sea fácilmente indexable por los buscadores más populares. Ya no hay que esperar a un "experto" para que realice esas sugerencias, es ahora una obligación extra de desarrollador web preocuparse porque todo el sitio sea "amigable" para el web crawler (el programa que indexa el contenido de los sitios web).
Dos cosas que puedes practicar en tus sitios para hacerlos más amigables para los buscadores son:

- Usa los atributos alt y title en las etiquetas de las imágenes y de vinculos.
- Usa el atributo rel=nofollow cuando sea aplicable.
![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiP7lXqvUgPHIAAdwpQbpBaCK3EPY49js5VIDUt72ZLyKwgCqJIotrbVXWpgJmq05XbaM86dy71z2287EP4nrPLOpCGuTp-iLvPWo7JhKl4GkVQcKs2Ad0nps3P3cjQjy9hSMO-r0bUqpk/s320/pcweenies_1039.jpg)    4. Usa MVC ...
[MVC](https://en.wikipedia.org/wiki/Model-view-controller) (Model-View-Controller) es una "arquitectura" (de código) para cualquier aplicación general, aplicación web o sitio web. MVC consiste en separar la presentación del sitio (View) de la lógica del negocio (Model) y el control (Controller) de las dos anteriores. Es decir: separar el sitio o aplicación en capas, siempre que sea posible o aplicable, y que se tenga el tiempo para hacerlo.

Y tu, ¿Qué otras cosas crees que debería de saber un Desarrollador Web? ¿Quieres que hablemos de algún tema especifico? ¡deja tu opinión!