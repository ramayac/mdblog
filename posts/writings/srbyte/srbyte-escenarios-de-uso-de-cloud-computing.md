---
title: Escenarios de uso de Cloud Computing
date: 2010-01-22
author: Rodrigo A.
tags: cloud, computing
draft: false
post_id: blog-3515952828243908885.post-4784643796373100132
---

En este breve articulo, se pretende mostrar tres casos puntuales de escenarios de uso de [Cloud Computing](https://en.wikipedia.org/wiki/Cloud_computing) y sus ventajas sobre el alojamiento web tradicional. Las aplicaciones web siempre se han instalado en servidores conectados a lo que ahora se denomina "la nube". Sin embargo, las demandas y tecnologías utilizadas en esos servidores, ha cambiado substancialmente en años recientes, especialmente con la entrada de proveedores de servicios como Amazon, Google y Microsoft.

Estas compañías ya tienen largo rato de proveer una infraestructura (hardware) flexible para aplicaciones web, de manera que esta se adapte y escale a diversos escenarios.

Escenarios de Uso...

Paga por lo que consumes

La instalación de aplicaciones web, hasta hace algunos años, era similar a la cuenta de tu teléfono fijo. Tenes que pagar una cuota de teléfono, consumas o no tiempo de llamadas.  Asi mismo, cuando compras alojamiento tradicional, compras espacio y pagas una "cuota fija", lleguen visitantes a tu sitio o no.

Cloud Computing cambia la idea del consumo fijo. Los varios recursos consumidos por una aplicación web (CPU, memoria, ancho de banda) se contabilizan por unidad (comenzando desde cero), como los contadores de agua de tu casa, y se paga de acuerdo al consumo de cada uno.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhTA7Nh4GsOoY7SmomhM89aFYFEykgTgmMvS8Th6ikErdIZlRNMj8YTy0KseeVJV3yzLxIpZPQ0LzEpXqpaGHS88rsP-0O39-rCwkCHMuEvOFlhamf0GB9xXX6-J_d6wihGaXagXpQz_-Fp/s200/Mtron_Pushes_SSDs_to_120_GB-1.jpg)    

"... paga por lo que consumes(CPU, espacio, ancho de banda)."

Ejemplo: Supongamos que estas haciendo un sitio para hacerle competencia a [Cherada.com](https://www.cherada.com/), y supones que llegaran al menos 100,000 Salvadoreños a registrarse en el primer año, y como quieres estar preparado, compras el

mejor plan de alojamiento viendo la demanda futura que tendrá tu sitio por ese lapso de tiempo. Pero por cosas del destino (digamos mala publicidad, plataforma y mercadeo del sitio) casi nadie llega en ese año, a lo sumo, recibiste 1,000 visitas.

Si este penoso escenario se cumpliera, terminarías pagando el año completo con el mejor plan, independientemente de que lleguen 10,000 o 1,000 visitas. Con un proveedor de servicios, solo pagas lo que consumes, y si bien tenes que pagar una cuota por el servicio, en el escenario anteriormente mencionado, sera significativamente menos que el del alojamiento tradicional.

Suministro para eventos momentáneo (One time event provisioning)

Las aplicaciones web a veces sufren de "picos" de trafico debido a la cobertura de los medios, como comerciales de TV, apariciones en artículos periodísticos, eventos sociales (elecciones) y naturales (sismos), etc. Para poder soportar este tipo de carga en el sistema, sin que este se vuelva inestable o se convierta en un "servidor de excusas", muchas personas creen que necesitan comprar infraestructura o "arrendar equipo" para solventar la demanda de las visitas.

Si tu aplicación esta apoyada sobre una plataforma de servicio de un proveedor de Cloud Computing, este será el que se adaptará dinamicamente para proveer el recurso (CPU, memoria, ancho de banda) para que tu aplicación web supla la demanda de información por los visitantes que llegan repentinamente a tu sitio, sin que gastes ni un centavo en equipo físico.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgrnLworS20b7Oc4sIDgT2YfbW5CSzU_KycxL5QwpSHIm-hO1oEugUAY2O7hf0_KkgQrzmVJLj-kb3_OvRvtRSHVJXVjJoEtOPuUqkcTdsqerMRsOFAl4ZzLUqOHv_XmBwAhszoPiKH5X2V/s320/servers.jpg)    

"... no inviertas en equipo innecesario."

Ejemplo: ¿Cual fue el sitio web salvadoreño que mas sufrió el año pasado por las elecciones? Exacto, [https://www.tse.gob.sv](https://www.tse.gob.sv/) . Con Cloud Computing, el TSE podría haber suplido a la perfección la demanda de clientes al sitio, sin la necesidad de comprar o prestar equipo para un evento que sucede cada cinco años.

Crecimiento automatizado y tecnologías escalables

Si se tiene la capacidad de soportar eventos momentáneos, entonces las plataformas de Cloud Computing también facilitan el crecimiento gradual de curvas que presentan las aplicaciones web exitosas. En los escenarios usuales de crecimiento para aplicación web necesitan equipo especial (balanceadores de carga y clusters), que al emplear Cloud Computing, queda en la abstracción de la plataforma de nuestro proveedor de servicio. Y adicionalmente, muchas plataformas de Cloud Computing, proveen soporte para un "tier" de datos que excede el rendimiento de RDBMS (Relational Database Systems), como: Map Reduce, Big Table, etc.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjnEBTySSqLmWeANhWTzOJAS03LmsKba1f0-tRK9hSawu-gXA-brTS96FsHbsOgGuOFvlxvMnqoZcN2ujo70o2Vv9xd_LMRADHSHoGp36WZmf3oDW0jR3KNt71w5jz7Tfs8bCMk6lDtdnmO/s200/cloud-computing2.jpg)    

"... crece sobre una infraestructura escalable."

Ejemplo: Supongamos una aplicación web, en donde se almacena los artículos de todos los blogs Salvadoreños, Twitts, fotos en Flicker, videos de YouTube, y noticias de los periódicos mas importantes de El Salvador, y lo hace de manera DIARIA. De manera que esta aplicación nos permite "ver el estado de El Salvador, en un día especifico" (¿buena idea no?). ¿A lo largo de 5 años, cuanta información (en TB) creen que tendríamos almacenada? ¿dos o tres TB? La idea es que la velocidad de búsqueda de la información debe ser la misma a lo largo del tiempo. Cloud computing nos permite: no solo almacenar toda esta información en la nube, sino que también buscarla con eficiencia y velocidad, para que sin importar la cantidad de información almacenada, la experiencia de nuestra aplicación web sea la misma siempre.

Espero que estos tres sencillos ejemplos, ayuden en alguna medida a comprender los escenarios de uso de Cloud Computing. Saludos!