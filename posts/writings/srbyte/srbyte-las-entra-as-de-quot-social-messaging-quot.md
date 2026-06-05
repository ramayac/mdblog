---
title: Las entrañas de &quot;Social Messaging&quot;
date: 2010-11-18
author: Rodrigo A.
tags: facebook, hadoop, hbase, messaging
draft: false
post_id: blog-3515952828243908885.post-5311978265864445715
---

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEg-oqDOfQAKQBp4G46KueKuLv0STqV3dyuFxvcnKn7Zi5QQ4QZV-LH_LNoIyn3H-NxIGnodvvEylgy1yN2PGBhyMib98joeNosWF6EH4u9CYb0j_MYEcC5hGlpeqwzIua9jIGmNJor3uBLX/s200/fb-messages-225.jpg)    
Hace apenas tres días (15 Nov 2010), Facebook lanzo (por invitación) su nuevo servicio de "Social Messaging System", un servicio innovador y que inclusive nos proveerá con una dirección "@facebook.com" en su debido tiempo.

Social Messaging (solo Messaging de ahora en adelante) integra SMS (mensajes de celular), chat (de Facebook) y correo electrónico en una interfaz transparente, que gira en torno a la plataforma "Chat (de Facebook)" rediseñada para loas nuevos retos que Messaging presenta.

Messaging, más que una competencia al correo electrónico "tradicional", parece más el próximo paso lógico de Facebook para "acaparar el mercado de la comunicación escrita", sin cambiarte de pestaña, ventana o dispositivo.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhTBLn-pj0CdwxOoWEZPbdZOHuxX3twmtdBl1fVaWHKGrh4S6E5Tdaa0_NeBWqFt9C8qiGzxy-Bu_7dygvcY6Z4-iQel8kt7tiKZKDsNOVF9VthnrPV5Yw9c8abrk9WWlbcOhbltZ8HEVmh/s400/fb-messages.jpg)    

Todo esta en un solo lugar (el poder del "default"), uno solo busca a al contacto a la que le quiere enviar un mensaje, le dispara caracteres e independientemente del dispositivo o tecnología que use esta persona (SMS, email), el o ella recibirá el mensaje, y cuando te contesten, vas a recibir las respuestas como "una conversación" y siempre todo, en el mismo lugar: Facebook.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgOyT1XrZtEx2AnGfG0O1V8dkHwX7-l3IMBlhRqbPfdHTpEy8aMRkATmC1k_RfcpT-7yRpJD3yqjxMSWiiznlkBEGTgoYgzUwjwvdjmegcx89TdpJlJAf-4-c6VntIdbwQrehyphenhyphenqtGXToVRv/s400/ben-sam-texting-620.jpg)    

Pueda que a ustedes no le guste Facebook o sus políticas de privacidad, pero ignorando estas (políticas) y siendo totalmente objetivos, NO se puede negar que con Messaging:

- Facebook ganara mucho más uso y "adeptos" al centralizar la interfaz para interactuar con los contactos que se tengan agregados a una cuenta.
- Los retos para diseñar y producir una solución informática así no son nada sencillos.

Lo que a mi más me interesa es saber como todo esto funciona detrás de la pantalla blanca con azul, quiero saber como va a funcionar y que va a utilizar internamente Messaging, y a precisamente eso es a lo que vamos...

Entrañas digitales

Actualmente la infraestructura de mensajería de Facebook (Chat e Inbox) maneja más de 350 millones de usuarios enviando 15,000 millones (billones para los gringos anglopartlantes) de mensajes de persona a persona al mes. Cuando los ingenieros de Facebook monitorearon el uso de esta infraestructura, encontraron dos patrones (de uso):

1. Datos temporales y volátiles (Chat) 2. Datos que no dejan de "crecer" y de bajo acceso (Inbox).

El reto para los Ingenieros de Facebook, era elegir un sistema de almacenamiento (una base de datos) que sirviera para almacenar la información que se diluye de estos dos patrones de uso detectados. Ya que Messaging es una combinación de concepto y diseño de Inbox y Chat, nada tiene más sentido que buscar dicha solución. Así que lo que hicieron por "un par de semanas" fue hacer pruebas de rendimiento a diversos motores de base de datos: MySQL, Cassandra, HBase y otros. El ganador eventual fue [HBase](https://hbase.apache.org/)
![image](https://upload.wikimedia.org/wikipedia/en/e/e7/HBase_Logo.png)    

HBase es una base de datos distribuida, versionable, orientada a columnas clave-valor, open source e inspirada en el modelo de [BigTable](https://en.wikipedia.org/wiki/BigTable) (de Google). HBase se puede utilizar en escenarios en donde se necesita escritura/lectura aleatoria en tiempo real de muchos datos. Cuando hablamos de "muchos datos", hablamos por ejemplo de:

> "Mil millones de filas multiplicado por millones de columnas"
Como ya imaginan, son cantidades enormes de información. Y todo esto ejecutándose en agrupaciones de servidores baratos o "comodity hardware clusters" como se les conoce, para mantener bajos los costos. Algunas de las características que impulsaron a Facebook a elegir HBase son:

- Modelo de consistencia más simple que el de Cassandra.
- Buena escalabilidad y rendimiento para los patrones de datos de Facebook.
- Riqueza de características: auto balanceo de carga, failover, soporte de compresión de datos, múltiples "shards" (fragmentos) por servidor, etc.
- Utiliza [HDFS](https://en.wikipedia.org/wiki/Hadoop#Filesystems) (que Facebook ya usa con los sistemas que usan Hadoop), el sistema de ficheros distribuido con soporte para replicarse, validaciones de checksum end-to-end y balanceo de datos automático.
¡Y eso, es solo la base de datos! Messaging tiene varios aspectos interesantes y claves que lo hacen un proyecto sumamente interesante, por ejemplo:

- Los archivos adjuntos se almacenaran en [HayStack](https://www.facebook.com/note.php?note_id=76191543919)
- Están usando un servidor de aplicaciones escrito desde cero
- Usan [ZooKeeper](https://hadoop.apache.org/zookeeper/) para los servicios de "Discovery".
- Integración con servicios de "email account verificacion", relación de amistad, decisiones de privacidad y decisiones de envío (¿un mensaje se deberá enviar por chat o SMS?)
Y finalmente, el detalle que más me sorprende es que Facebook va a liberar 20 servicios nuevos, que utilizará Messaging hechos por solo [15 Ingenieros](https://www.blogger.com/goog_1857599340) [en tan solo un año](https://www.theregister.co.uk/2010/11/15/facebooks_largest_ever_engineering_project/).

Con la elección de HBase como capa de almacenamiento, este proyecto esta "ungido" por Facebook, y que después de leer un poco sobre el funcionamiento interno de Messaging, me atrevería decir que sí estoy esperando ver pronto los resultados de tan interesante proyecto en mi cuenta de Facebook, y en especial ¡las estadísticas de uso! :)

Si les interesa el tema, les recomiendo estos tres sitios donde pueden encontrar más información sobre este interesante proyecto: [The Underlying Technology of Messages](https://www.facebook.com/note.php?note_id=454991608919), [Facebook's New Real-time Messaging System: HBase to Store 135+ Billion Messages a Month](https://highscalability.com/blog/2010/11/16/facebooks-new-real-time-messaging-system-hbase-to-store-135.html) y [Facebook Messages Walkthrough Pics](https://facebook%20messages%20walkthrough%20pics/).

¡Saludos!