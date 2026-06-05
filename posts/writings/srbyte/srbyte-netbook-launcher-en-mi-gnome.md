---
title: Netbook Launcher en mi GNOME
date: 2010-04-17
author: Rodrigo A.
tags: netbook, ubuntu, gnome-do, launcher, gnome, remix
draft: false
post_id: blog-3515952828243908885.post-3222536249714833010
---

Hace un tiempo, tuve la oportunidad de instalar [Ubuntu Netbook Remix](https://www.canonical.com/projects/ubuntu/unr)(UNR) en una netbook, y aunque me gusto la interfaz, mi impresión personal sobre la interfaz general del UNR es que es vistos, pero lento al lanzar las aplicaciones. Pero como aun así me había gustado, decidí instalarlo y utilizarlo en mi desktop. Básicamente, lo que quería, es tener el "feeling" del UNR en mi escritorio GNOME, y esto fue una tarea bastante fácil de realizar.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgnj3VZNfVE9CYDzq1mMvwkyayp5uReXe0rk-Aeaq45ZlG3GwswysapXDwyhWqBGJZ0HDH7wV7D9xyChSW5DB4bFhhvnI4yRL1QQvHa-KFHpn_SUafE0Y7qAyCrQyP6DlUaYO8B39ubJbqA/s320/unr-favorites-small.png)    

La interfaz de UNR esta compuesta por tres programas integrados entre si:

- Maximus: un "daemon" que maximiza las ventanas para ahorrar espacio.
- Netbook Launcher: El lanzador de aplicaciones, que más parece un desplegado en el escritorio, con varias conveniencias.
- Y applets "especiales" en el panel de Gnome (el orden de los applets es: go-home-applet | window-picker-applet | notification-area-applet | mixer-applet | clock ).
En si el único programa de mi interés es el Netbook Launcher, que fácilmente se instala en Ubuntu 9.10:

#apt-get install netbook-launcher

Sobre el Netbook Launcher puedo decir que para ser un programa tan sencillo a veces se siente como un producto sin terminar. No se puede re-ordenar los iconos, ni aun en la seccion de favoritos. Y algunas carpetas de red, o dispositivos USB, no se montaran o desmontaran como uno espera, sino que hay que ir a Nautilus o a la terminal para desmontarlas. Tampoco puedo hacer un quick search sobre el área de iconos desplegados (como cuando se busca archivos en Thunar o Nautilus). Sin embargo, la conveniencia de la sección de "Favoritos" es suficiente (para mi) para obviar cualquier problema mencionado, y es más, hasta quitar cualquier otro tipo de lanzado en el panel de GNOME.

El Netbook Launcher jamas podrá reemplazar a Gnome-Do, en velocidad y extensibilidad - o inclusive a "dmenu" - pero ofrece una amena alternativa en cuanto a concepto y a funcionalidad de los lanzadores de aplicaciones tradicionales.

Ah, ¿y el resultado de la instalacion? una locura:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgH6M72YytkZb-_IC9ys-a6npf9CG-FrYoRwpak7iRCpclctVWUHW7ijz6_C5HU_ZgcX_p9Ov4w7A9Q6nc2zjlVCncjfGrrDsjE8sygqGtDM3YoiHyRzX18Zr0VD2iESi2ZKHSOkGMDj9Wh/s320/ramayacDesktopNetbookLauncher.png)    
"Sección Favoritos Netbook Launcher"

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjeEud3w1H9y3NH3aDt1cMRjUrw6yfe4WH1vt5vlvh6dR6mfW4W2SvDvpWEQTyLdXAExFeXMeFYK6yKgtG9KtEKxcadOVV-IcD2K9-hTSenEQChA_bbLhFiGUv7f_7dCpCS7nqnqaNpWCxk/s320/NetbookLauncherCompiz.png)    
"Netbook Launcher & Compiz Fusion"

La integración del Netbook Launcher con Compiz Fusion es transparente, y en general la experiencia es integrada y placentera. Realmente deseo usar la interfaz de Netbook Launcher con un monitor touch, creo que así le sacaría más el jugo a la interfaz. Sinceramente creo que es una curiosidad que vale la pena probar, así que para los aventureros que cambian constantemente su interfaz, y no saben a ciencia cierta que poner o quitar en su escritorio, si se animan esta es una funcional alternativa. El siguiente paso, creo que sería instalar [Gnome-Shell](https://live.gnome.org/GnomeShell) junto al Netbook Launcher, ¿no? ¡Saludos!