---
title: Adaptando tu Sitio Para Móviles - Parte 1
date: 2010-01-26
author: Rodrigo A.
tags: navegador, wifi, movil, web, ipod, iphone
draft: false
post_id: blog-3515952828243908885.post-6872410106786116984
---

Para nadie es sorpresa que cada vez son más las personas que poseen un smartphone o un dispositivo (no necesariamente un celular) con capacidades de navegación por la World Wide Web como por ejemplo un iPhone o un iPod touch. Estos dispositivos son muy prácticos para cuando deseas revisar tu correo electrónico, los feeds de tus blogs o simplemente entretenerte navegando en el mar de páginas web que se pueden encontrar hoy en día y de casualidad te encuentras fuera de tu casa, en alguna cafetería o centro comercial donde encontraste una red WiFi disponible.

Sabiendo esto, las empresas y organizaciones deben tomar en cuenta que muchos de sus clientes / usuarios ahora tienen la alternativa de acceder desde un dispositivo móvil a los servicios web (o simplemente el contenido web) que estos ofrecen pero existe un inconveniente: Dichos sitios web fueron diseñados para ser vistos desde una computadora con una resolución como mínimo de 1024x768. Esto significa que cuando intentas ver uno de estos sitios en tu dispositivo móvil, puede lucir como en la siguiente imagen:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEib09uHQMqPjoFJcrbQgoiNCwgS6xZ5f9zOHCw2wEiFgdBQlwsGyk-FV1fuc28ExiIX5O-a7OMs7iRagUtaomubTrr5vqSK3Bzn5uQ095tBJrf5MZKbMwojv5uNSCsV_vvfdC6DksLyPh4/s400/google-code-ipod.jpg)    
"Google Code - Uno de los sitios de Google que a diferencia de GMail, GReader y Google Calendar, aun no tiene una version para móviles. Screenshot tomado desde un iPod touch (izquierda) y desde un móvil con Android (derecha). Captura de pantalla desde el Android cortesia de [@jainux](https://twitter.com/jainux)
"

Como puedes notar, el sitio no está adecuado para ser visualizado desde un dispositivo móvil. Esto dificulta la navegación en el mismo, teniendo que hacer zoom in para poder leer el texto y también zoom out para luego ver la página completa y dirigirse a otras secciones de esta.

El primer framework que conocí, el cual realmente es un plugin para Wordpress, fué [WPTouch](https://www.bravenewcode.com/wptouch/). Al agregar este plugin a tu blog de Wordpress, este identifica automáticamente cuando estás visitando el blog desde un iPhone o iPod touch y modifica la plantilla adaptándolo a las dimensiones y paleta de colores de dicho dispositivo para que la lectura del blog sea lo más adecuada.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhKgn4NfvCvSgJ72tkwKCTCdo1oCNQofAjUf8gutm6z5vfu612ge4o0pLFdHndbkOGm2mzj6cdM1IFx9stmDOpkumVt_tJMb28AFpvcWuapCMjWA8723k8bYvyur8JYPV0Bn4HgkUlZ2WA/s400/wptouch-blog.jpg)    
"[Soda Caustica](https://sicotropico.wordpress.com/), blog de Virginia Lemus. Uno de los primeros blogs donde observe el uso del plugin WPTouch para blogs de Wordpress."

Existen servicios que te automatizan la tarea de crear una versión móvil de tu sitio, adaptando las dimensiones de la plantilla y eliminando contenido que posiblemente no funcionará en tu dispositivo o retrasará innecesariamente la carga del mismo. Entre estos sitios se encuentran [MoFuse,](https://www.mofusepremium.com/) [dotMobi](https://mtld.mobi/) y [mobifyMe](https://mobify.me/).

MoFuse ([link](https://www.mofusepremium.com/)) Este es el más sencillo de todos los servicios. Te permite crear una versión móvil en tan solo un par de clicks con un mínimo de personalización del mismo pero ofreciéndote a cambio un dominio .mofuse.com para tu sitio creado el cual luego puedes asociar con tu propio dominio. Ofrece diferentes planes de pago y entre ellos, un plan gratuito orientado a blogs. Para empezar, debes escribir la URL de tu blog o de los feeds generados a partir de tu blog.

DotMobi ([link](https://mtld.mobi/)) Es en forma general, un conjunto de servicios orientados a promover los sitios móviles y la movilización de sitios actuales, así como todo lo relacionado con sitios web para móviles, ofreciéndote muy valiosas herramientas, documentación y referencias. Entre estas herramientas posee una denominada [Instant Mobilizer](https://instantmobilizer.com/), la cual genera casi automáticamente una versión móvil de tu sitio web. Las features que más llaman la atencion son Click to call que agrega un vínculo con el número telefónico de tu empresa y Google Map Autogeneration que agrega un vínculo de Google Maps apuntando a la dirección de tu empresa. La desventaja es que no posee un dominio propio para asignarle a tus sitios sino que tienes que conseguir un dominio con uno de sus partners.

MobifyMe ([link](https://mobify.me/)) También te permite crear una versión móvil de tu sitio pero a diferencia de DotMobi y MoFuse, este no te lo hace de forma instantánea generando una plantilla automáticamente. Sacrificas esta feature obteniendo un mayor control sobre el contenido de tu nuevo sitio permitiéndote seleccionar los elementos que deseas mostrar tomándolos directamente de tu sitio original y adaptándolos a las dimensiones de un dispositivo móvil usando además hojas de estilo personalizadas.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhroQRtpa71CXL9cpdawoUkcozq7iDs-lyEcNBh04HLsyHUaPz4KAAaHniKtYtEYD4hj7f6PXXlrvlAMCf_Mpt9RMqo7FrqH9WcG1n-9dLKe7NINp2rBCx9SO7x93B8jNXOgw6hF6bh6Xo/s400/mobify+-+editing.png)    
"Editando el layout una página de ejemplo en mobify.me"

Estos tres servicios te generan un sitio con su propio hosting y un domain customizable (también asociable a tu propio dominio) además de ofrecerte plugins y scripts para autodetección de tipos de dispositivos, el cual, al agregarlo a tu sitio original, puedes hacer que te redireccione automáticamente a tu versión móvil cuando este detecta que el usuario está navegando desde uno de estos dispositivos. Aunque parezca que este no es tan sencillo como los otros dos, este es el más profesional de todos y el que más opciones te ofrece para la personalización de tu nuevo sitio móvil. Además, entre los planes de pago te ofrece un plan gratuito con el que puedes hacer más que suficiente para montar la versión móvil de un sitio sencillo o un blog.

[Acá](https://srbyte.mofuse.mobi/) puedes ver una versión del sr byte generada com MoFuse y [acá](https://srbyte.mobify.me/) una generada con MobifyMe. Esta última es la que mejor me ha parecido así que esperamos pronto se convertirá en la versión oficial del Sr Byte para móviles. Pronto habilitaremos su respectivo subdominio y la redirección automática.

Debido a la gran cantidad de contenido que deseaba postear sobre este tema, he dividido el post en dos partes. En el siguiente post continuaré con la segunda parte: "Adaptando tu Sitio Para Móviles - Parte 2: Do It Yourself"