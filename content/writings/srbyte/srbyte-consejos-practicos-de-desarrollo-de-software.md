---
title: Consejos practicos de desarrollo de software...
date: 2008-07-26
author: Rodrigo A.
tags: codigo, util, programacion, programador
draft: false
post_id: blog-3515952828243908885.post-1286650227634257416
---

Como desarrollador ocasional de software, me he aventurado en el área de la producción de software a la medida, para un par de individuos. Como muchos desarrolladores, me encuentro con el terrible y repetitivo paradigma de ser un: Programador-DBA-Diseñador-Tester, lo cual, siempre termina siendo un dolor de cabeza. Pero aunque sea un dolor de cabeza, no podemos negar que es una realidad que a todo estudiante de alguna carrera afín de la informática termina haciendo para conseguir un par de $$$ extra.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgYvcAUObL5Y3pepXrMptx_OnfWMjChvEuQgMGeLUnvfs3tu5mwRvVDJ-Kmbw3r38ClrO4ilf3Rmw3a6D9aPoeVQSJU_Skb_8CZztoA8akgqj2juRnsDmVCGIKJ9VkqXP3Cg1rRaOj82bc/s400/hidden.png)    
"Arrrgh!, que asco de GUI... x_x"

La teoría nos dice por ejemplo que un Programador NO debe hacer trabajo de diseño de la GUI del programa, pero la practica y la realidad (no mucha gente capaz en el medio) nos enseña que podemos hacer un poquito de todo tomando ciertas medidas, que en mi opinión son necesarias para mantener la cordura en todo el proceso de desarrollo de software.

Así que sin más retrasos, les presento una lista de consejos prácticos - que parten de mi experiencia - de desarrollo de software:

1. Planea entrevistas, visitas y un par de meses para entregar el software. Planea entrevistas con las personas que utilizaran el software y visitarlos varias veces al mes, la idea detrás de esto es saber cuanto conocimiento tienen la o las personas que usaran tu programa. Deberías de ser capaz de contestar estas preguntas de tus clientes: ¿Qué tipo de usuario es? ¿Es un usuario común, avanzado o un PowerUser? ¿Qué sistema operativo prefiere/utiliza?

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjWnEoCh9rEA0nytbMDU2uxu7WxYwy43-t4CDVuNykZTLoV5MuMiohyenHwO73JtjYPsIoX_EZaMKv_YLgBj-JLeXc9Pl1bBIjB3BLSYYEFIP7XScJEQAntv6HvhgiDUWC2SnQLa1pamak/s400/prospectiveboss1kopie.jpg)    
"Conoce a tu cliente para atender sus necesidades informáticas."

2. No esperes que tu cliente cambie de sistema operativo solo por tu programa. Pedir que tu cliente cambie de sistema operativo solo porque tu programa usa "cron", es desconsiderado y poco profesional. Puedes recomendar que cambie de sistema, pero no lo cambies solo porque si. Si eso sucede, implica que como profesional tienes que hacerte cargo de que la mingración sea 100% satisfactoria, sin excusas.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEg0Qiaw-Il7xNyFz6S6ExQNjzO0rZYtw-piBoeeQNATEMaTcsJLlpm2fjswcrB5-ldvZJKn2c01ajxCaZCTREM9BarphFR2XQIAlIfqqdtnixfgi1J3lCMBSjeFPcqZXJYuM0TCGARSnIY/s400/hypnoGuy.jpg)    
"Trata de no atentar contra la comodidad del usuario."

3. Utiliza Frameworks y herramientas de desarrollo populares. Utiliza .NET, Java SDK, RoR, etc. No te quedes con lenguajes o entornos de desarrollo desfazados como VB 6.0, Delphi 7.0 ó VisualFoxPro, se practico y ¡mantente al tanto de las herramientas que te hacen la vida más fácil! Al menos deberías de usar/conocer lo basico: Unit Testing, Code Coverage, Source Version Control, Automated Build/Deployment. Si no sabes que es nada de eso, más te vale comenzar a aprender, porque son herramientas que te ahorraran tiempo.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiHV8Ndk0cl5O5M_796jluZ0dU30_fTIEVZoijNU-xqJ8AeaPjVXrKt515MIo9CXms3eKP9EsI6b4TePCbD0V1VJ36o7x1nN3eptdimM1zYkFFNFWz-KXXgytzXl3WXMpcATlfs73m0rz8/s400/800px-Framework_complexity_of_the_Pater_Noster_lighthouse.jpg)    
"Utilizar un framework hará que tu trabajo sea más rápido."

4. No reinventar la rueda. Si algo ya esta hecho, es software libre (opensource) y lo entiendes, entonces: tomalo, modifícalo y usalo. Y reconoce al autor del código que usaste.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgGmp6OBhqp7Jlr0Ngcyab_yL0EGe_3yWu2iqaUaDTApDk8ajOTN8ze2ZyBmXP29BGlBO571V6sLWitTA_WWjHkIVgdcYWXl75lxEbBnJEsS31EiJx7IheGqB2qapX_3fITit6KniuyZmY/s400/wheel3.jpg)    
"Solo puedes mejorar la rueda... ¿Para que re-inventarla?"

5. Concentrate en solucionar solo lo que prometiste. Imagina que estas haciendo un programa tipo Agenda. Y de pronto se te ocurre que te gustaría que envié mensajes a celulares y correos electrónicos que sirvan de notificaciones... si no te lo pidieron NO lo hagas. Si agregas constantemente características a tu programa: terminaras con un "Frankenware".

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhHvNyI5ryUecVVisIWJyHj32ixLVRDeBUbym2S-VE5LzLTI8YEgYC1gxouvOMbOsoUYCE-2uACP5NO9b1i8QW2-wKE_c7gldHfGIqpcoqRcxDgE_NJQ6y04CGMxDbFU54wpsBxo1ksXbc/s400/candy.jpg)    
"No pierdas de vista tu objetivo (vivan las 3D manías)"

6. Busca la simplicidad y el minimalismo. Ingeniate para que tu GUI sea sencilla y clara. Ejemplos de simplicidad es el iPod, el iPhone y la pagina principal de Google.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgtQDvqlSrrMMSHlewakLiXpemHzXBp_i-XWQaXHJu8Ofns9qorD1o48qW0VgHY7PRtEjzxHSmqsHcAyoskwTY3RB3Enbkawqcfcxm7O7jZllNU4KSiFGu1bWuGtbA3brHuNKyxyN2DccM/s400/yourproduct.jpg)    
"La imagen lo dice todo, haz click para verla más grande."

7. Si puedes construye librerías/controles propios y ¡compartelos!. Crea una serie de librerías/controles personales que te ayuden a realizar tu trabajo más rápido y asegurate de compartirlo, ya sea con el mundo o con algún colega. Esto te servirá para madurar como programador y mejorar la legibilidad y el estilo de tu código.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEh_vdauvmDxQ0wWTuQEXSq2u32mYhRHDiPQsaEBdGj1m-ik0G26BzCngCaBM2tC-vObtF8P2r7BMWWDadK9Zu0Wvgpg8OrhuTtEOoNGz9TzRorYnNIM2f-VAHqbWPhfNjB496S2wwrEliw/s400/security-padlock.jpg)    
"Un programador celoso no es un buen programador."

También comparte tus experiencias, a más de alguien le pueden servir :) Y tú, ¿que tipo de experiencias has tenido como desarrollador de software?