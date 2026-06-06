---
title: MoMA, la mamá de los monos...
date: 2007-07-20
author: Rodrigo A.
tags: mono, software, linux, .net, libre, gnu
draft: false
post_id: blog-3515952828243908885.post-5401259346683657008
---

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEh3s9JmMTYoQIDuJfhvA4XFuiOma2efuglQ8oYEYC0iBcwzavHjQadzsa_Atp9_8jKhUhY6SZcsNfSCfGc8BOXWnZPxBuC6A-XCLl32SGil57OHxf5Ze1IqFHWHYZFrWfwkC8b433lqd0I/s400/mono.gif)    MoMA o Mono Migration Analyzer, es una herramienta que ayuda a identificar diferentes problemas que se tengan al portar las aplicaciones del Framework .NET al Framework de Mono. Aunque MoMA pueda ayudar a enseñar problemas potenciales, existen muchos factores completos que no pueden ser previstos mediante el uso de una sola herramienta. Pero la utilidad de MoMA es innegable. Si tienes una aplicación hecha en .NET que no pase ninguna prueba de MoMA... entonces ya tienes una idea del trabajo que tendrás que hacer para migrar esa aplicación al Mono Framework.
> Recordemos que la verdadera prueba resulta al correr tu aplicación sin problemas en Mono.
Pero para que se acostumbren a codificar bien, y tengan una idea de como usar MoMA, aquí esta una pequeña guia, ¡espero que les sea de utilidad!

nota aclaratoria: las capturas de pantallas ajustadas a menor tamaño no son tan importantes o interesantes como las de tamaño "grande", esto es claro, bajo el criterio del autor. Como siempre, pueden dar click a cualquier imagen para verla en tamaño real.

Guía ejemplo de uso de MoMA:

1.0 Descargando MoMA

Puedes descargar MoMA de [https://www.mono-project.com/Moma](https://www.mono-project.com/Moma) , luego ve a la seccion de descargas "Download" y allí encontraras el vinculo de descarga.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjC-uOcyECdcmbqSUKntZbuVv5Docd_yMF_FBRYTfdEmqywEeHQE9XHnIR8I9XzWUFOpbRs5KgHn3082nvmRZwasamzNxwatnkdf3y0vhmzA9esuI2NE6oO7fWAU57IaVHMF-UTjJKxTzk/s400/momamono.png)    
"Captura de pantalla del sitio web de MoMA con el URL marcado (muy marcado) para descargarlo"

Recuerden que MoMA necesita el framwork .NET 2.0 ó Mono 1.2+, como solo estoy usando GNU\Linux, usare Mono para correr MoMA.

1.1 Ejecutando MoMA

Después de descargar el archivo .Zip de MoMA, extraemos su contenido y podremos ver lo siguiente:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjDykjRFi7H63xqBH5HDVxBnB2wIZWSSkoxQLk7SF1eJi6LiSiGh5CI90WMxrB78wlhN7iNfd9lrvm-MPXZ8u5HC7jELMiieFQ2SQoUbBs_7lAugWuf-aAVK1UICgPDioY3dr_r5G9ih1Y/s400/momamono-carpeta.png)    
"Captura de pantalla del contenido del archivo .Zip descargado"

Ahi estan los archivos que necesitamos, muchos diran: ah!, son simples ejecutables de windows. Pues noooo, estos archivos "exe" estan compilados en un lenguaje intermedio (como el bytecode de Java) llamado MSIL (MicroSoft Intermedium Languaje). Y las extensiones "exe" y "dll" pues solo son extensiones re utilizadas. Vamos a usar Mono para correr/interpretar/compilar (como le quieran decir) el archivo "MoMA.exe". Abrimos una consola (en mi caso usare Konsole) y nos vamos a la carpeta donde se encuentra "MoMA.exe", ahi ejecutamos el comando:

> mono MoMA.exe
Aquí hay una captura por si se pierden:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhDO6Dd9byo9plNutIXAEhMPBT3SduapMmLK02xWEL_3lDc0jhu2zjxWPz3Mw9a3eMU2qQDgTTL0xdQSstFcVyfZEbzsld95QjRNPIYHEP7utST6ECoi8aHL0KpR0eE4aQHp2HesKu1rR0/s400/momamono-ejecutando.png)    
"Captura de la consola, usando mono para ejecutar MoMA.exe"

Y si todo sale bien, pues podremos ver la ventana de MoMA en nuestra pantalla:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjNk1puV-4rHqkY9ur522kYwjFV40PXgo3OEq1b2krJlgJvqGfeKAAnWVOlR_YzICLJEAcWwR4bMNoRa37hLr4xeyWIFo2RBtJy5ClPuAG-qKqBxoPFNfFVVedoqAFwKNsG-OJb786lKsY/s400/moma-ejecutando.png)    
"Captura de MoMA en GNU\Linux, usando KDE como Desktop Manager"

1.2 Usando MoMA

