---
title: Asegura tus Sitios Web con OWASP
date: 2011-12-16
author: Robertux
tags: XSS, programacion, java, seguridad
draft: false
post_id: blog-3515952828243908885.post-5224065773591661222
---

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEh34pXNlgBWsi1BkJMejyGkRB26QRRv-NAArONxvoDmfnEIFTz9aqFupS5piF3HvBHvYdbFp3SugfNCJzxpCdEACJ5JWuQQRrAuLj92LEu_8GJ6gkPHU2WiBwPWsSMGsWq4voyL2kwSe5k/s320/www.owasp.org+screen+capture+2011-12-15-23-32-23.png)    

"OWASP - Logo de la página oficial"

Uno de estos días estaba buscando en internet un par de técnicas para librarse de ataques XSS y SQL Injection cuando me topé con The Open Web Application Security Project, ([OWASP](https://www.owasp.org/index.php/Main_Page)). Un portal/wiki creado por una comunidad con filosofía libre en el cual encontré muchísimo material referente a las formas más populares de ataques en la web. Encontré tan valiosa información que tomé a bien compartirlo con ustedes.

La referencia y documentación sobre vulnerabilidades comunes en el diseño y técnicas de programación de aplicaciones web es muy extensa, descriptiva y variada, encontrando desde artículos en PDF hasta vídeos de youtube. Entre lo más valioso que encontré en este sitio fue [ESAPI](https://www.owasp.org/index.php/ESAPI), Enterprise Security API. Este consiste en un conjunto de librerías desarrolladas por la comunidad de OWASP y liberadas bajo la licencia BSD en distintos lenguajes de programación las cuales encapsulan controles de seguridad para evadir las vulnerabilidades más comunes en los sitios web incentivando además al desarrollador a hacer uso de los patrones de diseño básicos recomendados por las comunidades de software.

Me tomé la libertad de descargar y probar la librería en su versión para Java 1.5 para jugar un rato con ella pero se ve un tanto compleja como para juzgarla a la ligera. Por tal razón me reservo mis comentarios de su uso y cuando haya podido implementarla actualizo con mis comentarios.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiswO6-PXLRQq9Ud1psch_OnCNJn9i0aEtgAJ5WdDDRjSErckYUiBXaG_oNCMbypE2wdqGc8VNLautP2K8OLg2otyA5p3xdusuK3BjUPKnMAM-IC7_5kSdhTiSP5f6Al7Pc2tLMUu3iZD8/s320/Screen+shot+2011-12-15+at+11.39.56+PM.png)    

"Captura de pantalla del Read Me"

Como ejemplo les dejo un vídeo en el cual narran de manera bastante comprensible las variantes y ejemplos de un ataque por Cross Site Scripting (XSS) así como consejos básicos de cómo proteger tus sitios ante este ataque:

"OWASP - Cross Site Scripting"

Pueden encontrar más información en el la [página principal de OWASP](https://www.owasp.org/index.php/Main_Page).