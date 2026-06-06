---
title: OK Ya Instale Linux. Y ahora que? (Parte II)
date: 2008-11-28
author: Rodrigo A.
tags: gmail, herramienta, ubuntu, gadget, linux, recomendado, debian, google, tutorial, screenlet, gnome, gnu
draft: false
post_id: blog-3515952828243908885.post-8532126731644485549
---

Una vez que hemos puesto nuestro tema favorito, es hora de instalar un par de accesorios para tener todo a la mano y poder ver la información que nos interesa sin necesidad de dar muchos clics o gastar memoria abriendo programas pesados para simplemente chequear si hay un nuevo correo, ver la fecha actual o bajarle el volumen general a la computadora.

Estos accesorios los podemos agregar al panel principal (nos basaremos en el gnome-panel, standard del escritorio gnome, el cual viene incluido por defecto en el Ubuntu y Debian, distribuciones populares de Linux) mediante muchos programas aunque aca nos basaremos en tres diferentes accesorios: el Gnome Panel y los Screenlets.

Para agregar un nuevo elemento al panel de GNome simplemente hacemos clic derecho sobre el y seleccionamos la opción "Agregar al Panel". De la ventana que aparece seleccionamos los applets de nuestra elección. Algunos de los applets recomendados son el "Forzar la Salida" el cual permite cerrar una aplicación que deje de responder simplemente haciendo clic sobre ella, "Indicador de Distribución de Teclado" para los que tenemos teclado inglés y toca estar cambiando a distribucion latinoamericana y el "Deskbar", el cual es un buscador todo en uno.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjDxGr54pLTrgypXFlxbxoQcVOPidMBqCFRo05x-5izoaA0fszhSKLKTmRwTwuB_0rI6XWhRNGTDkStaprrEKkzgd-xT4U3K_tPH1eNHGaUXnt2fyPmAkjhrb8O9sjIfvP81EAFR2SaxQs/s400/GnomePanelShot.png)    
"Imagen del Gnome Panel mostrando algunos applets como el indicador del teclado, la fecha/hora, el deskbar, forzar salida y las estadísticas del CPU/RAM/Red. Mostrando también la ventana para agregar mas applets"

Otra muy conveniente forma de tener accesorios en nuestro escritorio es usando los Screenlets. Estos accesorios son tan útiles como variados y fáciles de administrar. No se requiere mas que hacer un apt-get install screenlets para tenerlos a nuestra disposicion. Al final, podemos hacer que nuestro escritorio luzca de esta forma (clic para agrandar):

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjZawRPODt0qAKtE5PA3A371lfd3CA6jqIUJmZKW0zlHW0QNmcg2CtRrxOQRiXQUlchyphenhyphenHf0l2WBk6dirGXQfvn8UxY8GQ3AXQOyMWR1US25Djg2Ua0D0XVoK__P6LH7bOrYyAIStOtcLQ4/s400/600px-Mydesk.jpg)    
"Escritorio de Gnome haciendo uso de varios screenlets como el ClockRings en el fondo del escritorio, Reloj, Calendario, CPU Stats, etc."

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhq7TtuU9UwfQPh7BWdZ5ynvqOH2UcjaxlIjXzD-sCvI3SH1dw8jyWuPXpVQim23G34rZ2xgA2ZVIJ2PmJL_I85PChKvLi69OchnNeaQT7TES-8DRba9L7n1VP41INmmJBsO6YKoKNXuB8/s400/linux+vista.jpg)    
"Otra serie de Screenlets los cuales junto con el tema y el fondo de escritorio dan la apariencia de Vista Gadgets en un Widows Vista."

La mayoria de screenlets no vienen por defecto al instalar la aplicación desde apt-get así que los podremos descargar de [gnome-look](https://gnome-look.org/index.php?xcontentmode=6700) o de la [página oficial del proyecto Screenlets](https://screenlets.org/index.php/Category:UserScreenlets) en la cual tambien te indican las instrucciones de instalación y uso. Entre los Screenlets recomendados se encuentran el [GMail Checker](https://screenlets.org/index.php/Gmail), [CPU Meter](https://screenlets.org/index.php/CPU_Meter_Vista%27ish), [Now Calendar](https://screenlets.org/index.php/Now_Calendar_Screenlet) (con capacidades de consumir eventos de calendarios como Google Calendar, via iCal), el [Sidebar](https://screenlets.org/index.php/Sidebar) y el [Circle Clock](https://screenlets.org/index.php/CircleClock), [Wireless Screenlet](https://screenlets.org/index.php/WirelessScreenlet2).

Para instalar estos Screenlets, te descargas el archivo comprimido de cada uno de los vínculos anteriores, lo descomprimes y copias la carpeta descomprimida en la carpeta /usr/share/screenlets/ o en la carpeta .screenlets/ luego abres el Screenlets Manager (Menu Aplicaciones -> Accesorios -> Screenlets) y los verás listados en dicha ventana para que nada mas los actives y los ubiques en la posición que quieras dentro de tu escritorio.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEis3k8RAuPpUkIZVKJgSlILWFocwlmq75c9zojeEHyH7AWTc4NkcAU2xs-XsUVRm6XgsTGWT9TclP4FXcta0JrC47uO6GZRzech7Mvd9QsQJGOON0usW43sJtLuygKeHINIZe36c7C7Vbk/s400/ScreenletsManager.jpg)    
"Apariencia de la ventana del Screenlets Manager"