---
title: Archivos &quot;Pitufo&quot; atacan mi SVN
date: 2010-02-07
author: Rodrigo A.
tags: archivos, svn, cmd, windows, azul
draft: false
post_id: blog-3515952828243908885.post-6302627561200427818
---

Hace unos días quería subir un cambio al [repositorio de código SVN](https://www.srbyte.com/2008/03/programemos-mejor-subversion.html), para añadir información extra de depuración a un Servlet para monitorear un extraño bug. Intentar enviar el cambio resultaba en un error en mi plugin de SVN (SubEclipse), que al principio confundí con un atributo "Lock" sobre el (archivo). Acepto que no leí en detalle el error y me deje ir por la primera impresión. Más después de unos momentos, comencé a notar varias irregularidades en los archivos "base"...

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgXiisVD_Io0lTjJhQCwNqg17HHz8U-o_kvcqzA-oRtRBqNkT0fF7N3qFMxVcnJGAkPVF89EzX9UwCA_cfvJ1cq1mFfCWGm9Qps0UyW-6bzMASEmb3ZJZDeF-sOoEnLHVn7xjfy0HtpGzfk/s320/svnERROR.png)    

Dos ideas vienen a la mente cuando uno ve una imagen como la anterior:

1. El archivo en el servidor esta corrupto. 2. Mi copia local esta corrupta. Respalde aceleradamente mis archivos con cambios a un lugar seguro, y procedo a hacer un "Override and Update". La irregularidad se mantiene. Accedo al servidor SVN por su interfaz web, con la esperanza de encontrar el problema en el servidor... ¡¡¡pero no el fregado archivo esta bien!!!.

Bien, parece que el problema es local... así que me dispongo a ver las infames carpetas ".svn", y a medida que me adentro en la jerarquia de subcarpetas del proyecto, comienzo a verlo atestado de "pequeños pitufitos archivitos azules", y cuando ingreso a la carpeta "text-base" (en .svn de la carpeta que aloja mis archivos modificados) me encuentro nuevamente con los infelices:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEi21XrbeKIRaPv_GfpAdEiRHas6CNze639qiaHqtSbXjMgrFWVkkaA275v37Qe033z9Nd761MMx_CJwpBLsMtz7kHHIbMHISnH-xlB12zk-Nk100TNldGca9lO0gEchhrGUEMrvyoO38U0v/s320/ohcrap!bluefilessvn.png)    

Leyendo más a detalle el error del SubEclipse este describe un "mismatch" de la firma de los archivos. Ah, ahora entiendo el problema. Windows XP tiene una característica que consiste en comprimir los "archivos antiguos" (que no cambian en 60 dias) para ahorrar espacio

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiY3bdl9FYUnAt-EComq_yXVcx11cJHSQDNy3evhIXe-YCVTPSjemu6d0uObZP7jHLXmoRgAQUmb93IV4E4nzBng4HCNh9VQl4z91nToaChbRdJYUXoEVK-oBE-YeY_RFufGhXWBHhhUnko/s320/comprimir.png)    

El explorador de archivos de Windows, "colorea" de azul los archivos que fueron comprimidos de esta forma. Ciertamente la caracteristica es util para ahorrar espacio, y transparente para muchisimos programas, pero no es bien recibido por los clientes de SVN.

Los clientes de SVN ([SubEclipse](https://subclipse.tigris.org/), [Tortoise](https://tortoisesvn.tigris.org/), etc) mantienen un registro de [firmas (hash) de los archivos](https://www.srbyte.com/2007/09/encriptar-y-hashing.html) base (los originales que se obtienen desde el servidor). Pero como Windows XP comprimió los archivos, la firma de archivos modificados cambió, y el SubEclipse protesto mucho por el "mismatch" de la firma de los archivos afectados, por esa simple razón no se podía enviar el cambio a fin de cuentas.

¿Qué hacer en este caso? Lo primero que se nos viene a la mente es hacer un CheckOut del proyecto, y no es una mala idea, pero si tienes varios proyectos (ocho por ejemplo) afectados por los "archivos azules", antes de ponerte a insultar la cuenta de [@BillGates](https://twitter.com/BillGates) usando Twitter, puedes usar el comando "[compact](https://www.microsoft.com/resources/documentation/windows/xp/all/proddocs/en-us/compact.mspx?mfr=true)
" con los siguientes argumentos:

> compact /u /s /a /q /i
Empleas el comando sobre la carpeta afectada (la opcion "/s" indica que se recorran los subdirectorios), de esta forma le indicas a compact que descomprima los "archivos azules" y los regrese a su estado normal.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEh6wpiD_5M246MsilZwyVwOOm4_-OmKMEahIsuTYKNAePG9OEdsTkdjWktT8b4kht_9Cruogtq3u3N03lBVjs0TjVgEkVyNWQM9Q_B7BL9VZdnY08CsoUFxDabHSWYkQC7tameyALlIk-vE/s320/allnormalagain.png)    

Yo ejecute el comando sobre mi carpeta de proyectos, me funciono a la perfección, y como pueden ver, me salvó de perder mucho tiempo haciendo CheckOuts y configurando mil cosas extra...

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgrRX1KQAu9sfsAghT_I008X7ZIwaw63jFGjk7wBqcdXcxVONSdlSk4PZMo2abkeMzBGZS3qKBJ_2qoYTjCzA9_8nZe9ZL_kAuThFJKiOeamfdfx1Uz_weKtjVG1l_wD6gCMSpdkftDvfeY/s320/talegofiles.png)    

Bien, ese fue un tip que salva de dolores de cabeza, espero que les sirva, cuidado con los pitufos de windows, ¡Saludos!

Nota: La solución se realizo en una maquina con Windows XP Pro + SP3, el autor NO da garantías de que esta funcione en otras versiones del SO utilizado.