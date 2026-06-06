---
title: ¿Bing le esta copiando a Google?
date: 2011-02-02
author: Rodrigo A.
tags: google, copia, trampa, bing
draft: false
post_id: blog-3515952828243908885.post-8232212087694880166
---

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEha0_ejpTiIUQIGGTVhBcrASik4FYrDcWUp_br6nUjlJv8Wu2HapDv54f5BuXRjpIG2T2-ZVB5vvWApR3WIqOG53o40809n8EMTzx88frE_QDKBJsv7AknRFgcoYZprnYfc-ksgCibiyZqs/s200/Bing-vs-Google.jpg)    
Para estas alturas, muchos ya se habrán dado cuenta sobre el incidente llamado "BingGate", en la que Google publico esta increíble historia en [searchengineland.com](https://searchengineland.com/google-bing-is-cheating-copying-our-search-results-62914) (con pruebas) de que el popular buscador Bing esta copiando y mostrando resultados de Google como parte de sus búsquedas. En pocas palabras, Bing esta (supuestamente) haciendo trampa. De forma muy resumida este es el problema:

> Bing esta usando la información de clicks de donde sea para mejorar sus resultados.
¿Y ese es el problema?

Si, verán, los buscadores web solían trabajar de la siguiente forma (perdón por un ejemplo tan sencillo):

Si la pagina A contiene la palabra "pupusas", entonces es acerca de pupusas.

Google se volvió popular por la siguiente mejora, a la que ellos llamaron [PageRank](https://en.wikipedia.org/wiki/PageRank): Si la página A (que tiene la palabra pupusas) hace vinculo a la pagina B con esa palabra, ese vinculo es un buen indicador de que la página B también es de pupusas. Especialmente si la página A es un sitio de confianza (varias paginas lo vinculan).

![image](https://upload.wikimedia.org/wikipedia/commons/thumb/6/69/PageRank-hi-res.png/220px-PageRank-hi-res.png)    

A grosso modo PageRank determina el orden en el que aparecen las paginas cuando se busca algo en Google. Existen un sin fin de indicadores que influyen en el PageRank, entre estos inclusive: la velocidad con la que carga una página web. Bing por su parte, agrego "click data" o "click information" como indicador para sus propios resultados (ClickRank), que funciona así: si estas en la página A y haces click en el vínculo de dice "pupusa" y te lleva a la página B, ese click hace más importante la página B, que un vínculo que diga "carro" a la página C.

Más información sobre PageRank aquí: [https://en.wikipedia.org/wiki/PageRank](https://en.wikipedia.org/wiki/PageRank)

¿Aun no esta claro el problema?

Google es un sitio web, y como tal muchísimas personas hacen click en sus resultados... así que alguien en Microsoft se le prendio el foco (le hizo Bing!) y dijo tengo una idea:

> ¿Qué sucede, si medimos los clicks en los resultados de Google, y eso nos sirve para posicionar los resultados en nuestro buscador?
Ese es el problema, y por eso es que Google saltó.

¿Pero cómo rayos hacen eso?

Existen varias teorías, las cuales son muy poco probables, la que más tiene sentido (para mi) es la de la "Bing Toolbar". Si tienes instalada la Bing Toolbar, y haces una búsqueda en la página web Google, supuestamente la Bing Toolbar manda a Bing:

1. La consulta que realizaste. 2. La página web en la que hiciste click. Todo eso para mejorar el los resultados de Bing. Si bien aun no esta claro el "como", la evidencia existe y es bastante interesante:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiNiA8e-zeehP2eqc5fMYOcOqPwuEK0EtEZPVbhtkgORe8AU_08Noarjhgf20XfWs3K1A5ocRO01qu7HR-LHpo20dI32GapyRryTQCdNneyy-5HkL3zQcIo_uQHra_xN-tZSNpb7Iy2_FAG/s320/example-1-500x122.png)    
Búsqueda de "hiybbprqag" en Google

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEh34o1QiKbBtcgIi_LEnPQ7OVyGetmH2-yEya4VACBd_aKf6zWMX18oxrCarXufYF-JjEZ7hRy7o_FznESszY1IehrrTiXBGP6AvzEg3Ocj0_L1ZT7K7JO_M_8Mb6SkFnZz6CTkZbYYvbB-/s320/example-1-bing-500x129.png)    
Búsqueda de "hiybbprqag" en Bing

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgKh86Mbwg7M-3yhoehDJ0pAaR5eA3f7uHYTcvvxkGffLgFyjHJgiRJtBlfGCaItNwtvx-VoIKejesVu00u3ecnzvNm3UMx29QBp_eZVSF2RWC6gFdSfS3f1k88EPFtmhTxqjHFiAs_u7dM/s320/mbzrxpgjys-Google-Search-500x125.jpg)    
Búsqueda de "mbrzrxpgjys" en Google

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEj6HZwttGk2peE3kVhs1aTU2ckE3mTJ0-6I3bhSevw-veJpO5KXFYmmT1TJXlh9LbS1Yrc4Op3SLER5YN3mJ9jKFQDEDHyJpGomdaRsFQFaqY7XEUPN58mOEkgncncsRlAS9F41Mo0E-YLf/s320/mbzrxpgjys-Bing-500x129.jpg)    
Búsqueda de "mbrzrxpgjys" en Bing

Vean más ejemplos aquí: [https://searchengineland.com](https://searchengineland.com/)

¿Cómo me afecta eso?

Los que más se benefician son los usuarios de Bing que tendría (supuestamente) resultados bastantes similares (o sino mejores) a los de Google.

¿Y eso es bueno o es malo?

Es malo para Google, ya que Bing ahora posee el [29% del market share de busquedas en Internet](https://techcrunch.com/2011/01/03/compete-says-bings-total-u-s-market-share-grew-to-29-in-november/), lo cual es bastante impresionante. Así que es lógico que Google se tome la molestia de incriminar a Bing y desenmascararlo públicamente. Probablemente y a pesar de la mala publicidad, Bing simplemente siga mejorando la relevancia de sus resultados. Google es el "status quo" en términos de búsquedas, y si Bing lo va a retar e intentar quitar su lugar, no me extraña que se las ingenien para mejorar sus resultados, aunque no sea de la mejor forma.

Pueden encontrar más información sobre estas noticias en el Blog oficial de Google: [https://googleblog.blogspot.com](https://googleblog.blogspot.com/), y las respuestas de Microsoft en ZDNET: [https://www.zdnet.com/](https://www.zdnet.com/)

Y ustedes, ¿qué opinan al respecto?