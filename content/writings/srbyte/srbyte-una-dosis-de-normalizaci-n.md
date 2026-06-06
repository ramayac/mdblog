---
title: Una Dosis de Normalización
date: 2012-03-06
author: Rodrigo A.
tags: colaborador, entidades, boyce-codd, normalización, base de datos, modelo
draft: false
post_id: blog-3515952828243908885.post-2402810206304774467
---

Nota: Este artículo fué escrito por el Ing. Alexander Calderon Peraza, docente de la Universidad de El Salvador sede Santa Ana, pueden leer más artículos en [https://basesdedatosues.blogspot.com/](https://basesdedatosues.blogspot.com/) y seguirlo en  [@calderonperaza](https://twitter.com/calderonperaza).

Desde el punto de vista de los programadores, administradores de bases de datos, y demás afines a la informática, normalización es un paradigma con el cual se busca que una base de datos relacional minimice los problemas de coherencia de datos.Este articulo busca refrescarnos las primeras formar normales, y abordar el primer paso en lo que muchos consideran formas normales avanzadas, no quiero profundizar en los detalles teóricos, pues esos los encuentran en los miles de libros aburridos de bases de datos que están en la red, en lugar de ello quiero llevarlos por un recorrido practico que nos dirija el día de hoy hasta la [Forma Normal de Boyce-Codd](https://en.wikipedia.org/wiki/Boyce%E2%80%93Codd_normal_form). En primer lugar planteamos un escenario: un modelo entidad relación para registra las visitas que realizan los empleados a distintos inmuebles registrados en la empresa, el empleado visita la vivienda y elabora unas observaciones al respecto, para realizar las visitas al empleado se le proporciona un vehículo para el rápido desplazamiento, la tabla RevisionVivienda almacena la información referida a las revisiones realizadas. Dedique 1 minuto para observar la figura y los datos ejemplo de la tabla RevisionVivienda.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjxZLStyvKtcwYAxeNI_cZvPZps3ih-bxfsgNY4Evk5zyzKIHM6r5uMDxXJnD0YAI-CPJaBKgzI_vppwvNL597qeqNDIemRTJBaOduToWUgvGJ9A4asWyILXAT7F0cjKUguZeZxIKR7-QXi/s1600/normalizacion1.png)    

¿Identifico la anomalía en dicha tabla? ¿Aún no? Observe las visitas hecha el día 1-1-12 del empleado EACP, ahora imagine que ese día el empleado se equivocó al momento de registrar los datos, pues el no llevo el vehículo P1111, sino que utilizo para dichas visitas el carro Z2222. Si usted debe hacer esta corrección a los datos ¿cuantas tuplas tendría que modificar?

Ahí está el problema, debería, si el modelo fuera adecuado solo modificar una tupla, pero en lugar de ello, debe modificar en este caso 3 tuplas, a pesar que los Carros son asignados todo el día a un empleado.

Veamos paso a paso el proceso de normalización para la tabla RevisionVivienda y así solucionar la anomalía.

1FN:

La tabla RevisionVivienda ya está en primera forma normal, posee una llave primaria, compuesta por el ID_Vivienda + Fecha y no tiene campos repetidos, importante es mencionar que una vivienda solo puede ser revisada una vez el mismo día, por ello se escogió esta llave primaria.

2FN:

Veamos, esta cita que todos los campos que no son llave primaria, deben depender funcionalmente de toda la llave primaria, y no de una parte de ella, los campos ID_empleado, Observacion y Id_Carro, cumplen con esta condición, pues dependen de toda la llave primaria, no de una parte de ella.

3FN:

Esta cita que aquellos campos que dependan transitivamente de la llave primaria por medio de otro campo, deben ser eliminados, veamos nuevamente los últimos 3 campos, ¿dependen transitivamente de otro campo? La respuesta es que NO, esta tabla si cumple la tercera forma normal.

Bueno y entonces qué pasa con nuestra anomalía, paso las primeras 3 formas normales. Para ello existe una solución, y es aplicar una tercera forma normal más fuerte, que se conoce como Boyce-Cod.

Forma Normal de Boyce-Codd:

En forma sencilla cita, que debemos cumplir la condición que todo determinante en la tabla debe ser llave candidata, pues entonces veamos los determinantes que existen, y si alguno de ellos NO puede identificar de forma única una tupla, dicha dependencia funcional debe eliminarse de la tabla.

1Id_Vivienda + Fecha -> Id_empleado, Observacion, Id_Carro                   (llave primaria OK)

2Id_Vivienda + Id_Empleado + Fecha -> Observacion, Id_Carro                (llave candidata OK)

3Id_Empleado + Fecha -> Id_Carro                                                            ( Problema)

Observemos la tercera dependencia funcional, el Carro es asignado a un empleado en un día especifico, dicho empleado tendrá asignado todo el día ese carro para realizar sus visitas a las viviendas, pero esta dependencia funcional NO es una llave candidata de la relación, y ahí está el problema.

Para solucionarlo debemos tomar esta dependencia funcional y convertirla en otra tabla, de la cual será la llave primaria el determinante de la dependencia. La solución se muestra en la siguiente figura:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjvgQ7mtaxQ3MFuTUFW_VNhTJxWtiUAxWO5s9RvFvjZAAIxvdGlz0U2gofwgx-S9nAqqWPh4XAO7BRJUJcHf2lCypEy6O7TYVjkd4-CYkFoL3B_wI8ObZCwVJzLVyucvBb2tEIm-4wvy0iy/s1600/normalizacion2.png)    

Noten que se quitó el campo Id_carro de la tabla RevisionVivienda, y se creó la nueva tabla AsignacionCarro.

Espero esta explicación les haya servido de mucho provecho, y como ven se ha explicado de manera sencilla y muy al grano, evitándonos horas y horas de leer páginas de un libro con muchos modelos matemáticos enredados. Esperamos en próximamente explicar la siguiente forma normal en otra dosis de normalidad.

¡Saludos!