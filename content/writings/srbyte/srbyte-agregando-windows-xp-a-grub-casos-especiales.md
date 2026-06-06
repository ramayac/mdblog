---
title: Agregando Windows XP a GRUB... casos especiales.
date: 2010-01-11
author: Rodrigo A.
tags: grub, windows, map
draft: false
post_id: blog-3515952828243908885.post-7390754271798906381
---

Nota rápida:

Recientemente me vi en la "obligación" de instalar un Windows XP a una maquina, que por dos años solo tuvo Ubuntu Linux. La maquina ya tiene dos discos duros, así que añadí un tercero que tenia tirado, y en ese instale el Windows, para no cambiar esquemas de particiones ni re-dimensionar absolutamente nada. La tarea sería sencilla:

1. Desconectar todos los discos. 2. Conectar el nuevo disco. 3. Instalar Windows en el disco conectado. 4. Añadir los discos desconectados. Dejando a Windows en el segundo disco esclavo. 5. Añadir a Grub la entrada de Windows XP (en el menu.lst). El problema radica es que a Windows XP no le gusta arrancar en dos situaciones:

- Que este alojado en una partición extendida (es decir, una partición secundaria).
- En un disco duro que no sea el primario.

Creo que esas limitantes, se merecen esta imagen:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEii4HGvz-Qu6eUk7bgpKmAEEmhGkWiKA8kCCvVyel08fO3iCwLrOZXD5Ug_Zlvx8dI-eENLFE2mc5SuUlQQFegnB6u5ghdrYD-d2nLeEYD28ABhVIoJKzNkma42z_qqScCSBDqk3D8RpCjj/s200/microsoft_lazy.jpg)    

Detalles, detalles, para no darle muchas vueltas al asunto, la basura no arrancaba. Así que hice el "truco legendario" de mapear los dispositivos en el GRUB, más no sabía que había cambiado ligeramente la sintaxis del comando map (irónicamente en la ayuda [este cambio no se menciona](https://www.gnu.org/software/grub/manual/html_node/map.html)).

En fin... para los que tienen un disco duro extra con windows xp, y quieren que arranque utilizando GRUB 2 (1.97 beta), si no saben como se hace, pues acá les dejo la sencilla solución:

```
title Winbugs map (hd0,0) (hd1,0) map (hd1,0) (hd0,0) rootnoverify (hd1,0) makeactive savedefault chainloader +1
```

En la "nueva" (para mi al menos, no tengo ni la mas remota idea de hace cuanto lo cambiaron) sintaxis  de map, se especifica el disco físico y su partición. Antes, era solo el disco, y por eso maldije mucho el día de ayer. Aclarando: esta configuración es para el caso de que se agrega un disco duro como esclavo, con una partición con Windows XP. ¡Espero que les sirva, saludos!