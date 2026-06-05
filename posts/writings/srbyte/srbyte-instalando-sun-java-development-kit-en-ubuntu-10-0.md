---
title: Instalando Sun Java Development Kit En Ubuntu 10.04
date: 2010-05-02
author: Carlos Peña
tags: srbyte
draft: false
post_id: blog-3515952828243908885.post-3599571272775863470
---

Luego de instalar exitosamente Ubuntu 10.04 en mi portátil pretendía instalar Java JDK para iniciar la configuración de mi entorno de desarrollo utilizado con mas frecuencia. Para ello me propuse utilizar [Ubuntu Software Center](https://en.wikipedia.org/wiki/Ubuntu_Software_Center) , la aplicación gráfica para el manejo de software incluida en Ubuntu.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgkatdld_l_y1tDCLFj4U99uQWk9lM0psdqyFy84KpAZQLYblG0kLY4q6nHv1jYkXw_yoy14I5uEQ5NqN_WL9yCwNEllngTYVk2-Pdhm6gnylwNQdsASYHeNNSU7yEzeBKHd_YiD5qZlAU/s400/Screenshot1.png)    
Captura de pantalla de Ubuntu Software Center.

Ubuntu Software Center permite realizar búsquedas de paquetes en los repositorios configurados, así que aproveché esta función para buscar el paquete que corresponde a Sun Java JDK y así instalarlo. Para mi sorpresa, en la lista de resultados no apareció algún registro que hiciera referencia al paquete Sun Java JDK, solamente aparecieron referencias a [OpenJDK](https://en.wikipedia.org/wiki/OpenJDK).

Luego de unos minutos de "abatimiento extremo" recordé que en Ubuntu existe una aplicación gráfica que permite manipular los repositorios de software de una manera sencilla. Se puede acceder a ella a través de la opción System > Administration > Software Sources.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEivTCXr5Br5IUqTeJepHxMDqa6U_huXq9h1GPjZv74smVxF_Zzk7RinSdcTZ_0-_UT1ZgW3BnFSNDPFw7xyXBLLH9Jb9Pr91deRqj8hZvtbv0VsrZiJ8F4SIDDFltUwMvleU12SiezzZ0E/s400/Screenshot-Software+Sources.png)    
Captura de pantalla de la opción Software Sources.

Para que el paquete Sun Java JDK aparezca en la lista de resultados se debe activar el repositorio Ubuntu Lucid Partner ya que [todos los paquetes que hacen referencia a Sun Java han sido desplazados de los Repositorios Multiverse al Repositorio Partner](https://www.ubuntugeek.com/sun-java-moved-to-the-partner-repository-in-ubuntu-10-04-lucid.html).

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgzjjWrYK8f4G8ob5xPIYhsylyaruQmwKwv06gDA7stVa3u-NK07XZYfANSGLjVfos3Lv7I28BGE0jL0Q7hayAL0D1lC_GV-1C4suGqjyFdYW67EVDSet8F89xtkIoOfvI_PrkhWO81tvI/s400/Screenshot-Software+Sources-1.png)    
Activando el Repositorio Partner.

Por ultimo hay que recargar la lista de repositorios - algo que se solicitará luego de realizar la modificación - y regresar a Ubuntu Software  Center para repetir la búsqueda. "Automágicamente" estará disponible el paquete para Sun JDK/JRE,  para instalarlo hay que posicionarse sobre el y presionar el botón "Install", se descargarán los paquetes adicionales necesarios y luego de una breve espera se podrá confirmar el éxito de la instalación.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEixHk2jhgth5b6RVKYRl6xVMZwSXSmKGj_tBwkVc7T_H384Q33D7ewCx8C2gpUzWRuyNQ4mhtxUNMDtqqDeYOLy2atW2o5oT7Gbi2MiDuUgTiSwSE8LN30xb0g4vcp0buQ7miN6OuleuJo/s400/Screenshot-Ubuntu+Software+Center.png)    
Confirmando la instalación de Sun JDK.

Todo este proceso puede realizarse a través de linea de comandos, sin embargo he optado por explicar el "método gráfico" para aprovechar la oportunidad de utilizar el Ubuntu Software Center. De nuevo les recuerdo que sus comentarios y sugerencias son bienvenidas. Saludos!.