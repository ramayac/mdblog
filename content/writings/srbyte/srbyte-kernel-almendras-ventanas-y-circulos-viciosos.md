---
title: Kernel, Almendras, Ventanas y Circulos Viciosos...
date: 2007-07-08
author: Rodrigo A.
tags: software, kernel, linux, libre, gnu, virtualizacion
draft: false
post_id: blog-3515952828243908885.post-1172145899024398733
---

En el oscuro y ofuscado mundo de los Sistemas Operativos, un Kernel (también conocido en español como núcleo) es el componente central que se responsabiliza en manejar los recursos del sistema (manejo de memoria por ejemplo) y la comunicación entre el Hardware (parte física de una computadora) y el Software (programas).

Como componente básico de cualquier sistema operativo, un Kernel provee el nivel mas bajo en la capa de abstracción antes de llegar al Hardware... algo mas "abajo" del sistema operativo... son unos y ceros... o cambios voltaje.

La tarea del Kernel es ejecutar procesos (Software, aplicaciones, programas, como le quieran decir) y facilitarles a los mismos la tarea de interacción con el Hardware, mediante la comunicación entre procesos y llamadas al sistema (system calls).

Existen una gran variedad de Kernels: Micro Kernel y Mono Kernel y además esta el ExoKernel, el Kernel Híbrido y el Nano Kernel...

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEiaodYomHH_WHHgoRxar4uJa4PKgDoj1a70NGALe9RwLDCLa8EyogeuuwJY3YeF6uUPrgBNJeE46WlYxi79YwIS_GZ8dMBZNq-3J5DuDMwfvMp5st0mmnlOrABd60yuhzcRBZKo2kdZiwQ/s400/kernel.jpg)    
"Macro (Mono), Micro y Exo Kernel... el concepto"

...el diseño y creación de estos obedece usualmente a la implementación de los "anillos de protección". Estos anillos obedecen a un mecanismo que busca proteger datos y funcionalidad: de fallos y de comportamiento malicioso.

Los dos tipos de Kernel más usados son: Micro Kernel y Mono Kernel... y entre ellos existen un sin fin de implementaciones. Esto es como decir que solo existen Negros y Blancos... mentira, hay morenos, mulatos, cheles, chilucos, etc... Lo mismo es con las diversas implementaciones que existen entre el Micro Kernel y el Mono Kernel.

La forma mas fácil de imaginar un Kernel es como una almendra...

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgcC9rC8PtbntxX5RnHNcFDrPT4eHtaibkYoHdZ-lJO8vZzeUfup3ox_cH0ApGIAuqVlIZr9C-4-dFI4S-Ybi6XNSltfNx2b82xNPA3vvjPda0TRDjVevK8dGu9ELtawI7kC-fXZ5c30v4/s400/almendras.jpg)    
"Kernel en ingles, es el centro comestible de una almendra"

Hay que mencionar, que Microsoft decidió omitir esta regla de "anillos de protección" desde Windows 3.x hasta Windows 98. ¿Por qué?... bueno, la comunicación entre los anillos de protección requiere MUCHOS ciclos de procesamiento. Y con Windows siendo un sistema operativo elementalmente visual, Microsoft no tenia mas remedio que ELIMINAR los anillos de protección para mejorar el rendimiento general del sistema. Claro, esto llevo a problemas terribles muy bien conocidos en todos los Windows de la serie 3.x y 9x, como la famosa BSOD (Pantalla Azul).

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgk2ZQWvYDVOLGuB1hZdp6DcsZwcfvndccom4oDhmkBQyExBAPv6JbL76aTIwBzNC4mAhWYjamEHgTy1aA-nQr6te3vJQz2xPxF2BbF93iNkf5lAxFbTAmJ5c2MQTgbxR8dPJc7NguKYGQ/s400/kernelwin98.jpg)    
"Así se vería el Kernel de Windows 98 después de una BSOD"

Ahora la serie de sistemas Windows XP/Vista eliminan (enormemente) este problema implementando en su Kernel los una vez desplazados anillos de protección... que se asemejan mucho con las capas de una cebolla... pero basta de hablar de comida.

El tema del Kernel no sera muy tratado por los usuarios de Windows... pero crean cuando les digo que es un tema muy interesante y elegido como tema de conversación (bueno, ni tanto) por muchos usuarios de GNU\Linux... y no es solo interesante, sino también controversial y [¡extenso!](https://en.wikipedia.org/wiki/Image:Unix-history.svg).

Lo "bonito" del Kernel de Linux, es que como es Open Source, todo mundo sabe como funciona y esto permite algunas características muy interesantes a la hora del famoso proceso de "compilación del Kernel".

Compilar el Kernel en un sistema, no se trata de una tarea exclusiva para "geeks", basta con leer un poco y entender la dinámica. Lo que si aburre, es el TIEMPO que tarda en compilarse, y sumado a esto, tener que re-iniciar la maquina para probar el nuevo kernel, y sino funciona, hacer el proceso de nuevo.

Pues para este circulo vicioso, existe un método que simplifica MUCHO las cosas, y además, sirve de ejemplo perfecto de la flexibilidad que presenta el Software Libre. Este método consiste en la utilización de una maquina virtual para probar nuestro Kernel, y definitivamente NO consiste en copiar el kernel a la maquina virtual y arrancarla o peor aun, compilar el kernel en una maquina virtual, NO NO NO! Es mucho mejor que eso (y menos tedioso), consiste en arrancar la imagen de un disco duro virtual con un Kernel (el que hemos compilado) EXTERNO a esta maquina virtual.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEghyphenhyphencoj5MSQAhpY9NAVsxJofXbPdf2TBa5E221C7p99F8YuXGLx5nmxGTFNxlXSiO7dgR1qXVAI13FMgm9sN_iob6zgGuEFCXIlZ5HvpSYdt6IYc6apMIrwDNV0TM1LfWzlTLKzCpGiYu8/s400/qemu-logo.png)    
"Logo de QEmu"

Esta fantástica característica parte de [QEMU (QEmu es un emulador de computadoras virtuales)](https://fabrice.bellard.free.fr/qemu/), se llama "Direct Linux Boot":

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEguro53lwqOg6ueN4_a58jfGvjZTMcGo2yreLeFNihUIEznzO60pb20XbxoXE6Rc_91e7EZClkJ-TkG_j_hM6mDbMWGxscg-tKdDrzNkFHaFqKRzaCfTCrWNfwLJbwMJfYeiMUwFgnU3Dg/s400/linuxboot.png)    
"Captura de pantalla de Qemu Laucher con Direct Linux Boot habilitado"

Direct Linux Boot, permite ahorrar tiempo de una manera increíble, al evitar estar re-iniciando la maquina constantemente y evitando el infame "Kernel Panic".

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjTHA6NQVBNNcTcf_aSGvx5JfCi9IqUQKzfYVBT7CDTGqcq1D2pWEXUcbwZXbCBOxAg-lKTJUV-Akvj2BhMGf9NL03llSovpIAHuIz4L5TIDiQXjZhZqJj-NsKYR2uFkDlPf6AzNVyjICs/s400/instant%C3%A1nea1.png)    
"Captura de pantalla de Direct Linux Boot con Kernel EXTERNO 2.6.21rac Click para hacerlo más grande"

En el transcurso de la semana escribo un tutorial sobre como compilar y probar un Kernel usando QEMU.

Hasta Luego!