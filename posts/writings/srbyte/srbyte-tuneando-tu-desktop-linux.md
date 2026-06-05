---
title: Tuneando tu Desktop Linux
date: 2010-03-09
author: Robertux
tags: gonme, linux, debian, xfce, screenlet, compiz, gtk
draft: false
post_id: blog-3515952828243908885.post-7427918057855109338
---

Desde la década pasada y gracias a proyectos como Gnome, Kde, Compiz y Beryl, los escritorios Linux dejaron de ser conocidos por ser poco vistosos y nada amigables para el usuario y se han vuelto una seria competencia ante los gráficos que suelen verse en los escritorios con MacOSX y Windows 7.

Aun así, a algunos les cuesta tunear un poco sus escritorios y pasan todo el tiempo con su configuración por default, sin darse cuenta que con un par de paquetes, temas e imágenes que combinen pueden darle una mejor apariencia a este.

Para ello, elaboraremos una serie de posts explicando los pasos para decorar de diferentes maneras tu escritorio Linux. En este primer post haremos que nuestro escritorio luzca de la siguiente forma:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiMZxnWk9cK8TkC7FuXsxEO4Su129Kxb1E_b2FjYURJDDY0xfKL_P5ffIKGY7TtFJBA62LsmuES0TkRu0_Qv9SN9k5EPIRuUJTKevJpw21nafs_3zg9REKTX771R6gdPImn7Pg3sLeHq2o/s400/desktopShot.png)    
"Escritorio Linux ligeramente decorado"

Esta configuración la he elaborado sobre mi Debian GNU/Linux Lenny con el gestor XFCE. Esta también puede ser aplicada sobre cualquier otra distro basada en Debian como por ej. Ubuntu o Linux Mint, además tambien aplica sobre escritorios Gnome, con algunos ligeros cambios.

Para empezar, descargamos la imagen de fondo de [esta url](https://www.ewallpapers.eu/view_wallpaper/springtime-1280-960-40.html) y lo aplicamos como fondo de pantalla en la ventana DesktopSettings (cabe mencionar que se puede acceder a los settings de Xfce haciendo clic derecho sobre el escritorio):

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEg_WxCx_U2xcqAy_X_OfBNDivk95bDdm1cGzD0eQLeiNABNWx76_On9_Aguo_n5ESZP-o1zH_pL9NZx3-_0dfm82dHfBTp54RTrFL72AOoewGkVJIqSZYHg_TcPys5aRyJJCvSfGaw9Tl4/s400/xfceDesktopSettings.png)    
"Ventana Desktop Settings del Xfce, con el fondo de pantalla SpringTime."

Personalmente no me gustan los paneles o barras de botones que vienen en los sistemas operativos asi que los elimino o dejo uno solo en la parte superior con la opcion auto-ocultable y transparencia. Esto lo definimos en la ventana Panel Settings:
![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEg3EP4_uf36oSwqLdLX2T-DUd-yYjlplgJpiFhouPuDcwvD3Yt8KsspjI4yJIXWkGZSMdZjuNd_GTcnwpFYUrnWFgzshAOC5gciD8B3g6JrK_tMHCD64C1c6wCA4QszFOgKwLF544WNjuY/s400/xfcePanelSettings.png)    
"Ventana Panel Settings con autohide activado y un unico panel en la parte superior"

Luego, procedemos a instalar los temas de la interfaz gráfica y los bordes de ventanas. Estos se llaman Axiom y al igual que muchos otros temas, se encuentran en [xfce-look.org](https://xfce-look.org/) para escritorios Xfce y [gnome-look.org](https://gnome-look.org/) para escritorios Gnome. Los temas Axiom para la interfaz y bordes de ventana se encuentran en [esta](https://xfce-look.org/content/show.php/axiom+gtk?content=90873) y [esta](https://xfce-look.org/content/show.php/axiom+xfwm?content=90145) URL respectivamente.

Luego de descargarlos, los descomprimimos y como usuario root colocamos ambas carpetas en la ruta /usr/share/themes/ para después aplicarlos en las ventanas de User Interface Settings y Window Manager Settings:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEi6Bw7Hy3x7SoKtSJ-MSw07ZqqGxX75L9C_fFE-lB8fuNwenM-xWZUXQtmhJ0vXu38WBv0WN8-JoQBQPkZdtEJPy8Vkt_UigX5QjRoWewZYMJMkgNY2pTnRkr0eRoH9e-WwF0fZfZBYIZI/s400/xfceUserInterfaceSettings.png)    
"Al agregar nuestro theme a la ruta /usr/share/themes/ este aparece automaticámente en las User Interface Settings y nos es posible seleccionarlo"

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjAt9Fr_PDxo6Q_lW6QM0v_8mCZ0xmqiXr6masW2NJL1YUn-0g383Bz2tOqIhyphenhyphenLIAFQXPmQSFrQX9bshUxsIeSkzjeZloH4sSeiYxvrCZgXMSc2ypC908mwgla1xL-FlJp87eiTD2YaYQ0/s400/xfceWindowManagerSettings.png)    
"Aplicando el borde de ventanas Axiom desde el Window Manager Settings"

