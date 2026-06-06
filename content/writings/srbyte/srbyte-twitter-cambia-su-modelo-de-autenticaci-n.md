---
title: Twitter Cambia su Modelo de Autenticación
date: 2010-05-15
author: Rodrigo A.
tags: twitter, seguridad, oauth
draft: false
post_id: blog-3515952828243908885.post-6456282532171020786
---

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgGAdYHv2lc5buGOSF0htvAZ2uxAjWEPYuDB444WxusZDcQ3r_DaTPaBpYHx77nfojSwoIkQc78jvjpjeca6YGkYZD40aBdDTjusYTPyW14xTrXqcA4wpveq6LvP5RGt3Ir44Ykx2ccyH4/s400/twitterauth.jpg)    Desde hace ya un buen tiempo twitter ha permitido a los desarrolladores de aplicaciones hacer login mediante dos métodos distintos: autenticación básica y [OAuth](https://oauth.net/). Mediante autenticación básica, nosotros le confiábamos nuestro nombre de usuario y clave al cliente de twitter y este lo utilizaba para intentar hacer login y permitirnos hacer updates y leer los timelines desde dicha aplicación. La segunda y más reciente alternativa es OAuth, es un método mediante el cual la autenticación del usuario se hace directamente desde una página de [twitter.com](https://twitter.com/) y si la autenticación es válida, twitter le devuelve un token a la aplicación para que esta nos brinde sus servicios con acceso a nuestra cuenta.

Definitivamente la segunda alternativa es mucho más segura ya que nuestros datos de usuario y clave de twitter eran únicamente compartidos con twitter.com y las aplicaciones nada más reciben un token de autenticación por lo cual estos no pueden abusar de nuestras cuentas y nosotros podemos revocarles el permiso de usarlas en cualquier momento desde la página de settings, en nuestra cuenta de twitter.

Debido a esta mejora en seguridad que se experimenta con el uso de OAuth, twitter por fin ha decidido dar de baja al método de autenticación básica de manera que OAuth será la única alternativa por la cual los clientes de twitter podrán autenticarnos. Esta noticia la han publicado en el blog [The Twitter Engineering](https://engineering.twitter.com/2010/05/tracing-traffic-through-our-stack.html), blog oficial de los desarrolladores de twitter.

Este cambio entrará en vigencia a partir del 30 de junio del presente año. Pueden seguir la cuenta regresiva en [esta URL](https://www.countdowntooauth.com/) y además seguir al usuario de twitter [@twitterapi](https://twitter.com/twitterapi) para obtener más información sobre cambios técnicos en twitter y su api.

En qué afectará esto a los usuarios?

En absolutamente nada. Simplemente tendrán una forma mucho más segura de iniciar sesión en sus clientes de twitter y un mejor control de cuáles aplicaciones tienen acceso a sus cuentas.

En qué afectará esto a los desarrolladores de aplicaciones para twitter?

En mucho! Hay una gran cantidad de aplicaciones que aún funcionan con la autenticación básica y a partir de la fecha antes mencionada estas aplicaciones simplemente dejarán de funcionar así que empieza la cuenta regresiva para que cada grupo de desarrolladores empiece a implementar OAuth en sus aplicaciones. Uno de los populares servicios de hosting de imágenes vinculadas a twitter, el conocido [twitpic](https://www.twitpic.com/), recientemente hizo el cambio a autenticación por medio de OAuth. Cabe mencionar que es de los servicios populares que más se tardaron en implementar esta medida pero al final de cuentas la implementaron a tiempo.

Si necesitas documentación para tener una idea de cómo implementar OAuth en tu cliente de twitter, puedes leer la [wiki de twitter](https://apiwiki.twitter.com/OAuth-FAQ) relacionada al OAuth, donde encontrarás ejemplos escritos en los lenguajes de programación más populares. También puedes hacer uso de apis de terceros que ya encapsulen la funcionalidad del OAuth, como por ejemplo [twitter4j](https://twitter4j.org/en/index.html) para java, [twitterauth](https://github.com/mbleigh/twitter-auth) para Ruby on Rails y [twitteroo](https://rareedge.com/twitteroo/blog/2007/02/10/twitter-net-api/) para C#.net.