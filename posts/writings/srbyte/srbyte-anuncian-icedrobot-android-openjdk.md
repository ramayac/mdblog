---
title: Anuncian IcedRobot (Android + OpenJDK)
date: 2011-02-11
author: Rodrigo A.
tags: fosdem, openjdk, icedrobot, android
draft: false
post_id: blog-3515952828243908885.post-4773937781174860620
---

IcedRobot es el nombre que recibe un nuevo proyecto que apunta a ejecutar aplicaciones de Android en el OpenJDK ofreciendo otra alternativa a "Dalvik" (la maquina virtual de Android) y Apache Harmony.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiA3z3xR7bdIcsEwkad8LBDHyMzfO7P8yplWx1oxhaYlaPCwLYl2hMhJcRVmJeTjgqYPkxH5E9yGNB6mEQUKW3mpC-0hri2WO4DfUkS_loBmp4ZCecaJlIk5qeVPE-E4fiND-ynKv5OlkS6/s320/500px-FOSDEM.jpg)    
Free and Open Source Software Developers European Meeting El proyecto fue revelado al público en e[l FOSDEM 2011](https://wiki.debian.org/Java/DevJam/2011/Fosdem/JavaSpeakers#IcedRobot) por los developers David Fu y Mario Torre. Este (el proyecto) se separará en tres partes: GNUDroid creara una simplemente de Android utilizando software libre, el GNU Classpath y OpenJDK ( OpenJDK sera renombrado como IcedRobot Micro Edition para la plataforma Android). GNUBishop añadira funcionalidad extra para crear la IcedRobot Standard Edition. Y finalmente tendremos a "Daneel" que sería la VM contraparte de Dalvik.

![image](https://www.icedrobot.org/downloads/logo/robot.png)    
Logo de IcedRobot Por el momento, no hay código disponible al publico, pero Torre dice que se esta trabajando para "separar a Dalvik del kernel Linux personalizado que Android usa", lo que les permitiría ejecutar el código de forma separada e independiente (stand-alone). Luego, se esperan cambiar a Dalvik por la VM de OpenJDK y agregar un traductor para convertir el DEX bytecode de Android a Java ByteCode.

En pocas palabras (y espero que de forma más clara) van a hacer lo siguiente:

1. Cambiar el kernel de Android para separarlo Dalvik 2. Usar OpenJDK para ejecutar aplicaciones de Android 3. Implementar un "traductor" de aplicaciones de Android para que se ejecuten en OpenJDK. Es un proyecto que seguramente tomara al menos 6 meses para verlo funcionando. Igual, no podemos dejar de estar pendientes, ya que es una excelente noticia para los afortunados dueños de teléfonos Android y amantes de Software Libre :)

Existen más detalles en [el blog de Mario Torre](https://www.jroller.com/neugens/entry/introducing_icedrobot), y pueden encontrar la presentacion que dio en el FOSDEM 2011, en formato pdf, [dando click aqui](https://www.icedrobot.org/downloads/fosdem11/icedrobot-fosdem-2011-02-05.pdf).

Los dejo con un vídeo de FOSDEM 2011:

¡Saludos!