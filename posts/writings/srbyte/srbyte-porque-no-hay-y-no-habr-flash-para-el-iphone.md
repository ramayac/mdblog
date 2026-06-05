---
title: Porque no hay (y no habrá) Flash para el iPhone?
date: 2010-02-12
author: Rodrigo A.
tags: html, ipod, iphone, javascript, flash
draft: false
post_id: blog-3515952828243908885.post-2900024211659290681
---

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjDoBXpPofbAoAl1-wHQJR2-YuvNlOHeHoFY1pqJr2WVRHVD28FbfajVv5kY01L-qjmgsqkuMyAWzOvXqMjRRuTec99TT_fPA_yEIAF5eA70enY3J6S2I4TtVeybdgwY4v1Xx0k3t75PMc/s400/iphone-no-flash.jpg)    
"Este es el mensaje que te aparece al navegar al sitio get.adobe.com/flashplayer informándote que no existe una plugin disponible para tu dispositivo"

Leyendo el articulo de [GadgetLab](https://www.wired.com/gadgetlab/2008/11/adobe-flash-on/) me he dado cuenta de muchas razones sólidas por las cuales Apple no permite que se puedan ejecutar aplicaciones Flash en su navegador Safari para el iPhone ni mucho menos como aplicaciones nativas del iPhone/iPod Touch. En si, estas razones se pueden resumir de la siguiente manera:

- Pérdida de la exclusividad, calidad en el diseño y usabilidad de las aplicaciones
- No mas necesidad de una App Store centralizada y amarrada a iTunes
- No mas restricciones ni espera por aprobacion de aplicaciones
- Baja significativa en el rendimiento y duracion de la bateria
- Posibles vulnerabilidades del player de Adobe que pudieran servir para hackear el aparato.

Esto demuestra que la falta de Flash no fue nada mas algo que Apple pasó por alto sino que fue una decisión premeditada. Ellos diseñaron el browser del aparato para que no fuera capaz de ejecutar este plugin y sus razones son sólidas. Esto y las [últimas declaraciones de Steve Job](https://techcrunch.com/2008/03/05/adobes-flash-not-good-enough-for-steve-jobs/) [s](https://techcrunch.com/2008/03/05/adobes-flash-not-good-enough-for-steve-jobs/) nos indican que no debemos esperar que Apple esté trabajando en un próximo lanzamiento de Flash para el iPhone.

Adobe Flash es actualmente utilizado en tres areas principales del web: reproductores de video, juegos online y sitios web con interfaces ricas. Estas caracteristicas se estan volviendo cada vez menos necesarias con el avance de otras tecnologias alternas que se estan volvieno un estandar en la web, por lo cual la necesidad de Flash decrece con el tiempo. Veamos algunos ejemplos:

Reproductores de Video: Con la llegada de [HTML5](https://en.wikipedia.org/wiki/HTML5) y sus implementaciones nativas de tags como [canvas] y [video], pueda que en un futuro ya no sea necesario utilizar Flash para realizar streamings de multimedia o para juegos online. Tanto la gente de [Youtube](https://www.youtube.com/html5) como los de [Vimeo](https://vimeo.com/blog:268) han lanzado ya sus versiones de video players usando tags [video] para browsers que ya poseen soporte para este nuevo standard del HTML.

Asumo que Apple optará en un futuro cercano por agregar soporte completo a las features de HTML5 en el mobile Safari mientras que los sitios que actualmente poseen aplicaciones Flash, las portarán a este nuevo standard, dejándolo de lado.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgT9LBz7EZ3QctZrLvw0CJNJ7VKIGTjjSL7ul8EVyR-bw7p96ugXp500Pfpn0GBUI9rP40-j3VDXQg8-oYikHWE6qMABH8zyVTM_PF1mHb-UiZG6OFKKWDNW-JJNBLnYeP_ytpx_JH4Eok/s400/youtube-html5-video.png)    
"Código fuente de un vídeo de Youtube mostrando el tag [video]"

Juegos online: Aceptémoslo. El iPhone y el iPod Touch se estan convirtiendo en una plataforma de juegos capaz de competir con el PSP y el Nintendo DSi, quiza no tanto en calidad y complejidad pero sí en popularidad. De hecho, al momento de escribir este post, leo [un artículo](https://www.tipb.com/2010/02/06/game-developers-iphone-nintendo-ds-sony-psp/) que indica que el iPhone se esta volviendo mas popular entre los desarrolladores de videojuegos por encima del Sony PSP y el Nintendo DSi. Es un fenómeno similar al que ocurrió entre el Nintendo Wii y el Sony PS3 en el cual, aunque el segundo tuviera mucho mejor hardware, el primero se volvió mas popular gracias a su simplicidad.

"NFS Undercover para iPhone. Uno de los juegos que mejor demuestra la capacidad que tiene este dispositivo como consola de videojuegos"

Sitios web con interfaces ricas: Antes Flash era la única alternativa para crear sitios web vistosos, con controles ricos y animaciones presentando una agradable experiencia de usuario con un framework bastante fácil de usar para los desarrolladores y diseñadores, permitiendo integrar la capacidad de Flash para crear animaciones basadas en fotogramas y lineas de tiempo, la galeria de controles ricos y el lenguaje de programación ActionScript. Hoy en día hoteles, restaurantes, sitios que promueven peliculas de cine y artistas musicales aun siguen usando Flash para darle vistosidad a sus sitios pero cada vez son mas los que prefieren optar por realizar estos mismos efectos usando nada mas que frameworks Javascript y todas las facilidades que ofrecen los web browsers modernos permitiendo una mejor integración de sus interfaces con servicios del lado del servidor y gastando mucho menos recursos en el cliente, ya que los sitios Flash tienen la caracteristica de consumir gran cantidad de ancho de banda y memoria además de no estar apegado a los estandares del W3C.

Hoy en día vemos que es mucho más conveniente hacer uso de librerías Javascript como [Scriptaculous](https://script.aculo.us/), [jQuery](https://jquery.com/) o [Dojo](https://www.dojotoolkit.org/) para interfaces ricas y animadas en lugar de Flash y aun más sabiendo lo que podremos hacer en un futuro cuando HTML5 sea un estandar popular en el web.

Debido a esto, la necesidad de un usuario por tener Flash en su navegador va desapareciendo segun va avanzando el desarrollo de las tecnologías alternas como HTML5 y las iPhone Apps, que como vimos, pueden reemplazar a Flash para las tareas en las cuales usualmente hacíamos uso de este. Por ello concluimos que esta es la razón por la cual los de Apple han decidido no incluir un plugin Flash en su navegador Safari y en su lugar, esperará que su App Store siga creciendo, ofreciendo una gran variedad de videojuegos a sus usuarios, así como también esperan que los desarrolladores se decidan de una vez por todas a reemplazar a Flash por animaciones e interfaces hechas con frameworks Javascript para así desaparecer casi por completo la necesidad de Flash.