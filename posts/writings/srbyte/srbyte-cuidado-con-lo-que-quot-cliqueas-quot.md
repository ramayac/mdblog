---
title: Cuidado Con Lo Que &quot;Cliqueas&quot;
date: 2010-04-02
author: Rodrigo A.
tags: seguridad, internet
draft: false
post_id: blog-3515952828243908885.post-5302489112529415368
---

¿Alguna vez  se han preguntado lo peligroso que podría ser el hacer clic sobre un link?. La respuesta es sencilla, puede ser muuuuuuuuuuy peligroso. Traigo a colación este tema, ya que en varias ocasiones he iniciado sesión en Windows Live Messenger y recibo mensajes de contactos desconectados. Esos mensajes contienen links - muy sosprechosos - y que hacen alusión a fotos bastante subidas de tono - básicamente, "pornografía" - o contenidos que para "el usuario común" son considerados atractivos e inofensivos. He aquí el problema, pues con hacer clic a ese link "inofensivo" podemos entregar valiosa información a cualquier mal intencionado ahí afuera: la [dirección IP](https://es.wikipedia.org/wiki/Direcci%C3%B3n_IP) de nuestra computadora.

![image](https://www.muypro.com/wp-content/uploads/2009/07/Hacker-500x333.jpg)    
Mala combinación: "Información correcta, en manos inadecuadas".

En este punto probablemente se cuestionen, ¿qué demonios puede saber alguien a partir de un simple conjunto de dígitos separados por puntos?, Básicamente se pueden conocer los siguientes datos:

- País en el que se localiza la dirección IP.
- Región.
- Ciudad.
- Latitud/Longitud.
- [Nombre De Dominio](https://es.wikipedia.org/wiki/Dominio_de_Internet) asociado con la dirección IP.
- Nombre del [ISP](https://es.wikipedia.org/wiki/Proveedor_de_servicios_de_Internet).
- Velocidad de conexión.
Pero más importante que conocer esa información existe el peligro de que una persona con ciertas habilidades informáticas - y malas intenciones - intente conectarse a nuestra computadora y explotar cualquier vulnerabilidad del sistema operativo que utilizamos. Acotemos el problema a conexiones residenciales, pues las grandes corporaciones poseen - en el mejor de los casos - estructuras de seguridad muy díficiles de burlar.

La mayoría utiliza un [Router](https://es.wikipedia.org/wiki/Enrutador) para conectarse a internet, muchas veces viene acompañado de un [Firewall](https://es.wikipedia.org/wiki/Cortafuegos_%28inform%C3%A1tica%29). Esta configuración presenta una primera capa de seguridad ya que la direccion IP visible desde Internet, es la dirección del Router, con lo cual nuestra computadora no puede ser contactada directamente, y de paso, el Firewall se encargará de bloquear la mayoría de accesos no autorizados - la efectividad viene dada en gran medida por la configuracion del Firewall -. Hasta cierto punto, no estamos tan descubiertos como parece.

Luego de esta breve contextualización solamente queda explicar [como obtener la dirección IP de una computadora a través de un link](https://www.gohacking.com/2009/05/how-to-find-the-ip-address-of-a-remote-computer.html). La manera más sencilla es utilizar una aplicación Web, cuyo objetivo principal será capturar la dirección IP cuando alguna persona de click en un link determinado, las direcciones ip pueden ser recuperadas a través de un log manejado por la aplicación. El siguiente paso es rastrear la dirección IP, para ello se puede acceder al siguiente [sitio](https://www.ip2location.com/demo.aspx) especializado en proveer informacion geografica a partir de la dirección IP.

El uso que se le dé a esta información depende de cada persona, mMe limitaré a explicar el proceso para obtener una dirección IP y como rastrearla. Lo demás queda a su imaginación, y más importante aún, es el hecho que debemos tener mucho cuidado con lo que "cliqueamos". Como bien dicen : "la abstinencia es la mejor prevención". Sus comentarios son bienvenidos, hasta la próxima.