Xfce tiene la bondad de permitirnos aplicar transparencias y sombras sobre las ventanas sin necesidad de compiz o alguna otra libreria especial para gráficos. Esto lo aplicamos desde la ventana Window Manager Tweaks:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjkUiVegx9l19iKHH9gKJ6MDD179KbTduD38Fwstwk62AdhjGbKsqoBTdXFctj_KS71AHJKa429iBlZCL2ThCdi0vVoGzHkDZ7wSOVadO8L3ujLLM75VZBC7RH_xeA3VXFldksfeXKn554/s400/xfceWindowManagerTweaks.png)    
"Ventana Window Manager Tweaks en la cual definimos transparencias para las ventanas inactivas y sombras sobre los bordes de las ventanas"

Los elementos que ven en el escritorio tales como el calendario, reloj, gmail, etc son nada mas y nada menos que [screenlets](https://www.screenlets.org/). Estos funcionan exactamente de la misma manera que los Vista Gadgets y de hecho, el concepto de gadgets adheridos al escritorio existía desde hace mucho antes en Linux via gdesklets para gnome. Los screenlets vienen en los repositorios de debian y ubuntu por lo que vasta hacer un apt-get install screenlets como usuario root para instalarlos. Además también se pueden descargar del sitio oficial. En la ventana Screenlets Manager se pueden escoger, iniciar y configurar los screenlets deseados:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhYCxm3AazG4j4x7q1AV3h3aCZznw1k4JM4hobPFpVsqLbiIjFx368yJFqXao0FsJYGAfx7m3kd3I2ABpXhVbB_Ere07YlFB8VPPClMmQ8Oi8M3DT2X5L93RM6C-Cxrs9cQ6RgVUUhnhyc/s400/screenletsMainWindow2.png)    
"Ventana ScreenletsManager mostrando los screenlets disponibles y sus opciones"

Todos los screenlets que se ven en la captura de pantalla al inicio del post vienen por defecto con el paquete screenlets por lo que no es necesario descargar nada mas y si aun así desean mas, estos pueden descargarse del sitio de [gnome-look](https://gnome-look.org/index.php?xcontentmode=6700&PHPSESSID=140e052c44e85f9d97e0fc8ff2604a99) y basta con descomprimirlos y copiarlos a la carpeta /home/tuUsuario/.screenlets/ para poder verlos en la ventana Screenlets Manager.

Los screenlets utilizados fueron:

- Battery
- Clear calendar
- Clock
- Gmail
- Launcher
- Now Playing
- Volume Control
- Window List

Todos los screenlets llevan sus temas por defecto, no se utilizó ningun tema especial y estos han tomado la transparencia que por defecto se le definió a las ventanas inactivas en la ventana Window Manager Tweaks. El único screenlet que requirió un poco mas de trabajo fue el botón de shut down que ven en la esquina inferior derecha. Este es un screenlet de tipo launcher con una imagen personalizada la cual pueden descargar de [esta](https://dryicons.com/images/icon_sets/colorful_stickers_part_4_icons_set/png/256x256/shut_down.png) [url](https://dryicons.com/images/icon_sets/colorful_stickers_part_4_icons_set/png/256x256/shut_down.png). Para que al presionar el botón aparezcan las opciones de apagar, hibernar, suspender, etc. se creo un pequeño archivo logout.sh cuyo contenido pueden ver en [esta url](https://srbyte.pastebin.com/vdkBrgda). Luego, en los settings del screenlet agregan la ruta a este archivo .sh para que el botón funcione y listo.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEh9oZ7gXmIxQaSrlC4eXyBsfsT8kzFN-JB3-jlBBaNVS50bJ02fi8w_f8-MJUEjt0lxoqxE-qIOevrVveMoLxEvGExQD5yW5uZT1cqg8hq4nBqB1sZX56mdXcDAHzujNnzea5BGvJWw_ZM/s400/screenletLauncherSettingsWindow.png)    

Siguiendo estos pasos pueden tener un escritorio Xfce similar al primer screenshot de este post. Espero sus comentarios para saber si lo pudieron llevar a cabo. Saludos!