Para usar MoMA solo necesitamos ejecutables hechos con el Framework .NET y dar unos cuantos clicks. Yo usare de ejemplo la librería de un amigo para comprobar si puede ser ejecutado o no con Mono. Ustedes pueden usar cualquier otro compilado (el que ustedes quieran) que corra en el Framework .NET 1.0 y 2.0, y que evidentemente no sea una librería especifica a la plataforma o parte de .NET (como System.dll). Esta es la librería CustomDataSource(CDS) Ver.1.0 de Roberto Linares, que puede ser encontrada [aquí](https://ingenieria.uesocc.edu.sv/foro/viewtopic.php?t=77).

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjh-0NBPBd-GYpw5ye9J4yyv5jNJ4NpqgpkEtriKctiyD2Go92GavMKo3Q5J_A9UqtDZsUxceJb8s1lVUV8o5MEY6D4VsD5q4X1wUH1RQtG0fbdtzKa1SYxGbJhyphenhyphen1I7gRKfyQcvSSoQrNA/s400/descargandoCDS-bin.rar.png)    
"Captura de pantalla, descarga de Bin.rar de CDS, foro ingenieria UES-FMOcc"

En el paso 2 (de 4) de MoMA, podemos agregar uno o varios ejecutables hechos con el Framework .NET, asi que despues de descargar los binarios de CDS, y extraerlos en mi escritorio, agregare el archivo "CrypterClassLib.dll" dando click en el boton con signo "+" y luego eligiendo el/los archivo a agregar, asi:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEj1TuNYSO0wYnlCB7wxgkA4fdrOadMjggZUH2Ev62hRlW6frMKKHpAH74LYP_iUnhYvEOrHP53f_nZUxXdVlv7pjSm-BJuuvK5XltOyP9N8tofmVX3GQPaSN2BDd2RhOZMhiGG5rh7H54s/s400/libroberto-moma.png)    
"Captura de pantalla, agregando archivos para analizar con MoMA"

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjcrC1Ycn66CDLXXqW88i9iZ8gZZ03TFSVtlW7jb25p3e2vsr3IEuumWecvohTT9J3wF4sl_8X1ODPEHuC7tkjKymY8Doz8JrM85oYDTmKNDhTzeoGzIdLQi1VFS49CoeDhfJMo4PNMRlg/s400/libroberto-moma-selec.png)    
"Captura de pantalla, ¿Listo para analizar?"

Ya con la lista de archivos que seran analizados, pues solo basta dar click en "Next"; para que MoMA analize el ejecutable y podamos ver un resumen de compatibilidad del archivo con mono. Al final obtendremos un resultado como este:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiIPjUBx3Tj-YbP9GE0R0-9PGyh_IMLsQvJXErf1FteJNE0l0UgNpVPyeXk9-ohP8Q748vQEFm9oAV0El42hzkxu3VXBpLSiLacSJRp-EG_fbZIaibHejKW3Yk20Q-P9Kcirzh4i2Hs1kI/s400/moma-todobien.png)    
"Captura de Pantalla, ¡Te salvaste Roberto!"

En este paso (3 de 4) podemos ver que en el análisis general, el .dll analizado paso las pruebas con exito. Asi que en teoria, el archivo "CrypterClassLib.dll" puede ejecutarse sin problemas con el Mono Framwork. Vaya Roberto... te salvaste del bochorno publico jaja ;)

Pero... ¿Que sucede cuando TODO sale mal?

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgpyHhSJOncVRW3GzLQ3zEObyQORJoP7Ib1Qrl_0Xa-8Odaawwt1TnKL8o4p4-iVuxfsttCIguygfefhHEbYpZG7Lz9xeaO0dKr8XAAIo0T7cetHVjYkKPLIF01R2wfSJ3lz4U8dpR_cHA/s400/moma-mal.png)    
"Captura de Pantalla, ¡OH NO!"

Bien, si sucede algo tragico como lo muestra la captura anterior, o peor... pues MoMA genera un util reporte (paso 4 de 4) que muestra de manera general, que se puede hacer para combatir esta incompatibilidad.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgr7jnlzA63o7MJ57-Q80xgnmoxEMHHsTWJQhQ0pB2UqZDoh5RX76fIw9Vx1kmwZdGD_X-xyBMv4qc2iKuLAlFiWD3p_hPT4g12kVyRzOFAS88eYV-662hZ__1VsQEXWc4N27cqgLl7zr8/s400/Report.jpg)    
"Captura reporte de ejemplo"

Con este reporte en mano y teniendo en cuenta ciertos tips de migracion/codificacion seremos capaces de migrar una aplicacion casi (especulativamente) sin problemas de un S.O. a otro.

1.3 Consideraciones finales

No olviden enviar el reporte (sea bueno o malo) de compatibilidad que genera MoMA (MoMA lo envia automáticamente):

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEi0B7gPHrYW3jUCA42ojlVIvEaCxUQ3SN_7TRkXScElgP9zw7ayoOgJPPsbJ6zBTFn4z814IJHNaoR1ahechxFIyrOuJWR-PJ7jrwO1H6u2sxSlK-fFOf-nk871waLGEAuRhfGnm_YSpZA/s400/envio-reporte.png)    
"Captura de pantalla, envío de reporte a Mono Project"

Esto ayuda para que los desarrolladores de Mono coloquen especial énfasis en los problemas mas comunes de compatibilidad encontrados, es tan fácil como dar un click al botón "Submit Report".

Una guía mas completa y extensa para corregir problemas de interoperabilidad puede ser hallada [aquí](https://www.mono-project.com/Guide:_Porting_Winforms_Applications).

Hasta Luego!