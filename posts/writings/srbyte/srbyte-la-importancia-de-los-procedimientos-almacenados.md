---
title: La importancia de los procedimientos almacenados
date: 2009-04-13
author: Rodrigo A.
tags: procedimientos almacenados, practicas, programacion, base de datos
draft: false
post_id: blog-3515952828243908885.post-3533771959719215532
---

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiF2IwxzZ8q3SwcDODnBjmGSVl6HViA4vtWatOb-ayWBQEu7TNePONFIbQM7FtWzrKw2qIDCj6Q0HOLgav2XmGKsYJw9PlG0BxfZSTYT3GrSPB0hnEb8QwgvpbqSptlDqFlP6uP1o4XQRn0/s320/codecode.jpg)    En esta ocasión, me quiero concentrar en algo muy importante del inmenso mundo (y misterioso para algunos) de las bases de datos, quiero hablarles de: Los Procedimientos Almacenados.
Los procedimientos almacenados (conocidos también como proc, sproc, stopro, o SP's por sus siglas en Ingles) son subrutinas que están disponibles para las aplicaciones que acceden a una base de datos relacional. Los procedimientos almacenados están, como su nombre lo indica, almacenados en el diccionario de datos de la base de datos.

¿Para que existen?, ¿para que usarlos, si puedo tener mis consultas bien bonitas metida en mi código (como consultas SQL Ad-hoc)? Esas preguntas exactas me hice yo hace algunos años. Que pena ser tan ingenuo. Gracias a Dios desperté de ese "lapsus ..." y si bien, no lo se todo, lo que se, lo tengo que compartir.

El uso más típico de los SP (de ahora en adelante asi me estaré refiriendo a ellos, en vez de escribir Procedimientos Almacenados) es proveer una validación integrada en la base de datos, asi como proveer mecanismos de control de acceso a la misma.

> Los SP son usados para consolidar y centralizar lógica que originalmente se implementaba en las aplicaciones.
Eso quiere decir, que en vez de tener una "clase", con una lista interminable de funciones que hacen esto:

```
static function String getConsultaUsuario(String usuarioid) { return "SELECT * FROM user WHERE userid = '" + usuarioid +"'"; }
```
...usted debería de tener una serie SP en su gestor de bases de datos, implementado como un API "consolidado y centralizado" (en la base de datos). Lo mismo sucede con muchas consultas SQL "grandes", o secuenciales: no podemos dejarlas regadas en el codigo, necesitamos moverlas a un solo lugar...

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEi-wBsZKXAmeGVFVeQh9t-x-ZcEUU426Gn_MdL_PasIZTWkYGKyAS2XQPM7bs6EwkCob_IpRTqYgPNJLBkRfOTm4VNo2BfDRvSE8LoCHDvK0KPydLpkee2O860cbHdMJDCap8dz32Z_bgWr/s320/mysqlstoredprocedures.jpg)    
"MySQL también tiene procedimientos almacenados"

¿Por que? bien, enumerare solo algunas razones:

- El mantenimiento de las consultas SQL es mas sencillo, porque las consultas SQL están centralizadas en el gestor.
- La seguridad se enriquece utilizando SP's, debido a que un buen DBA (Administrador de Bases de Datos) únicamente proveerá una lista de procedimientos (bien tipificados y documentados) que el programador llamara de acuerdo a sus necesidades. Asi, relevamos al programador de preocuparse de verificar que los INSERT, DELETE, etc... estén "bien", lo aliviamos de la "carga" de manipular directamente la base de datos. Y dejamos esa lógica en donde debe de ir: en el gestor.
- Suponga que debe añadir una condición extra a la función de ejemplo "getConsultaUsuario", tendrías que cambiar la cadena que se retorna en la función, probar los cambios, compilar o publicar tu aplicación. Los SP se pueden modificar, sin necesidad de recompilar y hacer un deploy completo de la aplicación.
- Y lo más importante de todo, los procedimientos almacenados SIEMPRE serán más rapidos que una consulta externa. Si no lo son, es por dos motivos: a) No puedes escribir procedimientos almacenados, y/o b) El gestor de base de datos que estas usando es deficiente.
Entonces podemos observar dos ventajas inmediatas: Reducir la carga de mantenimiento de código SQL, y aumentar la velocidad de la aplicación. La desventajas: Nos estaremos amarrando al gestor, y más de algún programador probablemente saltara por no poder "tocar" (manipular) directamente los datos en la BD. Esto no es necesariamente malo, pero más de alguno protestará.

Tengo que aclarar que no hay usar SP absolutamente todo el tiempo y en toda aplicación a desarrollar. Un caso practico para implementar SP en una aplicación, es en el que se tiene una sola base de datos, y diferentes módulos o clientes que se conectan constantemente a esta. ¿Se imaginan tener una lista de consultas en cada modulo o cliente que se conecta?, ¿Y si se modifica la base de datos?, ¿cuantas consultas y en cuantos lugares lo iremos a modificar?

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEggFChxXmsJy4D1e6gnzMqY1eOCnBU6e2FHxYVNiqGhy6weCUygJNmw_nr4v8UC7G6DGz2U-kPfIixdPoBhesX9o71napsH8j_Tvghzj2pJzWy83IlkocguVILJ7WCfh0hkRVmAmgajhOeK/s320/sqlalchemy.jpg)    
"Los SP son parte de la Alquimia de SQL que usted debe conocer"

Los SP en general, serán realmente útiles (y aplicables) en caso de estar desarrollando aplicaciones robustas, de categoría empresarial, o para servicios que deben ser seguros, que posean alta disponibilidad y con vistas a ser escalables.

¿Y dónde aprendo, donde veo ejemplos, etc? ... como siempre, [Google tiene la respuesta](https://www.google.com/search?q=store+procedures). ¿Y tú, estas utilizando Procedimientos Almacenados en tu Software?