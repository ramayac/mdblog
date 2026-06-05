---
title: Java EE 6
date: 2009-12-20
author: Rodrigo A.
tags: glassfish, java, php, netbeans
draft: false
post_id: blog-3515952828243908885.post-6003911371860394385
---

Casi 10 años han pasado, desde Java EE 1.2, y Sun anuncia ahora que [Java EE 6 esta listo para las empresas](https://java.sun.com/javaee/). En los tres años desde la mas grande actualización de la plataforma Java EE, un sin fin de cosas han cambiado en Java, y EE 6 es un reflejo de esos cambios.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjheqX2zBRu89Vs_-pUme5V14pcS-lWMn1FYZdTHjEEPIGQ5voeb4GrefaDwy1MVA35WtpPaKMZdGvx9R-PaWS7wfW-LWCS4hcej-AFhqhg0-5Y0XKcB6RVOTzo9X4v9ophrv0zV7G8al_Z/s320/5ehhcx.jpg)    

Entre las mejores características, es la instalación del perfil Web, soporte para servicios RESTful, y la inclusión de ultimo minuto de la inyección de dependencias. Sun también libero [NetBeans 6.8](https://netbeans.org/community/news/show/1449.html) y GlassFish Enterprise Application server 3. Ambos compatibles con EE 6. Lo bueno de esta versión, es que quizás en la empresa donde trabajas no necesitas todo el ecosistema de Java EE, así que solo decides instalar el "perfil Web". Este concepto de perfiles (que es una respuesta a las quejas de los desarrolladores de Java a través de los años) se emplea para instalaciones con propósitos específicos. Inicialmente solo el perfil Web esta disponible, pero Sun liberará más combinaciones en el futuro. Imaginen un perfil solo para servicios u otro para desarrollo, es menos trabajo y mas controles para los encargados de mantener y monitorear los servidores empresariales... o al menos, esa es la idea.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgOulWYX1rOeNjzWnVWNtmRSKR9SEBExCO_gMSwLWrszxMvg2oPH0CjmG-u6hMV0yrFioc5GomzpsmbPji0bcho757ilpEB6rdIaTzRhoLnNNOg_d8XqqGvjqKyknicuZrgfDV9BNPRXm4A/s320/glassfish_front_cover_full.GIF)    

Tambien GlassFish se puede "adelgazar" para que cumpla con propósitos web específicos, recordemos que la versión 3 de GF esta basada en una arquitectura muy modular y flexible, lo que procura obtener mejor rendimiento del hardware.

Un detalle interesante es el JSR 330, el tardista a la fiesta de Java EE 6. Esta especificación se origina en Google, y existe unicamente para proveer la característica de inyección de dependencias, característica de la que después hablaremos con gusto.

Por otra parte, una excelente noticia para los desarrolladores de PHP, es que NetBeans 6.8, provee soporte para la versiones 5.3 de ese lenguaje. Y tambien mejoraron la integración de Ant y Maven. A ver si algún "phpero" nos cuenta como esta la integración de NetBeans con este popular lenguaje.

Así que... si tu negocio, empresa o lugar donde trabajas, esta usando una versión prehistórica de Java, seria bueno de que les comentes de esta nueva versión, y también seria bueno que aproveches la oportunidad para ponerte al día con todas estas tecnologías que no tardaran mucho en llegar a los negocios.

Si quieren saber más sobre Java EE 6, recomiendo este genial articulo en  [DevX.com](https://www.devx.com/Java/Article/42351/1763/page/2), y si lo desean descargar [hagan clic acá](https://java.sun.com/javaee/downloads/index.jsp?userOsIndex=6&userOsId=windows&userOsName=Windows).

¡Saludos!