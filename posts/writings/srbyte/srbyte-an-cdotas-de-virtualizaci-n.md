---
title: Anécdotas de virtualización ...
date: 2009-07-12
author: Rodrigo A.
tags: ubuntu, svn, virtualizacion
draft: false
post_id: blog-3515952828243908885.post-7528541824882962586
---

La semana pasada, el servidor de Source Control ([Subversion](https://www.srbyte.com/2008/03/programemos-mejor-subversion.html)) en donde se encuentran alojados todos los proyectos informáticos de la empresa comienzo a presentar problemas SERIOS de rendimientos. Con problemas serios me refiero a tardar 2 o 3 minutos para realizar un commit de 13 Kbs...

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgls70BPdOU6x4M_9fBEQiCA7Nzhbo-ChmGusM30ZAdDTFkEixQd8eoIRBkgLMLHrTEjygmZac3SH6-yAZkF1jKk20XmRLBS5q5FltQgCKi8aktWVQ9uECmk7mhhyphenhyphentzYixBQ5H_OLMfq355/s320/subversion_logo-200x173.png)    
Todo esto comenzó a suceder justo en el momento en el que tres de los proyectos principales se estaban convirtiendo a TAGS ("Estable") y se comenzó a realizar un control mas minucioso de los commits realizados (se realizo lock en el main branch de los mismos). Así que la cosa se puso bastante seria cuando de un momento a otro, el rendimiento de este servidor, que se ha mantenido estable y confiable disminuyo, y en un momento, hasta se pensó que había "tronado", justo como trono el servidor CVS que le antecedía.

Como nunca había visto el dichoso servidor, pues me decidí a buscar en donde se encontraba, así como buscar al responsable del mantenimiento del mismo, o al menos al que había configurado el SubVersion ahí, y realizar la típica sesión de preguntas de novato curioso (de buena intención, pero molestas), y de paso como ya había leido de un problema similar en el servidor SVN del código de KDE, que es al menos, unas 10 veces mayor que el de la empresa, sabia "mas o menos" que decir/mencionar/sugerir para la mejora del problema. Resulta que el administrador mas inmediato de la maquina es un amigo ([RobMV](https://robmv.com/)), así que le comente la situación, y mas o menos la conversación se dio similar a esto:

> - Mira, el SVN esta lento, podes revisar que sucede?
> - Ok veamos... hmmmm, la máquina esta lenta. (Despues de un rato) Ah! tiene menos memoria RAM asignada
> - Asignada? - pregunte.
> - Si, asignada, es una máquina virtual.
> - Ah! Nice! (me brillan los ojos cuando dicen virtual).
Resulta que el servidor REAL (Sun Blade Server con 16 GB, y dos procesadores Intel Xeon de 3 GHz c/u) había pasado recientemente por una "reasignación de recursos" para TODAS (al menos 13) las máquinas virtuales que en el se ejecutan, para mejorar asi el rendimiento de otras máquinas [virtuales](https://es.wikipedia.org/wiki/Virtualizaci%C3%B3n)...

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjWJmCLh3ymcjupn7Eubh9NmM8zqzmTfuLMYzbObpza_FeIhJq4i5183VskpxTsD6J8-ebQ10nlquyOXtgETF270h3PSRDLfo1KeogJQjvIYQNkbkiEu_O2VKcbDIjUX_PLMuq7werbrdUY/s320/virtualizacion.jpg)    
El "Blade", posee 3 niveles de prioridades en las que resuelve o "cede" el uso del procesador para cada máquina virtual. El servidor SVN estaba en la prioridad más baja. Ademas, la memoria RAM del servidor SVN se redujo de 1 GB a 512 MB. Y para colmo el servidor SVN se ejecuta en nada mas y nada menos que Windows 2003 Server.... Ah!!!! con razón esta lento.

La solución inmediata, fue reasignar la prioridad al servidor, a un nivel de prioridad "alto" (nivel de producción), y santo remedio... No digo que no funciono, pero creo que si alguien tiene algun problema similar, existen más opciones para solventar el problema. Creo firmemente, que ese mismo servidor SVN puede ejecutarse perfectamente en esas condiciones, y con mejores resultados que como lo hacia antes.

Mi solución, radical y simple:
> quitar Windows 2003 Server y usar un SO especifico para maquinas virtuales.
Si bien Windows 2003 Server, es bastante estable, cualquier experto puede concordar conmigo con que este no es un Sistema Operativo optimizado para ejecutarse como una máquina virtual, ¿entonces para que molestarse en tenerlo instalado en una, y gastar además en su licencia?

Para las máquinas virtuales, siempre hay que usar una regla de oro:
> Usa un sistema operativo OPTIMIZADO para máquinas virtuales.
![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiC5lzpqExPr7Hw8aXE5kcMluyi2RJrMUNhrm1e4HW60Xrh3J0P6TZVxGdP1HI0yNhE6TKLXku1v2QtXufnr1UhYavaslaTHOiHpU86mkUAUWecbhm7RabdeNay0Jwjg03A6vUrfcfF4hVL/s320/vmware-player-ubuntu-install-1.png)    
Un PERFECTO ejemplo de esto: [Ubuntu Server Edition JeOS](https://www.ubuntu.com/products/whatisubuntu/serveredition/jeos) (que se pronuncia como jugo en Ingles: "Juice"). Beneficios inmediatos de usar JeOS son:

- Mejor rendimiento en el mismo "hardware" comparado a un sistema operativo completo no optimizado.
- Menos espacio en disco
- Menor cantidad de actualizaciones (mas consolidadas y de mas importancia), lo que reduce la cantidad de mantenimiento del mismo.
Sistemas operativos como Ubuntu Jeos están afinados, de manera que aprovechen el máximo rendimiento de productos como VMware y KVM, lo que se traduce en mas eficiencia para escenarios de virtualización mayores.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjOHV-c4Uea6fB_cTjOhH2vW32Zu3nneJje7RnDvFFD8SqlIFL_FW_839v8E1RCfKsObtB6fa7m_k8VX582v8DA8Xpy4AqwKU0wpe3VpU4IEUE95OpFPAKBXVeoGvrq805rW0cSTCrcD7xn/s320/365x230.jpg)    
"JeOS = núcleo de SO {Kernel, Drives, Login} + Mínimo Mantenimiento + Mínimo "user space tools"" Si la idea es "sacarle" el jugo a los equipos actuales, y mejorar el rendimiento sin incurrir en gastos por la "crisis", entonces hay que hacer conciencia sobre soluciones que usen Software Libre, y ofrecerlo como una opción REALISTA a los problemas informáticos empresariales.

¿En tu trabajo, usan virtualización?