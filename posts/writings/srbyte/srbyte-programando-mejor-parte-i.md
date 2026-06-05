---
title: Programando Mejor [Parte I]
date: 2007-02-17
author: Rodrigo A.
tags: codigo, programacion, software, lenguajes
draft: false
post_id: blog-3515952828243908885.post-3510358555177172075
---

En el mundo de la programación, existen miles de lenguajes (funcionales o no) que sirven para el único propósito de hacer mas fácil la comunicación con las computadoras... decirles como hacer una tarea, ordenarles que hacer y como hacerlo. Cada lenguaje sirve para propósito diferente, es tarea constante para el programador buscar [el lenguaje perfecto](https://rodrigoamaya.blogspot.com/2007/01/el-lenguaje-de-programacion-perfecto.html) para que se sienta cómodo y sea más productivo. Pero entre mas de 1000 lenguajes, esta tarea es como buscar a la pareja perfecta para trabajar en la torre de Babel...

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgKD8CKNAbCL480h7dMUS135RWOLry5iDzyp7wl4Lexot5wvyY4HT1ynC5lk5WnqJ74tWTERSOy6jOkhQxa_U73YKJ79wX73Gnjf9XG9xjcIctB3iaM63Qw8vNaTorTGHhyMrMv00yiAFM/s200/babel.jpg)    
Muchos no saben quizás que el programador siempre busca que su código se "vea bien". El programador experimentado puede conocerse por su código: los nombres de variables son significativos, utiliza sangrías (tabulación), comenta su código de forma ordenada, entre otras buenas costumbres:

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEjehWgBjSjiKI00PICPCHQu2zYl4-sLMj-rWRdrX7WMXiWfOoIIhCw6Fa_FmPiEZ88SewMDhBhcZfRucqSU8X66CxwU-KK7OrA8vhBBKaNVk6lBAt3LR0ozYac9WBYhmLkGP3wXQCE_IWs/s1600/codepro.jpg)  
En cambio el novato deja mucho que desear...

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgTmuLU3t50Xpxn-koJYmxO_AymbMoOy85a3skq_h2vfi6dxea28q2NzflQat6NsickWeIQeeg5qYTZ4q5HMvBmj-HLvJYEeELghrkQvivvHBbtQCoA0jcPXMI6b6LwkuqEMwfmOTs71JM/s400/codenewbie.jpg)  
Por algún lado tenemos que empezar, ¿no?. Uno de los problemas principales en la actualidad a la hora de programar, es la falta de modularidad en el código. Muchos (pero muchos) programadores creen que sus clases tienen que hacer de todo. Estas clases todo (o super clases) son en realidad, El Santo Grial del programador... es decir... la búsqueda inalcanzable de:

la ultima clase objeto super abstracta genérica y su jerarquía.

![image](https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEgZNQnlKy-bJk1oX2KZ7aejAOuh8D5nqTzIFg4UXyKQd87i1yOvB26TSX4Sm9ONIptlEmUiOAwdnXYTIdKP9T02cpW8iMC_fUqPQNCYg3-_LEY6TST6OtfZv7w5yw_KFw7cnAY0F2W7lcs/s400/character3.jpg)    
"Imagen tomada del comic: [c0ders](https://www.pello.info/coders/characters.html)
"

Y claro, en esa búsqueda se termina en una bola de código (basura) gigante. Un ejemplo visual perfecto para esto, es un juego que se asemeja mucho al problema que menciono, este juego es [Katamari Damacy](https://es.wikipedia.org/wiki/Katamari_Damacy), que consiste en rodar una bola pegajosa llamada katamari, a lo largo y ancho de distintas pantallas, recolectando todo tipo de objetos hasta que la bola se convierte en una gigantesca esfera de basura... casi lo que sucede con el código de muchos de nosotros:

La analogía es ideal: el código comienza pequeño (como la bola de katamari) y termina acumulando miles de lineas de código inútiles... una clase en donde falla la abstracción y la modularidad. La probabilidad de que un proyecto de software falle esta directamente relacionada con el tamaño del mismo. Y la relación entre lineas de código y bugs (errores) es completamente linear. Menos código significa menos errores. Si el código es mas corto, se evita el síndrome de ML; NL, es decir: "Muy Largo; No leí" (y en ingles TL;DR : "Too Long; Ditn't Read"). Si hay menos código para leer y es más entendible, son mas altas las probabilidades de que alguien realmente lo lea.

Hasta la próxima!